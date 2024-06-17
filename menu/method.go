package menu

import (
	"errors"
	"github.com/open4go/log"
	"github.com/r2day/f9z/product"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetProducts 获取该菜谱下的所有商品
func (m *Model) GetProducts() ([]*product.Model, error) {

	productModel := product.Model{}
	results := make([]*product.Model, 0)
	coll := m.Context.Handler.Collection(productModel.CollectionName())
	filter := bson.M{"_id": bson.M{"$in": m.Products}}

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
