package stocktxn

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "scm_"
	collectionNameSuffix = "_flow"
	modelName            = "stock_txn"
)

type Model struct {
	model.Model    `json:"_" bson:"_"`
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StoreID        string             `json:"store_id" bson:"store_id"`
	MaterialID     string             `json:"material_id" bson:"material_id"`
	Type           string             `json:"type"`
	Qty            float64            `json:"qty"`
	UnitCost       float64            `json:"unit_cost" bson:"unit_cost"`
	RefType        string             `json:"ref_type" bson:"ref_type"`
	RefID          string             `json:"ref_id" bson:"ref_id"`
	IdempotencyKey string             `json:"idempotency_key" bson:"idempotency_key"`
	Remark         string             `json:"remark"`
}

func (m *Model) ResourceName() string { return modelName }

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
