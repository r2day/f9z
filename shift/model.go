package shift

import (
	"github.com/open4go/model"
	"github.com/open4go/req5rsp/cst"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionNamePrefix 数据库表前缀
	// 可以根据具体业务的需要进行定义
	// 例如: sys_, scm_, customer_, order_ 等
	collectionNamePrefix = "finance_"
	// CollectionNameSuffix 后缀
	// 例如, _log, _config, _flow,
	collectionNameSuffix = "_log"
	// 这个需要用户根据具体业务完成设定
	modelName = "shift"
)

type MenuItem struct {
	ID       string  `json:"id" bson:"id,omitempty"`
	Name     string  `json:"name" bson:"name,omitempty"`
	Price    float64 `json:"price" bson:"price,omitempty"`
	Quantity int     `json:"quantity" bson:"quantity,omitempty"`
}

type Order struct {
	OrderID       string          `json:"order_id" bson:"order_id,omitempty"`
	TableID       *string         `json:"table_id" bson:"table_id,omitempty"`
	Items         []MenuItem      `json:"items" bson:"items,omitempty"`
	Status        cst.OrderStatus `json:"status" bson:"status,omitempty"`
	TotalAmount   float64         `json:"total_amount" bson:"total_amount,omitempty"`
	PaidAmount    float64         `json:"paid_amount" bson:"paid_amount,omitempty"`
	PaymentMethod string          `json:"payment_method" bson:"payment_method,omitempty"`
	CreateTime    string          `json:"create_time" bson:"create_time,omitempty"`
}

type Model struct {
	// 模型继承
	model.Model `json:"_" bson:"_"`
	// 基本的数据库模型字段，一般情况所有model都应该包含如下字段
	// 创建时（用户上传的数据为空，所以默认可以不传该值)
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	HandoverID      string             `json:"handover_id" bson:"handover_id,omitempty"`
	StartTime       string             `json:"start_time" bson:"start_time,omitempty"`
	EndTime         string             `json:"end_time" bson:"end_time,omitempty"`
	PreviousCashier string             `json:"previous_cashier" bson:"previous_cashier,omitempty"`
	OperatorId      string             `json:"operator_id" bson:"operator_id,omitempty"`
	NextCashier     string             `json:"next_cashier" bson:"next_cashier,omitempty"`
	Supervisor      string             `json:"supervisor,omitempty" bson:"supervisor,omitempty"`

	TotalOrders      int     `json:"total_orders" bson:"total_orders,omitempty"`
	TotalSalesAmount float64 `json:"total_sales_amount" bson:"total_sales_amount,omitempty"`
	TotalPaidAmount  float64 `json:"total_paid_amount" bson:"total_paid_amount,omitempty"`
	TotalRefund      float64 `json:"total_refund_amount" bson:"total_refund_amount,omitempty"`

	OpeningCash    float64 `json:"opening_cash" bson:"opening_cash,omitempty"`
	ClosingCash    float64 `json:"closing_cash" bson:"closing_cash,omitempty"`
	ExpectedCash   float64 `json:"expected_cash" bson:"expected_cash,omitempty"`
	CashDifference float64 `json:"cash_difference" bson:"cash_difference,omitempty"`

	PaymentSummary  map[string]float64 `json:"payment_summary" bson:"payment_summary,omitempty"`
	OpenOrders      []Order            `json:"open_orders" bson:"open_orders,omitempty"`
	LowStockItems   []string           `json:"low_stock_items" bson:"low_stock_items,omitempty"`
	SpecialNotes    string             `json:"special_notes" bson:"special_notes,omitempty"`
	Anomalies       []string           `json:"anomalies" bson:"anomalies,omitempty"`
	CancelledOrders int                `json:"cancelled_orders" bson:"cancelled_orders,omitempty"`
	RefundedAmount  float64            `json:"refunded_amount" bson:"refunded_amount,omitempty"`
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
