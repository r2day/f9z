package store

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/open4go/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// ErrNoMerchant 请求中没有可用的商户标识。
var ErrNoMerchant = errors.New("store: merchant id is required")

// requestMerchant 返回本次请求所属商户：优先使用中间件写入 context 的商户，其次是网关写入的 X-Tenant-ID / MerchantID。
func requestMerchant(c *gin.Context) string {
	if c == nil || c.Request == nil {
		return ""
	}
	for _, id := range []string{
		model.GetValueFromCtx(c.Request.Context(), model.MerchantKey),
		c.GetHeader("X-Tenant-ID"),
		c.GetHeader("MerchantID"),
	} {
		if id = strings.TrimSpace(id); id != "" && id != "*" {
			return id
		}
	}
	return ""
}

// FetchDataAndSaveOnCache 获取数据库并且存储到缓存中
func FetchDataAndSaveOnCache(c *gin.Context, longitude, latitude float64) (int64, *Model, []*Model, error) {
	// 标记最匹配的门店
	bestMatchStore := &Model{}
	allStores := make([]*Model, 0)
	// 只查询本商户的门店，geo 数据也按商户分开存放
	merchant := requestMerchant(c)
	if merchant == "" {
		return 0, bestMatchStore, allStores, ErrNoMerchant
	}
	handler := bestMatchStore.Init(c.Request.Context(), MongoDatabase, bestMatchStore.CollectionName())
	counter, err := handler.GetList(bson.D{{Key: "_.meta.merchant_id", Value: merchant}}, &allStores)
	if err != nil {
		return 0, bestMatchStore, allStores, err
	}

	key := merchantGeoKey(merchant)
	// 存入redis 数据库
	addStore(c.Request.Context(), key, allStores)
	// 获取最近的门店(列表)
	storeIDList := getStoresByLocation(c.Request.Context(), key, longitude, latitude)

	// 同时存入用户位置到数据redis（便于下一步计算其到各个门店的距离）
	userLocationId := addUserLocationForTemporary(c.Request.Context(), key, longitude, latitude)
	defer func() {
		removeUserTemporaryLocation(c.Request.Context(), key, userLocationId)
	}()

	// 取出最近的一个门店（用于直接返回给用户）
	storeId := ""
	if len(storeIDList) > 0 {
		storeId = storeIDList[0]
	}
	log.WithField("allStores", allStores).Warning("==========")
	storeCache := NewCache(c.Request.Context())
	for _, store := range allStores {
		if store.ID.Hex() == storeId {
			store.BestMatch = true
			bestMatchStore = store
		}
		// 计算距离(所有门店与客户的距离)
		dis := getDistance(c.Request.Context(), key, store.ID.Hex(), userLocationId, "km")
		store.Distance = dis

		// 存入门店信息
		err := storeCache.Dump(store)
		if err != nil {
			log.WithField("store", store).Error(err)
			continue
		}
	}
	return counter, bestMatchStore, allStores, nil
}
