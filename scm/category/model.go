package category

import (
	"github.com/open4go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collectionNamePrefix = "scm_"
	collectionNameSuffix = "_config"
	modelName            = "categories"
)

type Model struct {
	model.Model `json:"_" bson:"_"`
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Sort        int                `json:"sort"`
	Icon        string             `json:"icon"`
	Parent      string             `json:"parent,omitempty"`
}

func (m *Model) ResourceName() string { return modelName }

func (m *Model) CollectionName() string {
	return collectionNamePrefix + modelName + collectionNameSuffix
}
