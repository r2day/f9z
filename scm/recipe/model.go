package recipe

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "scm_"
	collectionNameSuffix = "_config"
	modelName            = "recipe"
)

type Line struct {
	MaterialID string  `json:"material_id" bson:"material_id"`
	Qty        float64 `json:"qty"`
	Unit       string  `json:"unit"`
	WasteRate  float64 `json:"waste_rate" bson:"waste_rate"`
	Note       string  `json:"note"`
}

type VariantMatch struct {
	PropertyID string `json:"property_id" bson:"property_id"`
	ValueID    int    `json:"value_id" bson:"value_id"`
}

type Variant struct {
	Name  string         `json:"name"`
	Mode  string         `json:"mode"`
	Match []VariantMatch `json:"match"`
	Lines []Line         `json:"lines"`
}

type Model struct {
	model.Model `json:"_" bson:"_"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProductID   string             `json:"product_id" bson:"product_id"`
	Name        string             `json:"name"`
	Version     int                `json:"version"`
	Status      string             `json:"status"`
	Producer    string             `json:"producer"`
	YieldQty    float64            `json:"yield_qty" bson:"yield_qty"`
	YieldUnit   string             `json:"yield_unit" bson:"yield_unit"`
	BaseLines   []Line             `json:"base_lines" bson:"base_lines"`
	Variants    []Variant          `json:"variants" bson:"variants"`
	Remark      string             `json:"remark"`
}

func (m *Model) ResourceName() string { return modelName }

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
