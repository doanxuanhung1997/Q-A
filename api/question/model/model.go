package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionQuestion = "question"
)

/*Question Model*/
type Question struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Content     string             `json:"content" bson:"content"`
	TagId       string             `json:"tagId" bson:"tag_id"`
	Like        int                `json:"like" bson:"like"`
	DisLike     int                `json:"dis_like" bson:"dis_like"`
	Image       string             `json:"image" bson:"image"`
	CreatedBy   string             `json:"createdBy" bson:"created_by"`
	CreatedDate time.Time          `json:"createdDate" bson:"created_date"`
	UpdatedBy   string             `json:"updatedBy" bson:"updated_by"`
	UpdatedDate time.Time          `json:"updatedDate" bson:"updated_date"`
}
