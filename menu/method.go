package menu

import (
	"errors"
	"github.com/open4go/log"
	"github.com/r2day/f9z"
	"github.com/r2day/f9z/product"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetProducts 获取该菜谱下的所有商品
func (m *Model) GetProducts() ([]*product.Model, error) {

	productModel := product.Model{}
	results := make([]*product.Model, 0)
	allResults := make([]*product.Model, 0)
	coll := m.Context.Handler.Collection(productModel.CollectionName())
	filter := bson.M{"_id": bson.M{"$in": m.Products}}
	if m.IsComboMode {
		// 如果是套餐模式则通过获取套餐列表
		for index, i := range m.Combo {
			productIds, err := f9z.ConvertToObjectID(i.Products)
			if err != nil {
				log.Log(m.Context.Context).WithField("i.Products", i.Products).Error(err)
				// 如果错误直接跳过
				continue
			}
			filter = bson.M{"_id": bson.M{"$in": productIds}}
			// 获取数据列表
			cursor, err := coll.Find(m.Context.Context, filter)
			if errors.Is(err, mongo.ErrNoDocuments) {
				log.Log(m.Context.Context).WithField("m", m).Error(err)
				return nil, err
			}

			if err != nil {
				log.Log(m.Context.Context).WithField("m", m).Error(err)
				return nil, err
			}

			if err = cursor.All(m.Context.Context, &results); err != nil {
				log.Log(m.Context.Context).WithField("m", m).Error(err)
				return nil, err
			}
			// 将该套餐的产品进行标识（但是不影响产品本身，只是在该类套餐中标识）
			for _, p := range results {
				p.CombID = m.ID.Hex()
				// 套餐有多个组合，每个组合设置一组id
				p.CombIndex = index
				// 该类产品的
				p.MaxPurchaseAllowed = i.Quantity
				// TODO 套餐价格(一般会跟原商品价格存在差异，因此这里会重新设定价格，但是这个价格可能会在下单时，
				// 根据商品id查询到的价格结果不一致，所以需要注意优化
				p.Price = i.Price
				// 将所有产品拼接
				allResults = append(allResults, p) // 直接使用 allResults
			}
		}
		// 当所有内容拼接完毕后返回
		return allResults, nil
	}

	// 非套餐类的数据走正常查询逻辑
	// 获取数据列表
	cursor, err := coll.Find(m.Context.Context, filter)
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Log(m.Context.Context).WithField("m", m).Error(err)
		return nil, err
	}

	if err != nil {
		log.Log(m.Context.Context).WithField("m", m).Error(err)
		return nil, err
	}

	if err = cursor.All(m.Context.Context, &results); err != nil {
		log.Log(m.Context.Context).WithField("m", m).Error(err)
		return nil, err
	}
	return results, nil
}
