package property

// GetPriceSumByIds 获取属性配置总价格
func (m *Model) GetPriceSumByIds(ids []int) (int64, error) {
	idSet := make(map[int]struct{})
	for _, id := range ids {
		idSet[id] = struct{}{}
	}

	var totalPrice int64
	for _, val := range m.Values {
		if _, ok := idSet[val.Id]; ok {
			totalPrice += val.Price
		}
	}
	return totalPrice, nil
}
