package member

import (
	"go.mongodb.org/mongo-driver/bson"
)

func (m *Model) GetByWxLoginID(id string) (*Model, error) {

	result := new(Model)
	coll := m.Context.Handler.Collection(m.Context.Collection)
	filter := bson.D{{Key: "identity.wx_login_id", Value: id}}

	// 获取数据列表
	err := coll.FindOne(m.Context.Context, filter).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
