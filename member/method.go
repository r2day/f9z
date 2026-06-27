package member

import (
	"go.mongodb.org/mongo-driver/bson"
)

func (m *Model) GetByWxLoginID(id string) (*Model, error) {

	result := new(Model)
	coll := m.Context.Handler.Collection(m.Context.Collection)
	filter := bson.D{bson.E{Key: "identity.wx_login_id", Value: id}}

	// 获取数据列表
	err := coll.FindOne(m.Context.Context, filter).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetLevelStatistics 按层级分组统计会员数量
func (m *Model) GetLevelStatistics() ([]LevelStat, error) {
	coll := m.Context.Handler.Collection(m.Context.Collection)

	// MongoDB Aggregation Pipeline
	pipeline := bson.A{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "identity.level_id", Value: bson.D{{Key: "$exists", Value: true}}},
			}},
		},
		bson.D{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: "$identity.level_id"},                                     // 按 level_id 分组
				{Key: "level_name", Value: bson.D{{Key: "$first", Value: "$identity.level"}}}, // 取层级名称
				{Key: "member_count", Value: bson.D{{Key: "$sum", Value: 1}}},                 // 统计数量
			}},
		},
		bson.D{
			{Key: "$sort", Value: bson.D{{Key: "_id", Value: 1}}}, // 按 level_id 升序排序
		},
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "level_id", Value: "$_id"},
				{Key: "level_name", Value: 1},
				{Key: "member_count", Value: 1},
				{Key: "_id", Value: 0},
			}},
		},
	}

	cursor, err := coll.Aggregate(m.Context.Context, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(m.Context.Context)

	var results []LevelStat
	if err = cursor.All(m.Context.Context, &results); err != nil {
		return nil, err
	}

	return results, nil
}

// LevelStat 返回结构
type LevelStat struct {
	LevelID     int    `json:"level_id" bson:"level_id"`
	LevelName   string `json:"level_name" bson:"level_name"`
	MemberCount int    `json:"member_count" bson:"member_count"`
}

// LevelStatMap 将 LevelStat 列表转换为 map，方便通过 level_id 快速查找
type LevelStatMap map[int]LevelStat

// ToLevelStatMap 将统计结果转换为 Map
func ToLevelStatMap(stats []LevelStat) LevelStatMap {
	m := make(LevelStatMap, len(stats))
	for _, stat := range stats {
		m[stat.LevelID] = stat
	}
	return m
}

// GetLevelStat 根据 level_id 获取统计信息（带默认值）
func (m LevelStatMap) GetLevelStat(levelID int) LevelStat {
	if stat, ok := m[levelID]; ok {
		return stat
	}
	// 返回默认空值
	return LevelStat{
		LevelID:     levelID,
		LevelName:   "未知层级",
		MemberCount: 0,
	}
}
