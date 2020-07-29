package controllers

import (
	"../../../db"
	"../../../helpers/constant"
	"../../../helpers/middlewares"
	modelAnswer "../../answer/model"
	modelQuestion "../../question/model"
	"../model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func VoteQuestion() gin.HandlerFunc {
	return func(c *gin.Context) {
		userAuth, errAuth := middlewares.VerifyToken(c)
		if errAuth != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": errAuth.Error(),
			})
			return
		}

		/*Parameter c.GetRawData*/
		rawBody, _ := c.GetRawData()
		inputVote := model.DataInputVote{}
		err := json.Unmarshal(rawBody, &inputVote)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageError001,
			})
			return
		}

		var dataQuestion modelQuestion.Question
		objectQuestionId, _ := primitive.ObjectIDFromHex(inputVote.QuestionId)
		errQuestion := db.Collection(modelQuestion.CollectionQuestion).FindOne(db.GetContext(), bson.M{"_id": objectQuestionId}).Decode(&dataQuestion)
		if errQuestion != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageError004,
			})
			return
		}

		var questionVoted model.QuestionVoted
		errVoted := db.Collection(model.CollectionQuestionVoted).FindOne(db.GetContext(), bson.M{"user_id": userAuth.Id.Hex(), "question_id": inputVote.QuestionId}).Decode(&questionVoted)
		if errVoted == nil {
			if constant.TypeVoteLike == inputVote.TypeVoted {
				if questionVoted.Voted != constant.TypeVoteLike {
					questionVoted.Voted = constant.TypeVoteLike
					db.Collection(model.CollectionQuestionVoted).UpdateOne(db.GetContext(), bson.M{"_id": questionVoted.Id}, bson.M{"$set": questionVoted})
					dataQuestion.Like += 1
					dataQuestion.DisLike -= 1
				} else {
					db.Collection(model.CollectionQuestionVoted).DeleteOne(db.GetContext(), bson.M{"_id": questionVoted.Id})
					dataQuestion.Like -= 1
				}
			} else {
				if questionVoted.Voted != constant.TypeVoteDisLike {
					questionVoted.Voted = constant.TypeVoteDisLike
					db.Collection(model.CollectionQuestionVoted).UpdateOne(db.GetContext(), bson.M{"_id": questionVoted.Id}, bson.M{"$set": questionVoted})
					dataQuestion.DisLike += 1
					dataQuestion.Like -= 1
				} else {
					db.Collection(model.CollectionQuestionVoted).DeleteOne(db.GetContext(), bson.M{"_id": questionVoted.Id})
					dataQuestion.DisLike -= 1
				}
			}
			db.Collection(modelQuestion.CollectionQuestion).UpdateOne(db.GetContext(), bson.M{"_id": dataQuestion.Id}, bson.M{"$set": dataQuestion})
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeSuccess,
				"message": constant.MessageSuccess003,
			})
			return
		} else {
			questionVoted.Id = primitive.NewObjectID()
			questionVoted.UserId = userAuth.Id.Hex()
			questionVoted.QuestionId = inputVote.QuestionId
			questionVoted.Voted = inputVote.TypeVoted
			questionVoted.CreatedDate = time.Now().UTC()
			questionVoted.UpdatedDate = time.Now().UTC()

			_, errVoteNew := db.Collection(model.CollectionQuestionVoted).InsertOne(db.GetContext(), questionVoted)
			if inputVote.TypeVoted == constant.TypeVoteLike {
				dataQuestion.Like += 1
			} else {
				dataQuestion.DisLike += 1
			}
			db.Collection(modelQuestion.CollectionQuestion).UpdateOne(db.GetContext(), bson.M{"_id": dataQuestion.Id}, bson.M{"$set": dataQuestion})
			if errVoteNew == nil {
				c.JSON(http.StatusOK, gin.H{
					"code":    constant.CodeSuccess,
					"message": constant.MessageSuccess003,
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code":    constant.CodeFail,
					"message": errVoteNew.Error(),
				})
				return
			}
		}
	}
}

