package address

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "member_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_config"
	// 这个需要用户根据具体业务完成设定
	modelName = "address"
)

// Model 属性配置
// 例如：辣度
// 例如：温度
type Model struct {
	// 模型继承
	model.Model `json:"_" bson:"_"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// 收货人
	AcceptName string `json:"accept_name" bson:"accept_name"`
	// 手机号
	Mobile string `json:"mobile"  bson:"mobile"`
	// 性别
	Sex int `json:"sex" bson:"sex"`
	// 地址
	Address string `json:"address" bson:"address"`
	// 地址
	AddressShotName string `json:"addressShotName" bson:"address_shot_name"`
	// 经纬度
	Latitude float64 `json:"latitude" bson:"latitude"`
	// 经纬度
	Longitude float64 `json:"longitude" bson:"longitude"`
	// 门牌号
	DoorNum string `json:"door_num" bson:"door_num"`
	// 账号
	AccountID string `json:"account_id" bson:"account_id"`
	//
	Inner bool `json:"inner"  bson:"inner"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
