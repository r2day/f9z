package campaign

import (
	"math"
	"testing"
	"time"
)

func TestCalcBenefitTable(t *testing.T) {
	cases := []struct {
		name   string
		m      Model
		amount float64
		want   float64
	}{
		{"负金额", Model{Type: TypeFullReduce, Rule: Rule{Reduce: 10}}, -1, 0},
		{"满减未达门槛", Model{Type: TypeFullReduce, Rule: Rule{Threshold: 50, Reduce: 10}}, 49.99, 0},
		{"满减达到门槛", Model{Type: TypeFullReduce, Rule: Rule{Threshold: 50, Reduce: 10}}, 50, 10},
		{"满减不超过订单金额", Model{Type: TypeFullReduce, Rule: Rule{Reduce: 30}}, 20, 20},
		{"满减为负按 0", Model{Type: TypeFullReduce, Rule: Rule{Reduce: -5}}, 20, 0},
		{"新人立减", Model{Type: TypeNewcomer, Rule: Rule{Reduce: 5}}, 20, 5},
		{"折扣 8 折", Model{Type: TypeDiscount, Rule: Rule{Discount: 0.8}}, 100, 20},
		{"折扣封顶", Model{Type: TypeDiscount, Rule: Rule{Discount: 0.8, MaxReduce: 12}}, 100, 12},
		{"折扣=0 无效", Model{Type: TypeDiscount, Rule: Rule{Discount: 0}}, 100, 0},
		{"折扣=1 无效", Model{Type: TypeDiscount, Rule: Rule{Discount: 1}}, 100, 0},
		{"折扣>1 无效", Model{Type: TypeDiscount, Rule: Rule{Discount: 1.2}}, 100, 0},
		{"配送减免", Model{Type: TypeDelivery, Rule: Rule{DeliveryReduce: 3}}, 100, 3},
		{"秒杀优先用立减", Model{Type: TypeFlash, Rule: Rule{Reduce: 8, Discount: 0.5}}, 100, 8},
		{"秒杀折扣", Model{Type: TypeFlash, Rule: Rule{Discount: 0.5}}, 100, 50},
		{"秒杀无规则", Model{Type: TypeFlash}, 100, 0},
		{"发券不扣金额", Model{Type: TypeCoupon, Rule: Rule{Reduce: 10}}, 100, 0},
		{"买赠不扣金额", Model{Type: TypeGift, Rule: Rule{Reduce: 10}}, 100, 0},
		{"充值不扣金额", Model{Type: TypeRecharge, Rule: Rule{Reduce: 10}}, 100, 0},
		{"未知类型", Model{Type: "x", Rule: Rule{Reduce: 10}}, 100, 0},
		{"订单金额为 0 时满减不封顶（clamp 的 max=0 视为不限）", Model{Type: TypeFullReduce, Rule: Rule{Reduce: 10}}, 0, 10},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.m.CalcBenefit(tc.amount); math.Abs(got-tc.want) > 1e-9 {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
}

func TestInTime(t *testing.T) {
	cases := []struct {
		name       string
		start, end int64
		now        int64
		want       bool
	}{
		{"不限时", 0, 0, 100, true},
		{"未开始", 200, 0, 100, false},
		{"开始时刻", 100, 0, 100, true},
		{"已结束", 0, 50, 100, false},
		{"结束时刻", 0, 100, 100, true},
		{"区间内", 50, 150, 100, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := &Model{StartAt: tc.start, EndAt: tc.end}
			if got := m.InTime(tc.now); got != tc.want {
				t.Fatalf("got %v", got)
			}
		})
	}
}

func TestIsOnline(t *testing.T) {
	now := time.Now().Unix()
	cases := []struct {
		name string
		m    Model
		want bool
	}{
		{"上线且不限时", Model{Status: StatusOnline}, true},
		{"草稿", Model{Status: StatusDraft}, false},
		{"下线", Model{Status: StatusOffline}, false},
		{"上线但已过期", Model{Status: StatusOnline, EndAt: now - 3600}, false},
		{"上线但未开始", Model{Status: StatusOnline, StartAt: now + 3600}, false},
		{"上线且在区间内", Model{Status: StatusOnline, StartAt: now - 60, EndAt: now + 3600}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.m.IsOnline(); got != tc.want {
				t.Fatalf("got %v", got)
			}
		})
	}
}
