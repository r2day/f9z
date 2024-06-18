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
