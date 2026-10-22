package member

import (
	"errors"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestLevelStatMap(t *testing.T) {
	m := ToLevelStatMap([]LevelStat{
		{LevelID: 1, LevelName: "银卡", MemberCount: 3},
		{LevelID: 2, LevelName: "金卡", MemberCount: 1},
		{LevelID: 1, LevelName: "银卡(新)", MemberCount: 5}, // 重复 ID：后者覆盖
	})
	cases := []struct {
		name string
		id   int
		want LevelStat
	}{
		{"存在", 2, LevelStat{LevelID: 2, LevelName: "金卡", MemberCount: 1}},
		{"重复 ID 取最后一条", 1, LevelStat{LevelID: 1, LevelName: "银卡(新)", MemberCount: 5}},
		{"不存在返回默认值", 9, LevelStat{LevelID: 9, LevelName: "未知层级"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := m.GetLevelStat(tc.id); got != tc.want {
				t.Fatalf("got %+v", got)
			}
		})
	}
	if len(ToLevelStatMap(nil)) != 0 {
		t.Fatal("nil 输入应得到空 map")
	}
}

// 未传商户号时直接拒绝，不访问数据库（这里未初始化 Handler）。
func TestGetByWxLoginIDRequiresMerchant(t *testing.T) {
	got, err := (&Model{}).GetByWxLoginID("wx-1", "")
	if got != nil || !errors.Is(err, mongo.ErrNoDocuments) {
		t.Fatalf("got=%v err=%v", got, err)
	}
}
