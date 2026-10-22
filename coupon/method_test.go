package coupon

import (
	"math"
	"testing"
)

func TestCalcBenefit(t *testing.T) {
	cases := []struct {
		name   string
		m      Model
		amount float64
		want   float64
	}{
		{"现金券", Model{Type: TypeCash, Amount: 5}, 20, 5},
		{"现金券不超过订单金额", Model{Type: TypeCash, Amount: 30}, 20, 20},
		{"未达门槛", Model{Type: TypeCash, Threshold: 50, Amount: 5}, 49, 0},
		{"达到门槛", Model{Type: TypeCash, Threshold: 50, Amount: 5}, 50, 5},
		{"配送券", Model{Type: TypeDelivery, Amount: 3}, 20, 3},
		{"折扣券 9 折", Model{Type: TypePercent, Amount: 0.9}, 100, 10},
		{"折扣券=0 无效", Model{Type: TypePercent, Amount: 0}, 100, 0},
		{"折扣券>=1 无效", Model{Type: TypePercent, Amount: 1}, 100, 0},
		{"赠品券不扣金额", Model{Type: TypeGift, Amount: 5}, 100, 0},
		{"未知类型", Model{Type: "x", Amount: 5}, 100, 0},
		// 现状：负金额的现金券会原样返回负数（没有下限保护）。
		{"现金券负金额（现状）", Model{Type: TypeCash, Amount: -5}, 20, -5},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.m.CalcBenefit(tc.amount); math.Abs(got-tc.want) > 1e-9 {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}
