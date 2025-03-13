package order

import (
	"github.com/open4go/model"
	"github.com/open4go/req5rsp/cst"
	"github.com/r2day/f9z/pay"
	"github.com/r2day/f9z/property"
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
	// 订单变更时间
	Workflow []WorkflowInfo `json:"workflow" bson:"workflow"`
	// Refund 退款申请
	Refund []*RefundApply `json:"refund" bson:"refund"`
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
	// 当前排队号
	CurrentPos int64 `json:"current_pos" bson:"-"`
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
	// 待支付金额(单位：元）这个金额是可能包含小数点
	RealPrice float64 `json:"real_price" bson:"real_price"`
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
	// 	ClosedAt 关闭订单时间
	ClosedAt int64 `json:"closed_at" bson:"closed_at"`
}

// CustomerInfo 客户信息
type CustomerInfo struct {
	Mobile  string `json:"mobile" bson:"mobile"`
	Name    string `json:"name" bson:"name"`
	Id      string `json:"id" bson:"id"`
	Account string `json:"account" bson:"account"`
	Avatar  string `json:"avatar" bson:"avatar"`
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
	// 地理位置(可以快速的让用户进行订单导航）
	// 以及外卖订单的发布
	LBS LBSInfo `json:"lbs" bson:"lbs"`
}

// LBSInfo 地理位置信息
type LBSInfo struct {
	Longitude float64 `json:"longitude"  bson:"longitude"`
	Latitude  float64 `json:"latitude"  bson:"latitude"`
}

// Buckets 商品信息
type Buckets struct {
	// 商品id
	ID string `json:"id"`
	// 名称
	Name string `json:"name"`
	// 数量
	Number int `json:"number"  bson:"number"`
	// 单价
	Price float64 `json:"price"  bson:"price"`
	// 原始价格
	OriginAmount string `json:"origin_amount"  bson:"origin_amount"`
	// 单位
	Unit string `json:"unit"  bson:"unit"`
	// 属性列表
	Property []property.Model `json:"property"  bson:"property"`
	// 以下的id可以快速的通过查找property表里的价格这个id对应的价格进行统计
	PropsID []int `json:"props"  bson:"props"`
	// 单品属性简述
	PropsText string `json:"props_text"  bson:"props_text"`
	// 图片
	Image string `json:"image"  bson:"image"`
	// 属性参数
	PropsItem []PropsItemInfo `json:"props_item"  bson:"-"`
}

// PendingOrder 数据展示
type PendingOrder struct {
	// 订单来源(系统根据订单来源终端自动赋值）
	Id       string       `json:"id"`
	Price    float64      `json:"price"`
	Number   int          `json:"number"`
	Time     int64        `json:"time"`
	Customer CustomerInfo `json:"customer"`
	Merchant MerchantInfo `json:"merchant"`
}

// WorkflowInfo 订单操作记录
type WorkflowInfo struct {
	// Label 标签
	Label string `json:"label" bson:"label"`
	// Description 描述
	Description string `json:"description" bson:"description"`
	// Operator 操作人
	Operator string `json:"operator" bson:"operator"`
	// OperatorID 操作人ID
	OperatorID string `json:"operator_id" bson:"operator_id"`
}

type PropsItemInfo struct {
	IsDefault int    `json:"is_default"`
	Id        int    `json:"id"`
	Code      string `json:"code"`
	Value     string `json:"value"`
	Price     int    `json:"price"`
}

// RefundApply 退款申请记录
type RefundApply struct {
	ApplyId string           `json:"id"`     // 申请单id
	Type    RefundReasonType `json:"type"`   // 申请单类型：1取消订单 2退款
	Status  RefundStatus     `json:"status"` // 申请状态:  0:待处理 1:已审批 2:已完成
	Audit   string           `json:"audit"`  // 审批类型:  1:同意 2:退款
	Reason  string           `json:"reason"` // 原因
	// Amount 一般情况下取消订单可以直接全额退款
	Amount int64 `json:"amount"` // 退款金额
	// OperatorID 操作人
	Operator string `json:"operator"` // 处理人
	// OperatorID 操作人ID
	OperatorID  string     `json:"operator_id" bson:"operator_id"`
	Applier     string     `json:"applier"`                          // 申请发起人id, 用户或商户
	ApplierType RefundType `json:"applier_type" bson:"applier_type"` // 发起人的用户类型 1用户 2商户
	// 	CreatedAt 创建时间
	CreatedAt int64 `json:"created_at" bson:"created_at"`
}

// PendingApply 返回待处理申请
func (m *Model) PendingApply() *RefundApply {
	for _, apply := range m.Refund {
		if apply.Status == RefundProcessed { // Status 为 1 表示已经同意，可以进行申请退款
			return apply
		}
	}
	return nil // 如果没有待处理申请，返回 nil
}

// UpdateApply 返回待处理申请
// newList := m.UpdateApply(target)
// m.Refund = newList
func (m *Model) UpdateApply(target *RefundApply) []*RefundApply {
	newList := make([]*RefundApply, 0)
	for _, apply := range m.Refund {
		if apply.ApplyId == target.ApplyId { // Status 为 0 表示待处理
			newList = append(newList, target)
		} else {
			// 保持原来的列表不变
			newList = append(newList, apply)
		}
	}
	return newList // 如果没有待处理申请，返回 nil
}

// ResourceName 返回资源名称
func (m *Model) ResourceName() string {
	return modelName
}

// CollectionName 返回表名称
func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
