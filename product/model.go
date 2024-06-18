package product

import (
	"github.com/open4go/model"
	"github.com/r2day/f9z/property"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ProductID      int64         `json:"product_id" bson:"product_id"`
	CreatedAt      time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at" bson:"updated_at"`
	GoodsMealsInfo []interface{} `json:"goods_meals_info,omitempty" bson:"goods_meals_info"`
	IsAdd          int           `json:"is_add,omitempty" bson:"is_add"`
	UseSpec        bool          `json:"use_spec" bson:"use_spec"`
	Entity         []EntityInfo  `json:"entity"`
	StallCode      string        `json:"stall_code" bson:"stall_code"`
	Specs          []SpecsInfo   `json:"specs"`
	IsFollowSuit   int           `json:"is_follow_suit,omitempty"`
	IsLabel        int           `json:"is_label"`
	// 销售属性
	SellTimeStatus  int     `json:"sell_time_status" bson:"sell_time_status"`
	IsSell          bool    `json:"is_sell" bson:"is_sell"`
	PackCost        string  `json:"pack_cost" bson:"pack_cost"`
	UnitType        int     `json:"unit_type" bson:"unit_type"`
	Sort            int     `json:"sort"`
	Price           float64 `json:"price"`
	Unit            string  `json:"unit"`
	MembershipPrice int     `json:"membership_price" bson:"membership_price"`

	// 发布属性
	PublishClient []string `json:"publish_client" bson:"publish_client"`

	// 商品本身属性
	Name          string `json:"name"`
	Type          int    `json:"type"`
	GoodsType     int    `json:"goods_type" bson:"goods_type"`
	Content       string `json:"content"`
	UseProperty   int    `json:"use_property" bson:"use_property"`
	IsUseProperty bool   `json:"is_use_property" bson:"is_use_property"`
	CategoryID    string `json:"category_id" bson:"category_id"`
	// 属性列表
	Properties []primitive.ObjectID `json:"properties" bson:"properties"`
	// 属性列表（仅用于展示，读取property 表后渲染到这里）
	Property []property.Model `json:"property" bson:"-"`
	// 统计数据 & 限制
	Sales     int `json:"sales"`
	Stock     int `json:"stock"`
	MinBuyNum int `json:"min_buy_num" bson:"min_buy_num"`

	// 展示
	Images   string   `json:"images"`
	CoverImg string   `json:"cover_img" bson:"cover_img"`
	ImageArr []string `json:"imageArr"`
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
