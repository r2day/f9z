package coupon

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "active_"
	collectionNameSuffix = "_config"
	modelName            = "coupon"
)

const (
	TypeCash     = "cash"
	TypePercent  = "percent"
	TypeGift     = "gift"
	TypeDelivery = "delivery"
)

const (
	StatusActive   = "active"
	StatusDisabled = "disabled"
)

// Model 优惠券模板
type Model struct {
	model.Model  `json:"_" bson:"_"`
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name"`
	Type         string             `json:"type"`
	Threshold    float64            `json:"threshold"`
	Amount       float64            `json:"amount"`
	ValidDays    int                `json:"valid_days" bson:"valid_days"`
	StartAt      int64              `json:"start_at" bson:"start_at"`
	EndAt        int64              `json:"end_at" bson:"end_at"`
	Stock        int                `json:"stock"`
	Issued       int                `json:"issued"`
	PerUserLimit int                `json:"per_user_limit" bson:"per_user_limit"`
	Overlay      bool               `json:"overlay"`
	Status       string             `json:"status"`
	Content      string             `json:"content"`
	Cover        string             `json:"cover"`
	GiftProductID string            `json:"gift_product_id" bson:"gift_product_id"`
}

func (m *Model) ResourceName() string { return modelName }

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
