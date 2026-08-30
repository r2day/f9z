package outbound

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "scm_"
	collectionNameSuffix = "_flow"
	modelName            = "outbound"
)

type Line struct {
	MaterialID string  `json:"material_id" bson:"material_id"`
	Qty        float64 `json:"qty"`
	Unit       string  `json:"unit"`
	Note       string  `json:"note"`
}

type Model struct {
	model.Model `json:"_" bson:"_"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Code        string             `json:"code"`
	StoreID     string             `json:"store_id" bson:"store_id"`
	Reason      string             `json:"reason"`
	Status      string             `json:"status"`
	Lines       []Line             `json:"lines"`
	PostedAt    int64              `json:"posted_at" bson:"posted_at"`
	Remark      string             `json:"remark"`
}

func (m *Model) ResourceName() string { return modelName }

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
