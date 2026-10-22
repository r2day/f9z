package f9z_test

import (
	"context"
	"errors"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/open4go/req5rsp/cst"
	"github.com/r2day/db"
	"github.com/r2day/f9z/categories"
	"github.com/r2day/f9z/finance"
	"github.com/r2day/f9z/member"
	"github.com/r2day/f9z/menu"
	"github.com/r2day/f9z/order"
	"github.com/r2day/f9z/pay"
	"github.com/r2day/f9z/product"
	"github.com/r2day/f9z/property"
	"github.com/r2day/f9z/store"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 集成测试：仅在 R2DAY_IT=1 时运行。
// MongoDB 默认 mongodb://127.0.0.1:27017（可用 R2DAY_IT_MONGO 覆盖），使用临时库 f9z_it，结束后删除。
// Redis 默认 127.0.0.1:6379（可用 R2DAY_IT_REDIS 覆盖）的 db 15，只删除本测试用到的 key。
func itSkip(t *testing.T) {
	t.Helper()
	if os.Getenv("R2DAY_IT") != "1" {
		t.Skip("集成测试：设置 R2DAY_IT=1 并启动本地 MongoDB / Redis 后运行")
	}
}

func itDB(t *testing.T) (context.Context, *mongo.Database) {
	t.Helper()
	itSkip(t)
	uri := os.Getenv("R2DAY_IT_MONGO")
	if uri == "" {
		uri = "mongodb://127.0.0.1:27017"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	t.Cleanup(cancel)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		t.Fatal(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		t.Fatalf("无法连接 MongoDB: %v", err)
	}
	d := client.Database("f9z_it")
	_ = d.Drop(ctx)
	t.Cleanup(func() {
		_ = d.Drop(context.Background())
		_ = client.Disconnect(context.Background())
	})
	return ctx, d
}

var itRedisKeys = []string{"store_position", store.CacheInfoPrefix}

func itRedis(t *testing.T) *redis.Client {
	t.Helper()
	itSkip(t)
	addr := os.Getenv("R2DAY_IT_REDIS")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}
	c := redis.NewClient(&redis.Options{Addr: addr, DB: 15})
	if err := c.Ping(context.Background()).Err(); err != nil {
		t.Fatalf("无法连接 Redis: %v", err)
	}
	c.Del(context.Background(), itRedisKeys...)
	prev := db.RDB
	db.RDB = c
	t.Cleanup(func() {
		c.Del(context.Background(), itRedisKeys...)
		db.RDB = prev
		_ = c.Close()
	})
	return c
}

func insert(t *testing.T, ctx context.Context, d *mongo.Database, coll string, doc any) primitive.ObjectID {
	t.Helper()
	res, err := d.Collection(coll).InsertOne(ctx, doc)
	if err != nil {
		t.Fatal(err)
	}
	id, _ := res.InsertedID.(primitive.ObjectID)
	return id
}

// FetchDataAndSaveOnCache 是小程序“附近门店”接口直接调用的函数，这里用 gin 测试上下文驱动真实的 Mongo + Redis。
func TestIT_StoreFetchDataAndSaveOnCache(t *testing.T) {
	ctx, d := itDB(t)
	rdb := itRedis(t)
	prev := store.MongoDatabase
	store.MongoDatabase = d
	t.Cleanup(func() { store.MongoDatabase = prev })

	coll := (&store.Model{}).CollectionName()
	mk := func(merchant, storeID, name string, lng, lat float64) primitive.ObjectID {
		s := &store.Model{StoreID: storeID, Name: name, Lbs: store.LbsInfo{Longitude: lng, Latitude: lat}}
		s.Meta.MerchantID = merchant
		return insert(t, ctx, d, coll, s)
	}
	nearA := mk("merchant-a", "S-A1", "A 近店", 114.0310, 22.5410)
	mk("merchant-a", "S-A2", "A 远店", 114.0800, 22.5400)
	nearB := mk("merchant-b", "S-B1", "B 最近店", 114.0301, 22.5401)

	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest("GET", "/stores/nearby", nil)
	// 请求带有商户 A 的标识，但函数并不按商户过滤
	c.Request.Header.Set("MerchantID", "merchant-a")

	counter, best, all, err := store.FetchDataAndSaveOnCache(c, 114.0300, 22.5400)
	if err != nil {
		t.Fatal(err)
	}
	// 现状：查询条件为空，返回了所有商户的门店（见 PR 问题 1）
	if counter != 3 || len(all) != 3 {
		t.Fatalf("counter=%d len=%d", counter, len(all))
	}
	if best.ID != nearB || !best.BestMatch {
		t.Fatalf("最近门店应为商户 B 的门店: %+v", best)
	}
	for _, s := range all {
		if s.Distance <= 0 || s.Distance > 10 {
			t.Fatalf("距离异常: %s %v", s.Name, s.Distance)
		}
		if s.ID == nearA && s.BestMatch {
			t.Fatal("只有一个门店应标记为 BestMatch")
		}
	}
	// 临时用户坐标在返回前被删除，只剩 3 个门店
	if n := rdb.ZCard(ctx, "store_position").Val(); n != 3 {
		t.Fatalf("store_position 成员数=%d", n)
	}
	// 门店信息按 StoreID 写入缓存
	got, err := store.NewCache(ctx).Load("S-B1")
	if err != nil || got.Name != "B 最近店" || got.Distance <= 0 {
		t.Fatalf("Load=%+v err=%v", got, err)
	}
	if _, err := store.NewCache(ctx).Load("S-404"); !errors.Is(err, redis.Nil) {
		t.Fatalf("不存在的门店应返回 redis.Nil: %v", err)
	}
}

func TestIT_StoreCacheEmptyStoreIDOverwrites(t *testing.T) {
	ctx := context.Background()
	itRedis(t)
	cache := store.NewCache(ctx)
	cases := []struct {
		name string
		m    *store.Model
	}{
		{"第一个门店", &store.Model{Name: "一号店"}},
		{"第二个门店", &store.Model{Name: "二号店"}},
	}
	for _, tc := range cases {
		if err := cache.Dump(tc.m); err != nil {
			t.Fatalf("%s: %v", tc.name, err)
		}
	}
	// 现状：StoreID 为空时所有门店写到同一个 hash 字段 ""，后写覆盖先写（见 PR 问题 4）
	got, err := cache.Load("")
	if err != nil || got.Name != "二号店" {
		t.Fatalf("got=%+v err=%v", got, err)
	}
}

func TestIT_StoreGeoHelpers(t *testing.T) {
	ctx := context.Background()
	rdb := itRedis(t)
	s1 := &store.Model{ID: primitive.NewObjectID(), Lbs: store.LbsInfo{Longitude: 114.031, Latitude: 22.541}}
	s2 := &store.Model{ID: primitive.NewObjectID(), Lbs: store.LbsInfo{Longitude: 114.080, Latitude: 22.540}}
	bad := &store.Model{ID: primitive.NewObjectID(), Name: "非法坐标", Lbs: store.LbsInfo{Longitude: 500, Latitude: 500}}
	store.AddStore(ctx, []*store.Model{s1, s2, bad}) // 非法坐标只记录日志并跳过
	if n := rdb.ZCard(ctx, "store_position").Val(); n != 2 {
		t.Fatalf("成员数=%d", n)
	}

	cases := []struct {
		name     string
		lng, lat float64
		want     []string
	}{
		{"按距离升序", 114.030, 22.540, []string{s1.ID.Hex(), s2.ID.Hex()}},
		{"10km 外无结果", 116.40, 39.90, []string{}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := store.GetStoresByLocation(ctx, tc.lng, tc.lat)
			if fmt.Sprint(got) != fmt.Sprint(tc.want) {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}

	// 现状：用户临时坐标与门店写在同一个 geo key 中，并发请求之间可能把别人的坐标当作“最近门店”（见 PR 问题 5）
	user := store.AddUserLocationForTemporary(ctx, 114.0301, 22.5401)
	near := store.GetStoresByLocation(ctx, 114.0300, 22.5400)
	if len(near) == 0 || near[0] != user {
		t.Fatalf("临时用户坐标应出现在最近结果首位: %v", near)
	}
	if d := store.GetDistance(ctx, s1.ID.Hex(), user, "km"); d <= 0 || d > 1 {
		t.Fatalf("距离=%v", d)
	}
	if d := store.GetDistance(ctx, s1.ID.Hex(), "not-exist", "km"); d != 0 {
		t.Fatalf("不存在的成员距离应为 0: %v", d)
	}
	if n := store.RemoveUserTemporaryLocation(ctx, user); n != 1 {
		t.Fatalf("删除数=%d", n)
	}
	if n := store.RemoveUserTemporaryLocation(ctx, user); n != 0 {
		t.Fatalf("重复删除数=%d", n)
	}
}

// 财务统计按 "meta.tenant_id" 过滤，但模型把租户写在顶层 tenant_id、时间写在 "_.meta.created_time"，
// 所以按模型正常写入的流水统计结果永远为 0（见 PR 问题 2）。
func TestIT_FinanceSummaryFieldMismatch(t *testing.T) {
	ctx, d := itDB(t)
	f := &finance.Model{}
	coll := f.CollectionName()
	now := time.Now().Unix()

	typed := &finance.Model{TenantID: "t1", FlowType: cst.FinanceIncome, Amount: 1000, OrderID: "o1", MemberID: "u1"}
	typed.Meta.CreatedTime = now
	typed.Meta.MerchantID = "t1"
	insert(t, ctx, d, coll, typed)

	f.Init(ctx, d, coll)
	sum, err := f.GetMonthlySummary(ctx, "t1")
	if err != nil {
		t.Fatal(err)
	}
	if fmt.Sprint(sum["total_income"]) != "0" || fmt.Sprint(sum["total_orders"]) != "0" {
		t.Fatalf("现状应统计不到: %v", sum)
	}
	trend, err := f.GetLast7DaysTrend(ctx, "t1")
	if err != nil || len(trend) != 0 {
		t.Fatalf("现状趋势应为空: %v %v", trend, err)
	}

	// 只有手工写成顶层 meta.tenant_id / meta.created_time 的文档才会被统计
	raw := func(flow cst.FinanceStatus, amount int64, orderID, memberID string) bson.M {
		return bson.M{"meta": bson.M{"tenant_id": "t1", "created_time": now}, "flow_type": flow, "amount": amount, "order_id": orderID, "member_id": memberID}
	}
	insert(t, ctx, d, coll, raw(cst.FinanceIncome, 100, "o1", "u1"))
	insert(t, ctx, d, coll, raw(cst.FinanceIncome, 200, "o2", "u1"))
	insert(t, ctx, d, coll, raw(cst.FinanceRefund, 30, "o1", "u2"))
	insert(t, ctx, d, coll, bson.M{"meta": bson.M{"tenant_id": "t2", "created_time": now}, "flow_type": cst.FinanceIncome, "amount": int64(999)})

	sum, err = f.GetMonthlySummary(ctx, "t1")
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{"total_income": "300", "total_refund": "30", "total_orders": "2", "total_users": "2", "month": time.Now().Format("2006-01")}
	for k, v := range want {
		if fmt.Sprint(sum[k]) != v {
			t.Fatalf("%s=%v want %s (%v)", k, sum[k], v, sum)
		}
	}
	trend, err = f.GetLast7DaysTrend(ctx, "t1")
	if err != nil || len(trend) != 1 {
		t.Fatalf("trend=%v err=%v", trend, err)
	}
	if fmt.Sprint(trend[0]["income"]) != "300" || fmt.Sprint(trend[0]["refund"]) != "30" || trend[0]["order_count"] != 2 {
		t.Fatalf("trend=%v", trend[0])
	}
}

func TestIT_MemberQueries(t *testing.T) {
	ctx, d := itDB(t)
	coll := (&member.Model{}).CollectionName()
	mk := func(merchant, wx string, level int) {
		m := &member.Model{}
		m.Identity.WxLoginID = wx
		m.Identity.LevelID = level
		m.Identity.Level = fmt.Sprintf("L%d", level)
		m.Meta.MerchantID = merchant
		insert(t, ctx, d, coll, m)
	}
	mk("merchant-a", "wx-1", 1)
	mk("merchant-b", "wx-1", 1)
	mk("merchant-b", "wx-2", 2)

	cases := []struct {
		name, wx, merchant string
		wantErr            error
	}{
		{"同一 openid 按商户隔离：A", "wx-1", "merchant-a", nil},
		{"同一 openid 按商户隔离：B", "wx-1", "merchant-b", nil},
		{"其他商户查不到", "wx-2", "merchant-a", mongo.ErrNoDocuments},
		{"商户为空", "wx-1", "", mongo.ErrNoDocuments},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := &member.Model{}
			m.Init(ctx, d, coll)
			got, err := m.GetByWxLoginID(tc.wx, tc.merchant)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err=%v", err)
			}
			if tc.wantErr == nil && (got.Meta.MerchantID != tc.merchant || got.Identity.WxLoginID != tc.wx) {
				t.Fatalf("got=%+v", got)
			}
		})
	}

	// 现状：等级统计没有商户条件，统计的是所有商户的会员（见 PR 问题 3）
	m := &member.Model{}
	m.Init(ctx, d, coll)
	stats, err := m.GetLevelStatistics()
	if err != nil {
		t.Fatal(err)
	}
	sm := member.ToLevelStatMap(stats)
	if sm.GetLevelStat(1).MemberCount != 2 || sm.GetLevelStat(2).MemberCount != 1 || sm.GetLevelStat(1).LevelName != "L1" {
		t.Fatalf("stats=%+v", stats)
	}
}

