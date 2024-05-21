package pay

import (
	"github.com/open4go/model"
	"github.com/open4go/req5rsp/cst"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "pay_"
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
}

type PayInfo struct {
	// 预支付信息
	PreTrade PrepareTrade `json:"prepare_trade" bson:"prepare_trade"`
	// Amount 显示金额
	Amount string `json:"amount"`
	// 支付总金额(单位：分）
	PayAmount int64 `json:"pay_amount" bson:"pay_amount"`
	// 第三方支付单号
	TransactionID string `json:"transaction_id" bson:"transaction_id"`
	// 支付成功时间
	PayedAt int `json:"payed_at" bson:"payed_at"`
	// 退款成功时间
	RefundedAt int `json:"refunded_at" bson:"refunded_at"`
	// Channel 支付渠道
	Channel cst.ChannelType `json:"channel" bson:"channel"`
	// Method 支付方式
	Method cst.PayMethod `json:"method" bson:"method"`
	// Status 支付状态 0未支付 1已支付 2退款中 3已退款
	Status cst.PayStatus `json:"status" bson:"status"`
}

// PrepareTrade 预支付返回信息
// 例如：微信下单后产生的签名值
// 可以用于继续完成支付
type PrepareTrade struct {
	// NonceStr 随机字符串
	NonceStr string `json:"noncestr"`
	// Package 固定值
	Package string `json:"package"`
	// Timestamp 时间戳（单位：秒）
	Timestamp string `json:"timestamp"`
	// Sign 签名
	Sign string `json:"sign"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
