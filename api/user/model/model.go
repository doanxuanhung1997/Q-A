package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionUser = "user"
)

/*User Model*/
type User struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username     string             `json:"username" bson:"username"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password" bson:"password"`
	RegisterCode string             `json:"registerCode" bson:"register_code"`
	RegisterDate time.Time          `json:"registerDate" bson:"register_date"`
	FirstName    string             `json:"firstName" bson:"first_name"`
	LastName     string             `json:"lastName" bson:"last_name"`
	DateOfBirth  string             `json:"dateOfBirth" bson:"date_of_birth"`
	PhoneNumber  string             `json:"phoneNumber" bson:"phone_number"`
	Address      string             `json:"address" bson:"address"`
	Image        string             `json:"image" bson:"image"`
	CreatedBy    string             `json:"createdBy" bson:"created_by"`
	CreatedDate  time.Time          `json:"createdDate" bson:"created_date"`
	UpdatedBy    string             `json:"updatedBy" bson:"updated_by"`
	UpdatedDate  time.Time          `json:"updatedDate" bson:"updated_date"`
}

/*Data of Struct ResponseDataLogin*/
type ResponseLogin struct {
	Id    string `json:"Id" `
	Email string `json:"email" `
	Token string `json:"token" `
}
