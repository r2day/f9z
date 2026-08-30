package material

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "scm_"
	collectionNameSuffix = "_config"
	modelName            = "material"
)

type Model struct {
	model.Model       `json:"_" bson:"_"`
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Code              string             `json:"code"`
	Name              string             `json:"name"`
	Kind              string             `json:"kind"`
	CategoryID        string             `json:"category_id" bson:"category_id"`
	BaseUnit          string             `json:"base_unit" bson:"base_unit"`
	PurchaseUnit      string             `json:"purchase_unit" bson:"purchase_unit"`
	PurchaseFactor    float64            `json:"purchase_factor" bson:"purchase_factor"`
	LastCost          float64            `json:"last_cost" bson:"last_cost"`
	AvgCost           float64            `json:"avg_cost" bson:"avg_cost"`
	SafetyQty         float64            `json:"safety_qty" bson:"safety_qty"`
	ShelfLifeDays     int                `json:"shelf_life_days" bson:"shelf_life_days"`
	DefaultSupplierID string             `json:"default_supplier_id" bson:"default_supplier_id"`
	Images            string             `json:"images"`
	Content           string             `json:"content"`
	Status            string             `json:"status"`
	Spec              string             `json:"spec"`
}

func (m *Model) ResourceName() string { return modelName }

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
