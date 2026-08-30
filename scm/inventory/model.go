package inventory

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "scm_"
	collectionNameSuffix = "_stock"
	modelName            = "inventory"
)

type Model struct {
	model.Model `json:"_" bson:"_"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StoreID     string             `json:"store_id" bson:"store_id"`
	MaterialID  string             `json:"material_id" bson:"material_id"`
	Qty         float64            `json:"qty"`
	AvgCost     float64            `json:"avg_cost" bson:"avg_cost"`
}

func (m *Model) ResourceName() string { return modelName }

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