func TestIT_GetByOrderID(t *testing.T) {
	ctx, d := itDB(t)
	o := &order.Model{}
	o.Identity.OrderNo = "NO-1"
	insert(t, ctx, d, o.CollectionName(), o)
	// pay / categories 模型没有 identity 字段，只有手工文档才能命中
	insert(t, ctx, d, (&pay.Model{}).CollectionName(), &pay.Model{})
	insert(t, ctx, d, (&categories.Model{}).CollectionName(), &categories.Model{Name: "饮品"})

	cases := []struct {
		name string
		fn   func(no string) (int, error)
		no   string
		want int
	}{
		{"order 命中", func(no string) (int, error) {
			m := &order.Model{}
			m.Init(ctx, d, m.CollectionName())
			r, err := m.GetByOrderID(no)
			return len(r), err
		}, "NO-1", 1},
		{"order 未命中返回空列表", func(no string) (int, error) {
			m := &order.Model{}
			m.Init(ctx, d, m.CollectionName())
			r, err := m.GetByOrderID(no)
			return len(r), err
		}, "NO-404", 0},
		{"pay 模型无 identity 字段，查不到", func(no string) (int, error) {
			m := &pay.Model{}
			m.Init(ctx, d, m.CollectionName())
			r, err := m.GetByOrderID(no)
			return len(r), err
		}, "", 0},
		{"categories 模型无 identity 字段，查不到", func(no string) (int, error) {
			m := &categories.Model{}
			m.Init(ctx, d, m.CollectionName())
			r, err := m.GetByOrderID(no)
			return len(r), err
		}, "", 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			n, err := tc.fn(tc.no)
			if err != nil || n != tc.want {
				t.Fatalf("n=%d err=%v", n, err)
			}
		})
	}
}

