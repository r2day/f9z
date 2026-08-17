package campaign

import "time"

func (m *Model) InTime(now int64) bool {
	if now == 0 {
		now = time.Now().Unix()
	}
	if m.StartAt > 0 && now < m.StartAt {
		return false
	}
	if m.EndAt > 0 && now > m.EndAt {
		return false
	}
	return true
}

func (m *Model) IsOnline() bool {
	return m.Status == StatusOnline && m.InTime(0)
}

// CalcBenefit 按订单金额试算优惠（元）。发券/买赠/充值不在这里扣金额。
func (m *Model) CalcBenefit(orderAmount float64) float64 {
	if orderAmount < 0 {
		return 0
	}
	rule := m.Rule
	if rule.Threshold > 0 && orderAmount < rule.Threshold {
		return 0
	}
	switch m.Type {
	case TypeFullReduce, TypeNewcomer:
		return clamp(rule.Reduce, 0, orderAmount)
	case TypeDiscount:
		if rule.Discount <= 0 || rule.Discount >= 1 {
			return 0
		}
		off := orderAmount * (1 - rule.Discount)
		if rule.MaxReduce > 0 && off > rule.MaxReduce {
			off = rule.MaxReduce
		}
		return clamp(off, 0, orderAmount)
	case TypeDelivery:
		return clamp(rule.DeliveryReduce, 0, orderAmount)
	case TypeFlash:
		if rule.Reduce > 0 {
			return clamp(rule.Reduce, 0, orderAmount)
		}
		if rule.Discount > 0 && rule.Discount < 1 {
			return clamp(orderAmount*(1-rule.Discount), 0, orderAmount)
		}
	}
	return 0
}

func clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if max > 0 && v > max {
		return max
	}
	return v
}
