package member

import (
	"errors"
	"github.com/open4go/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *Model) GetByWxLoginID(id string) (*Model, error) {

	result := new(Model)
	coll := m.Context.Handler.Collection(m.Context.Collection)
	filter := bson.D{{Key: "identity.wx_login_id", Value: id}}

	// 获取数据列表
	cursor, err := coll.Find(m.Context.Context, filter)
	if errors.Is(err, mongo.ErrNoDocuments) {
		log.Log().WithField("id", id).Error(err)
		return nil, err
	}
	if err != nil {
		log.Log().WithField("id", id).Error(err)
		return nil, err
	}
	if err = cursor.All(m.Context.Context, &result); err != nil {
		log.Log().WithField("id", id).Error(err)
		return nil, err
	}
	return result, nil
}
