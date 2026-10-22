package property

import "testing"

func TestGetPriceSumByIds(t *testing.T) {
	m := &Model{Values: []PropertySetting{{Id: 1, Price: 100}, {Id: 2, Price: 250}, {Id: 3, Price: 0}}}
	cases := []struct {
		name string
		ids  []int
		want int64
	}{
		{"空", nil, 0},
		{"单个", []int{2}, 250},
		{"多个", []int{1, 2, 3}, 350},
		{"重复 ID 只计一次", []int{1, 1, 1}, 100},
		{"不存在的 ID 忽略", []int{9}, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := m.GetPriceSumByIds(tc.ids)
			if err != nil || got != tc.want {
				t.Fatalf("got %d err %v", got, err)
			}
		})
	}
}
