package process

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "scm_"
	collectionNameSuffix = "_flow"
	modelName            = "process"
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
	Status      string             `json:"status"`
	InLines     []Line             `json:"in_lines" bson:"in_lines"`
	OutLines    []Line             `json:"out_lines" bson:"out_lines"`
	PostedAt    int64              `json:"posted_at" bson:"posted_at"`
	Remark      string             `json:"remark"`
}

func (m *Model) ResourceName() string { return modelName }

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
