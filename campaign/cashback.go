package campaign

import "math"

// ConsumeMinAmount 消费门槛：优先 ConsumeThreshold，其次 Threshold，再 RechargeAmount（兼容）。
func (r Rule) ConsumeMinAmount() float64 {
	if r.ConsumeThreshold > 0 {
		return r.ConsumeThreshold
	}
	if r.Threshold > 0 {
		return r.Threshold
	}
	return r.RechargeAmount
}

// CalcCashback 计算应入账的 gift_balance（元，四舍五入为整数元以匹配 Assets.GiftBalance int）。
// 优先固定金额；否则用比例；都没有则默认 5% 且上限 CashbackMax（若未设上限则 50 元）。
func (r Rule) CalcCashback(orderAmount float64) float64 {
	if orderAmount <= 0 {
		return 0
	}
	var raw float64
	if r.CashbackAmount > 0 {
		raw = r.CashbackAmount
	} else if r.CashbackPercent > 0 {
		raw = orderAmount * r.CashbackPercent
	} else {
		raw = orderAmount * 0.05
	}
	max := r.CashbackMax
	if max <= 0 && r.CashbackAmount <= 0 && r.CashbackPercent <= 0 {
		max = 50
	}
	if max > 0 && raw > max {
		raw = max
	}
	if raw < 0 {
		return 0
	}
	return math.Round(raw)
}
