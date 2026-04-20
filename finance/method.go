package finance

import (
	"context"
	"github.com/open4go/req5rsp/cst"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/open4go/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetMonthlySummary 获取当月财务总结
// 返回：总收入（分）、总退款（分）、总订单数、总参与用户数
func (m *Model) GetMonthlySummary(ctx context.Context, tenantID string) (map[string]interface{}, error) {
	coll := m.Context.Handler.Collection(m.Context.Collection)

	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Unix()

	filter := bson.D{
		{Key: "meta.tenant_id", Value: tenantID}, // 注意使用 meta. 前缀
		{Key: "meta.created_time", Value: bson.D{{Key: "$gte", Value: startOfMonth}}},
	}

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "total_income", Value: bson.D{{Key: "$sum", Value: bson.D{{Key: "$cond", Value: bson.A{
				bson.D{{Key: "$eq", Value: bson.A{"$flow_type", cst.FinanceIncome}}},
				"$amount", 0,
			}}}}}},
			{Key: "total_refund", Value: bson.D{{Key: "$sum", Value: bson.D{{Key: "$cond", Value: bson.A{
				bson.D{{Key: "$eq", Value: bson.A{"$flow_type", cst.FinanceRefund}}},
				"$amount", 0,
			}}}}}},
			{Key: "total_orders", Value: bson.D{{Key: "$addToSet", Value: "$order_id"}}},
			{Key: "total_users", Value: bson.D{{Key: "$addToSet", Value: "$member_id"}}},
		}}},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		log.Log(ctx).WithField("tenant_id", tenantID).WithError(err).Error("获取当月财务总结失败")
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return map[string]interface{}{
			"total_income": 0,
			"total_refund": 0,
			"total_orders": 0,
			"total_users":  0,
			"month":        now.Format("2006-01"),
		}, nil
	}

	r := results[0]
	return map[string]interface{}{
		"total_income": r["total_income"],
		"total_refund": r["total_refund"],
		"total_orders": len(r["total_orders"].(primitive.A)),
		"total_users":  len(r["total_users"].(primitive.A)),
		"month":        now.Format("2006-01"),
	}, nil
}

// GetLast7DaysTrend 获取最近7天每日财务趋势
func (m *Model) GetLast7DaysTrend(ctx context.Context, tenantID string) ([]map[string]interface{}, error) {
	coll := m.Context.Handler.Collection(m.Context.Collection)

	end := time.Now()
	start := end.AddDate(0, 0, -6) // 最近7天

	filter := bson.D{
		{Key: "meta.tenant_id", Value: tenantID},
		{Key: "meta.created_time", Value: bson.D{
			{Key: "$gte", Value: start.Unix()},
			{Key: "$lte", Value: end.Unix()},
		}},
	}

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: bson.D{{Key: "$dateToString", Value: bson.D{
				{Key: "format", Value: "%Y-%m-%d"},
				{Key: "date", Value: bson.D{{Key: "$toDate", Value: bson.D{{Key: "$multiply", Value: bson.A{"$meta.created_time", 1000}}}}}},
			}}}},
			{Key: "income", Value: bson.D{{Key: "$sum", Value: bson.D{{Key: "$cond", Value: bson.A{
				bson.D{{Key: "$eq", Value: bson.A{"$flow_type", cst.FinanceIncome}}},
				"$amount", 0,
			}}}}}},
			{Key: "refund", Value: bson.D{{Key: "$sum", Value: bson.D{{Key: "$cond", Value: bson.A{
				bson.D{{Key: "$eq", Value: bson.A{"$flow_type", cst.FinanceRefund}}},
				"$amount", 0,
			}}}}}},
			{Key: "order_count", Value: bson.D{{Key: "$addToSet", Value: "$order_id"}}},
		}}},
		{{Key: "$sort", Value: bson.D{{Key: "_id", Value: 1}}}},
	}

	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		log.Log(ctx).WithError(err).Error("获取最近7天财务趋势失败")
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	trend := make([]map[string]interface{}, 0, len(results))
	for _, r := range results {
		trend = append(trend, map[string]interface{}{
			"date":        r["_id"],
			"income":      r["income"],
			"refund":      r["refund"],
			"order_count": len(r["order_count"].(primitive.A)),
		})
	}

	return trend, nil
}
