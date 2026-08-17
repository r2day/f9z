package coupon

func (m *Model) CalcBenefit(orderAmount float64) float64 {
	if m.Threshold > 0 && orderAmount < m.Threshold {
		return 0
	}
	switch m.Type {
	case TypeCash, TypeDelivery:
		if m.Amount > orderAmount {
			return orderAmount
		}
		return m.Amount
	case TypePercent:
		if m.Amount <= 0 || m.Amount >= 1 {
			return 0
		}
		return orderAmount * (1 - m.Amount)
	default:
		return 0
	}
}
