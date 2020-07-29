package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	CollectionQuestionVoted = "question_voted"
	CollectionAnswerVoted   = "answer_voted"
)

/*QuestionVoted Model*/
type QuestionVoted struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserId      string             `json:"userId" bson:"user_id"`
	QuestionId  string             `json:"questionId" bson:"question_id"`
	Voted       int                `json:"voted" bson:"voted"`
	CreatedDate time.Time          `json:"createdDate" bson:"created_date"`
	UpdatedDate time.Time          `json:"updatedDate" bson:"updated_date"`
}

/*AnswerVoted Model*/
type AnswerVoted struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserId      string             `json:"userId" bson:"user_id"`
	AnswerId    string             `json:"answerId" bson:"answer_id"`
	Voted       int                `json:"voted" bson:"voted"`
	CreatedDate time.Time          `json:"createdDate" bson:"created_date"`
	UpdatedDate time.Time          `json:"updatedDate" bson:"updated_date"`
}

/*QuestionVoted Model*/
type DataInputVote struct {
	QuestionId string `json:"questionId"`
	AnswerId   string `json:"answerId"`
	TypeVoted  int    `json:"typeVoted"`
}