func VoteAnswer() gin.HandlerFunc {
	return func(c *gin.Context) {
		userAuth, errAuth := middlewares.VerifyToken(c)
		if errAuth != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": errAuth.Error(),
			})
			return
		}

		/*Parameter c.GetRawData*/
		rawBody, _ := c.GetRawData()
		inputVote := model.DataInputVote{}
		err := json.Unmarshal(rawBody, &inputVote)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageError001,
			})
			return
		}
		var dataAnswer modelAnswer.Answer
		objectAnswerId, _ := primitive.ObjectIDFromHex(inputVote.AnswerId)
		errQuestion := db.Collection(modelAnswer.CollectionAnswer).FindOne(db.GetContext(), bson.M{"_id": objectAnswerId}).Decode(&dataAnswer)
		if errQuestion != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageError005,
			})
			return
		}

		var answerVoted model.AnswerVoted
		errVoted := db.Collection(model.CollectionAnswerVoted).FindOne(db.GetContext(), bson.M{"user_id": userAuth.Id.Hex(), "answer_id": inputVote.AnswerId}).Decode(&answerVoted)
		if errVoted == nil {
			if constant.TypeVoteLike == inputVote.TypeVoted {
				if answerVoted.Voted != constant.TypeVoteLike {
					answerVoted.Voted = constant.TypeVoteLike
					db.Collection(model.CollectionAnswerVoted).UpdateOne(db.GetContext(), bson.M{"_id": answerVoted.Id}, bson.M{"$set": answerVoted})
					dataAnswer.Like += 1
					dataAnswer.DisLike -= 1
				} else {
					db.Collection(model.CollectionAnswerVoted).DeleteOne(db.GetContext(), bson.M{"_id": answerVoted.Id})
					dataAnswer.Like -= 1
				}
			} else {
				if answerVoted.Voted != constant.TypeVoteDisLike {
					answerVoted.Voted = constant.TypeVoteDisLike
					db.Collection(model.CollectionAnswerVoted).UpdateOne(db.GetContext(), bson.M{"_id": answerVoted.Id}, bson.M{"$set": answerVoted})
					dataAnswer.DisLike += 1
					dataAnswer.Like -= 1
				} else {
					db.Collection(model.CollectionAnswerVoted).DeleteOne(db.GetContext(), bson.M{"_id": answerVoted.Id})
					dataAnswer.DisLike -= 1
				}
			}
			db.Collection(modelAnswer.CollectionAnswer).UpdateOne(db.GetContext(), bson.M{"_id": dataAnswer.Id}, bson.M{"$set": dataAnswer})
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeSuccess,
				"message": constant.MessageSuccess004,
			})
			return
		} else {
			answerVoted.Id = primitive.NewObjectID()
			answerVoted.UserId = userAuth.Id.Hex()
			answerVoted.AnswerId = inputVote.AnswerId
			answerVoted.Voted = inputVote.TypeVoted
			answerVoted.CreatedDate = time.Now().UTC()
			answerVoted.UpdatedDate = time.Now().UTC()

			_, errVoteNew := db.Collection(model.CollectionAnswerVoted).InsertOne(db.GetContext(), answerVoted)
			if inputVote.TypeVoted == constant.TypeVoteLike {
				dataAnswer.Like += 1
			} else {
				dataAnswer.DisLike += 1
			}
			db.Collection(modelAnswer.CollectionAnswer).UpdateOne(db.GetContext(), bson.M{"_id": dataAnswer.Id}, bson.M{"$set": dataAnswer})
			if errVoteNew == nil {
				c.JSON(http.StatusOK, gin.H{
					"code":    constant.CodeSuccess,
					"message": constant.MessageSuccess004,
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code":    constant.CodeFail,
					"message": errVoteNew.Error(),
				})
				return
			}
		}
	}
}
