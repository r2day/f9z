package categories

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "product_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_config"
	// 这个需要用户根据具体业务完成设定
	modelName = "categories"
)

// Model 属性配置
// 例如：辣度
// 例如：温度
type Model struct {
	// 模型继承
	model.Model `json:"_" bson:"_"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name"`
	Sort     int                `json:"sort"`
	Icon     string             `json:"icon"`
	Products []string           `json:"products" bson:"products"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
