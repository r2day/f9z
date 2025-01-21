package menu

import (
	"github.com/open4go/model"
	"github.com/r2day/f9z/product"
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
	modelName = "menu"
)

// Model 属性配置
// 例如：辣度
// 例如：温度
type Model struct {
	// 模型继承
	model.Model `json:"_" bson:"_"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name            string             `json:"name"`
	IsShowBackstage int                `json:"is_show_backstage" bson:"is_show_backstage"`
	Sort            int                `json:"sort"`
	GoodsType       int                `json:"goods_type"  bson:"goods_type"`
	IsSell          bool               `json:"is_sell" bson:"is_sell"`
	Icon            string             `json:"icon"`
	// 商品id列表（该菜谱下的所有商品id）
	Products []primitive.ObjectID `json:"products" bson:"products"`
	// GoodsDisplay 货物展示 不存储(兼容当前客户端）
	GoodsDisplay []*product.Model `json:"goods_list" bson:"-"`
	// 发布到的门店
	Stores []string `json:"stores" bson:"stores"`
	// 更新方式
	UpdateType int `json:"update_type" bson:"update_type"`
	// 套餐配置
	Combo []ComboInfo `json:"combo" bson:"combo"`
}

// ComboInfo 套餐信息
type ComboInfo struct {
	Price    int      `json:"price"`
	Quantity int      `json:"quantity"`
	Products []string `json:"products"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
