package order

import (
	"github.com/open4go/model"
	"github.com/open4go/req5rsp/cst"
	"github.com/r2day/f9z/pay"
	"github.com/r2day/f9z/product"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "order_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_flow"
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
	// 订单来源(系统根据订单来源终端自动赋值）
	Source SourceType `json:"source" bson:"source"`
	// 取货方式（客户可选择）
	Pick PickUpType `json:"pick" bson:"pick"`
	// 订单状态时间节点
	Stp StatusTimePoint `json:"stp" bson:"stp"`
	// 状态
	Status cst.OrderStatus `json:"status" bson:"status"`
	// 订单标识信息
	Identity IdentityInfo `json:"identity" bson:"identity"`
	// 客户信息
	Customer CustomerInfo `json:"customer" bson:"customer"`
	// 商户信息
	Merchant MerchantInfo `json:"merchant" bson:"merchant"`
	// 商品列表
	Buckets []*Buckets `json:"buckets" bson:"buckets"`
	// 支付信息
	Pay pay.PayInfo `json:"pay" bson:"pay"`
	// 价格信息
	Price PriceInfo `json:"price" bson:"price"`
}

// IdentityInfo 标识信息
type IdentityInfo struct {
	// 备注信息 （后台管理员操作）
	Remark string `json:"remark" bson:"remark"`
	// 脚注信息（客户提交）
	Postscript string `json:"postscript" bson:"postscript"`
	// 排队序号 （自动生成）
	SortNum string `json:"sort_num" bson:"sort_num"`
	// 订单号 （自动生成）
	OrderNo string `json:"order_no" bson:"order_no"`
	// 座位号 (自主选择/店铺分配）
	TableNo string `json:"table_no" bson:"table_no"`
}

type PriceInfo struct {
	// 原价(单位元）
	OriginalPrice string `json:"original_price" bson:"original_price"`
	//Sale price 特惠價(单位元）
	SalePrice string `json:"sale_price" bson:"sale_price"`
	//Reduced price 減價(单位元）
	ReducedPrice string `json:"reduced_price" bson:"reduced_price"`
	//Retail price 零售價(单位元）
	RetailPrice string `json:"retail_price" bson:"retail_price"`
	// 支付金额(单位：分）
	PayPrice int64 `json:"pay_price" bson:"pay_price"`
}

// StatusTimePoint 订单状态关键时间点
type StatusTimePoint struct {
	// 订单创建时间
	CreatedAt int64 `json:"created_at" bson:"created_at"`
	// 	PayedAt 支付时间
	PayedAt int64 `json:"payed_at" bson:"payed_at"`
	// 	ReviewAt 评价时间
	ReviewAt int64 `json:"review_at" bson:"review_at"`
	// 	CompletedAt 完成时间
	CompletedAt int64 `json:"completed_at" bson:"completed_at"`
}

// CustomerInfo 客户信息
type CustomerInfo struct {
	Mobile  string `json:"mobile" bson:"mobile"`
	Name    string `json:"name" bson:"name"`
	Id      string `json:"id" bson:"id"`
	Account string `json:"account" bson:"account"`
}

// MerchantInfo 商户信息
type MerchantInfo struct {
	// 联系方式
	Mobile string `json:"mobile" bson:"mobile"`
	// 名称
	Name string `json:"name" bson:"name"`
	// 唯一标识
	Id string `json:"id" bson:"id"`
	// 位置
	Address string `json:"address" bson:"address"`
}

// Buckets 商品信息
type Buckets struct {
	// 商品id
	ID string `json:"id"`
	// 名称
	Name string `json:"name"`
	// 数量
	Number int `json:"number"`
	// 原始价格
	OriginAmount string `json:"origin_amount"  bson:"origin_amount"`
	// 单位
	Unit string `json:"unit"  bson:"unit"`
	// 属性
	Property []product.PropertyInfo `json:"property"  bson:"property"`
	// 图片
	Image string `json:"image"  bson:"image"`
	// 数量
	Amount string `json:"amount"  bson:"amount"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
