package property

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
	modelName = "property"
)

// Model 属性配置
// 例如：辣度
// 例如：温度
type Model struct {
	// 模型继承
	model.Model `json:"_" bson:"_"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Checked bool               `json:"checked"  bson:"checked"`
	Price   int64              `json:"price"  bson:"price"`
	Values  []PropertySetting  `json:"values"  bson:"values"`
	Name    string             `json:"name"  bson:"name"`
	Desc    string             `json:"desc"  bson:"desc"`
	// multiple_selection
	MultipleSelection bool `json:"multiple_selection"  bson:"multiple_selection"`
}

// PropertySetting 属性设置
// 例如：辣，中辣，特辣
// 例如：常温，冷，冻
type PropertySetting struct {
	IsDefault int    `json:"is_default,omitempty"`
	Id        int    `json:"id"`
	Code      string `json:"code"`
	Value     string `json:"value"`
	// 单位分
	Price int64 `json:"price"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
