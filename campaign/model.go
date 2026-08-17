package campaign

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "active_"
	collectionNameSuffix = "_config"
	modelName            = "campaign"
)

// 活动类型：发券 / 满减 / 折扣 / 新人 / 限时特价 / 买赠 / 免配送 / 充值赠
const (
	TypeCoupon    = "coupon"
	TypeFullReduce = "full_reduce"
	TypeDiscount  = "discount"
	TypeNewcomer  = "newcomer"
	TypeFlash     = "flash"
	TypeGift      = "gift"
	TypeDelivery  = "delivery"
	TypeRecharge  = "recharge"
)

const (
	StatusDraft   = "draft"
	StatusOnline  = "online"
	StatusOffline = "offline"
)

// Rule 活动优惠规则（金额单位：元）
type Rule struct {
	Threshold     float64 `json:"threshold"`
	Reduce        float64 `json:"reduce"`
	Discount      float64 `json:"discount"`
	MaxReduce     float64 `json:"max_reduce" bson:"max_reduce"`
	GiftProductID string  `json:"gift_product_id" bson:"gift_product_id"`
	GiftQty       int     `json:"gift_qty" bson:"gift_qty"`
	DeliveryReduce float64 `json:"delivery_reduce" bson:"delivery_reduce"`
	RechargeAmount float64 `json:"recharge_amount" bson:"recharge_amount"`
	RechargeGift   float64 `json:"recharge_gift" bson:"recharge_gift"`
}

// Audience 投放人群
type Audience struct {
	Scope    string `json:"scope"` // all|new|level
	LevelIDs []int  `json:"level_ids" bson:"level_ids"`
	NewDays  int    `json:"new_days" bson:"new_days"`
}

// Model 营销活动
type Model struct {
	model.Model `json:"_" bson:"_"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Type        string             `json:"type"`
	Status      string             `json:"status"`
	Cover       string             `json:"cover"`
	Content     string             `json:"content"`
	StartAt     int64              `json:"start_at" bson:"start_at"`
	EndAt       int64              `json:"end_at" bson:"end_at"`
	StoreIDs    []string           `json:"store_ids" bson:"store_ids"`
	ProductIDs  []string           `json:"product_ids" bson:"product_ids"`
	CouponID    string             `json:"coupon_id" bson:"coupon_id"`
	Rule        Rule               `json:"rule" bson:"rule"`
	Audience    Audience           `json:"audience" bson:"audience"`
	Stock       int                `json:"stock"`
	Taken       int                `json:"taken"`
	Priority    int                `json:"priority"`
	Sort        int                `json:"sort"`
}

func (m *Model) ResourceName() string { return modelName }

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
