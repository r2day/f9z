package supplier

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "scm_"
	collectionNameSuffix = "_config"
	modelName            = "supplier"
)

type Item struct {
	MaterialID string  `json:"material_id" bson:"material_id"`
	SupplySku  string  `json:"supply_sku" bson:"supply_sku"`
	Unit       string  `json:"unit"`
	Price      float64 `json:"price"`
	LeadDays   int     `json:"lead_days" bson:"lead_days"`
}

type Model struct {
	model.Model `json:"_" bson:"_"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Contact     string             `json:"contact"`
	Phone       string             `json:"phone"`
	Address     string             `json:"address"`
	Status      string             `json:"status"`
	Remark      string             `json:"remark"`
	Materials   []Item             `json:"materials" bson:"materials"`
}

func (m *Model) ResourceName() string { return modelName }

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
