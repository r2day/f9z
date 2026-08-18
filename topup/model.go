package topup

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "store_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_config"
	// 这个需要用户根据具体业务完成设定
	modelName = "topup"
)

// Model 充值列表配置
type Model struct {
	// 模型继承
	model.Model `json:"_" bson:"_"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Sales      int                `json:"sales"   bson:"sales"`
	Image      string             `json:"image"   bson:"image"`
	Value      string             `json:"value"   bson:"value"`
	DeletedAt  time.Time          `json:"deleted_at"   bson:"deleted_at"`
	Sort       int                `json:"sort"   bson:"sort"`
	StatusText string             `json:"status_text"   bson:"status_text"`
	Status     int                `json:"status"   bson:"status"`
	CreatedAt  time.Time          `json:"created_at"   bson:"created_at"`
	Desc       string             `json:"desc"   bson:"desc"`
	StoreId    int                `json:"store_id"   bson:"store_id"`
	Gifts      []string           `json:"gifts"   bson:"gifts"`
	Type       int                `json:"type"   bson:"type"`
	FullImage  string             `json:"full_image"   bson:"full_image"`
	Name       string             `json:"name"   bson:"name"`
	SellPrice  string             `json:"sell_price"   bson:"sell_price"`
	// 商户id (商品需要选择对应的销售门店)
	MerchantId string `json:"merchant_id"   bson:"merchant_id"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
