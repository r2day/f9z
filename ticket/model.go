package ticket

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "active_"
	collectionNameSuffix = "_flow"
	modelName            = "ticket"
)

const (
	StatusUnused  = "unused"
	StatusUsed    = "used"
	StatusExpired = "expired"
	StatusVoid    = "void"
)

// Model 用户领券记录
type Model struct {
	model.Model `json:"_" bson:"_"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Code        string             `json:"code"`
	CouponID    string             `json:"coupon_id" bson:"coupon_id"`
	CampaignID  string             `json:"campaign_id" bson:"campaign_id"`
	MemberID    string             `json:"member_id" bson:"member_id"`
	AccountID   string             `json:"account_id" bson:"account_id"`
	Status      string             `json:"status"`
	Amount      float64            `json:"amount"`
	Threshold   float64            `json:"threshold"`
	CouponType  string             `json:"coupon_type" bson:"coupon_type"`
	ExpireAt    int64              `json:"expire_at" bson:"expire_at"`
	UsedAt      int64              `json:"used_at" bson:"used_at"`
	OrderID     string             `json:"order_id" bson:"order_id"`
}

func (m *Model) ResourceName() string { return modelName }

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
