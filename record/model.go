package record

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "active_"
	collectionNameSuffix = "_flow"
	modelName            = "record"
)

// Model 活动参与 / 核销流水
type Model struct {
	model.Model  `json:"_" bson:"_"`
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CampaignID   string             `json:"campaign_id" bson:"campaign_id"`
	CouponID     string             `json:"coupon_id" bson:"coupon_id"`
	TicketID     string             `json:"ticket_id" bson:"ticket_id"`
	MemberID     string             `json:"member_id" bson:"member_id"`
	OrderID      string             `json:"order_id" bson:"order_id"`
	StoreID      string             `json:"store_id" bson:"store_id"`
	Type         string             `json:"type"`
	Benefit      float64            `json:"benefit"`
	OrderAmount  float64            `json:"order_amount" bson:"order_amount"`
	Remark       string             `json:"remark"`
}

func (m *Model) ResourceName() string { return modelName }

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
