package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionTag      = "tag"
)

/*Tag Model*/
type Tag struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Tag         string             `json:"tag" bson:"tag"`
	Image       string             `json:"image" bson:"image"`
	CreatedBy   string             `json:"createdBy" bson:"created_by"`
	CreatedDate time.Time          `json:"createdDate" bson:"created_date"`
	UpdatedBy   string             `json:"updatedBy" bson:"updated_by"`
	UpdatedDate time.Time          `json:"updatedDate" bson:"updated_date"`
}
