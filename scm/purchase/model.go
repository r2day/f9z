package purchase

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "scm_"
	collectionNameSuffix = "_flow"
	modelName            = "purchase"
)

type Line struct {
	MaterialID string  `json:"material_id" bson:"material_id"`
	Qty        float64 `json:"qty"`
	Unit       string  `json:"unit"`
	UnitCost   float64 `json:"unit_cost" bson:"unit_cost"`
	Note       string  `json:"note"`
}

type Model struct {
	model.Model `json:"_" bson:"_"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Code        string             `json:"code"`
	StoreID     string             `json:"store_id" bson:"store_id"`
	SupplierID  string             `json:"supplier_id" bson:"supplier_id"`
	Status      string             `json:"status"`
	ExpectAt    string             `json:"expect_at" bson:"expect_at"`
	Lines       []Line             `json:"lines"`
	Remark      string             `json:"remark"`
}

func (m *Model) ResourceName() string { return modelName }

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
