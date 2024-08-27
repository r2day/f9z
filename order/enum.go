package order

// PickUpType 取餐方式
type PickUpType int

// SourceType 订单来源
type SourceType int

// RefundType 订单来源
type RefundType int

// RefundReasonType 退款原因
type RefundReasonType int

// RefundStatus 退款申请单处理状态
type RefundStatus int

// RefundAudit 退款审计
type RefundAudit int

const (
	// HelpYourSelf 自提
	HelpYourSelf PickUpType = iota
	// Takeout 外卖
	Takeout
	// DineIn 店内就餐
	DineIn
	// Express 快递
	Express
	// NoNeed 无需取货（虚拟商品，或补差价）
	NoNeed
)

const (
	// WxMiniApp 微信小程序
	WxMiniApp SourceType = iota
	// DouYinMiniApp 抖音小程序
	DouYinMiniApp
	// AlipayMiniApp 支付宝小程序
	AlipayMiniApp
	// MeiTuanMiniApp 美团小程序
	MeiTuanMiniApp
	// PosClient 点餐系统
	PosClient
	// WebClient 网页系统
	WebClient
)

const (
	// InvalidApply 无效申请
	InvalidApply RefundType = iota
	// MerchantApply 商户申请
	MerchantApply
	// CustomerApply 客户申请
	CustomerApply
)

const (
	// InvalidRefundReasonApply 无效申请
	InvalidRefundReasonApply RefundReasonType = iota
	// RefundReasonCancelOrder 取消订单
	RefundReasonCancelOrder
	// RefundReasonRefundOrder 退定商品
	RefundReasonRefundOrder
)

const (
	// RefundApplying 申请退款中
	RefundApplying RefundStatus = iota
	// RefundProcessed 退款已处理（操作人进行处理）
	RefundProcessed
	// RefundDone 退款完成 （回调通知收到退款回调）
	RefundDone
)

//const (
//	// RefundAuditA 申请退款中
//	RefundAuditA RefundAudit = iota
//	// RefundProcessed 退款已处理（操作人进行处理）
//)
