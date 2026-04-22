package finance

import (
	"github.com/open4go/model"
	"github.com/open4go/req5rsp/cst"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "finance_"
	collectionNameSuffix = "_flow"
	modelName            = "flow"
)

// Model 财务流水记录
type Model struct {
	model.Model `json:"_" bson:"_"`

	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`

	// 关联信息
	TenantID   string `json:"tenant_id" bson:"tenant_id"`
	StoreID    string `json:"store_id" bson:"store_id"`
	MemberID   string `json:"member_id" bson:"member_id"`
	OrderID    string `json:"order_id" bson:"order_id"`
	MerchantID string `json:"merchant_id" bson:"merchant_id"`

	// 核心字段
	FlowType cst.FinanceStatus `json:"flow_type" bson:"flow_type"` // 流水类型：收入/退款/支出
	PayType  cst.PayMethod     `json:"pay_type" bson:"pay_type"`   // 支付方式

	Amount     int64   `json:"amount" bson:"amount"` // 金额（单位：分）
	RealAmount float64 `json:"real_amount" bson:"real_amount"`

	// 支付平台相关
	TransactionID string `json:"transaction_id" bson:"transaction_id"`
	OutTradeNo    string `json:"out_trade_no" bson:"out_trade_no"`
	RefundID      string `json:"refund_id" bson:"refund_id,omitempty"`

	// 状态与描述
	Remark   string `json:"remark" bson:"remark"`
	Operator string `json:"operator" bson:"operator"`
}

// ResourceName 和 CollectionName
func (m *Model) ResourceName() string {
	return modelName
}

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
