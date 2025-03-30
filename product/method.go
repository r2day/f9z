package product

import (
	"errors"
	"github.com/open4go/log"
	"github.com/r2day/f9z/property"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetProperty 获取属性配置
func (m *Model) GetProperty() ([]*property.Model, error) {
	// 判断 m.Properties 是否为空
	if len(m.Properties) == 0 {
		return nil, nil
	}
	
	productModel := property.Model{}
	results := make([]*property.Model, 0)
	coll := m.Context.Handler.Collection(productModel.CollectionName())
	filter := bson.M{"_id": bson.M{"$in": m.Properties}}

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

// SalesCount 更新销量信息
func (m *Model) SalesCount(salesIncrement int) error {
	coll := m.Context.Handler.Collection(m.CollectionName())
	// 创建过滤器和更新文档
	filter := bson.M{"_id": m.ID}
	update := bson.M{"$inc": bson.M{"sales": salesIncrement}}

	// 更新产品销量信息
	result, err := coll.UpdateOne(m.Context.Context, filter, update)
	if err != nil {
		log.Log(m.Context.Context).Error(err)
		return err
	}
	log.Log(m.Context.Context).WithField("salesIncrement", salesIncrement).
		WithField("id", m.ID.Hex()).WithField("result", result).
		Debug("after sales count success")
	return nil
}
