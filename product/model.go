package product

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
	modelName = "detail"
)

// Model 订单信息
type Model struct {
	// 模型继承
	model.Model `json:"_" bson:"_"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

// Buckets 商品信息
type Buckets struct {
	ID           string         `json:"id"`
	Number       int            `json:"number"`
	OriginAmount string         `json:"origin_amount"  bson:"origin_amount"`
	Price        string         `json:"price"`
	Unit         string         `json:"unit"`
	Property     []PropertyInfo `json:"property"`
	Image        string         `json:"image"`
	Amount       string         `json:"amount"`
	Name         string         `json:"name"`
}

type SpecsInfo struct {
	Values []struct {
		Id    int         `json:"id"`
		Image interface{} `json:"image"`
		Value string      `json:"value"`
	} `json:"values"`
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type PropertyInfo struct {
	IsOpenCheckbox bool              `json:"is_open_checkbox"`
	Id             int               `json:"id"`
	Price          int               `json:"price"`
	Values         []PropertySetting `json:"values"`
	Name           string            `json:"name"`
	Desc           *string           `json:"desc,omitempty"`
}

type PropertySetting struct {
	IsDefault int     `json:"is_default,omitempty"`
	Id        int     `json:"id"`
	Code      string  `json:"code"`
	Value     string  `json:"value"`
	Price     float64 `json:"price"`
}

type EntityInfo struct {
	SpecId          string      `json:"spec_id"`
	TradeMark       string      `json:"trade_mark"`
	Id              string      `json:"id"`
	Stock           string      `json:"stock"`
	SpecText        interface{} `json:"spec_text"`
	Spec            interface{} `json:"spec"`
	Image           string      `json:"image"`
	Num             int         `json:"num"`
	Price           float64     `json:"price"`
	MembershipPrice int         `json:"membership_price"  bson:"membership_price"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