func TestIT_MenuAndProduct(t *testing.T) {
	ctx, d := itDB(t)
	pColl := (&product.Model{}).CollectionName()
	propID := insert(t, ctx, d, (&property.Model{}).CollectionName(), &property.Model{Price: 3})
	p1 := insert(t, ctx, d, pColl, &product.Model{Price: 10, Properties: []primitive.ObjectID{propID}})
	p2 := insert(t, ctx, d, pColl, &product.Model{Price: 20})
	menuID := primitive.NewObjectID()

	t.Run("普通菜谱按 Products 查询", func(t *testing.T) {
		m := &menu.Model{ID: menuID, Products: []primitive.ObjectID{p1}}
		m.Init(ctx, d, m.CollectionName())
		got, err := m.GetProducts()
		if err != nil || len(got) != 1 || got[0].ID != p1 || got[0].CombID != "" {
			t.Fatalf("got=%v err=%v", got, err)
		}
	})

	t.Run("套餐模式：标记套餐信息并覆盖价格，非法 ID 的分组被跳过", func(t *testing.T) {
		m := &menu.Model{ID: menuID, IsComboMode: true, Combo: []menu.ComboInfo{
			{Products: []string{"bad-id"}},
			{Products: []string{p1.Hex(), p2.Hex()}, Price: 15, Quantity: 2},
			{Products: []string{p2.Hex()}, Quantity: 1},
		}}
		m.Init(ctx, d, m.CollectionName())
		got, err := m.GetProducts()
		if err != nil || len(got) != 3 {
			t.Fatalf("got=%d err=%v", len(got), err)
		}
		for _, p := range got[:2] {
			if p.CombID != menuID.Hex() || p.CombIndex != 1 || p.Price != 15 || p.MaxPurchaseAllowed != 2 {
				t.Fatalf("套餐 1: %+v", p)
			}
		}
		last := got[2]
		if last.CombIndex != 2 || last.Price != 20 || last.OriginPrice != 20 {
			t.Fatalf("套餐 2 价格为 0 时应保留原价: %+v", last)
		}
	})

	t.Run("商品属性与销量", func(t *testing.T) {
		p := &product.Model{ID: p1, Properties: []primitive.ObjectID{propID}}
		p.Init(ctx, d, pColl)
		props, err := p.GetProperty()
		if err != nil || len(props) != 1 || props[0].Price != 3 {
			t.Fatalf("props=%v err=%v", props, err)
		}
		for _, inc := range []int{3, 2, -1} {
			if err := p.SalesCount(inc); err != nil {
				t.Fatal(err)
			}
		}
		var doc bson.M
		if err := d.Collection(pColl).FindOne(ctx, bson.M{"_id": p1}).Decode(&doc); err != nil {
			t.Fatal(err)
		}
		if fmt.Sprint(doc["sales"]) != "4" {
			t.Fatalf("sales=%v", doc["sales"])
		}
	})
}
