package member

import (
	"github.com/open4go/model"
	"github.com/r2day/f9z/product"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "member_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_data"
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
	// 用户来源
	From UserFrom `json:"source" bson:"source"`
	// 订单标识信息
	Identity IdentityInfo `json:"identity" bson:"identity"`
	// 客户信息
}

// IdentityInfo 标识信息
type IdentityInfo struct {
	// 备注信息 （后台管理员操作）
	AccountID string `json:"account_id" bson:"account_id"`
	// 昵称
	Nickname string `json:"nickname" bson:"nickname"`
	// 用户名
	UserName string `json:"user_name" bson:"user_name"`
	// 生日
	Birthday string `json:"birthday" bson:"birthday"`
	// 性别
	Gender UserGender `json:"gender" bson:"gender"`
	// 微信登陆信息
	WxLoginID string `json:"wx_login_id" bson:"wx_login_id"`
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
	// 支付中
	PayAt int64 `json:"pay_at" bson:"pay_at"`
	// 	PayedAt 支付成功时间
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
