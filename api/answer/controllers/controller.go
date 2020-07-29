package controllers

import (
	"../../../db"
	"../../../helpers/common"
	"../../../helpers/constant"
	"../../../helpers/middlewares"
	"../model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func CreateAnswer() gin.HandlerFunc {
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
		inputAnswer := model.Answer{}
		err := json.Unmarshal(rawBody, &inputAnswer)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageError001,
			})
			return
		}

		var answer model.Answer
		answer.Id = primitive.NewObjectID()
		answer.QuestionId = inputAnswer.QuestionId
		answer.Answer = inputAnswer.Answer
		answer.TagId = inputAnswer.TagId
		answer.Image = inputAnswer.Image
		answer.CreatedDate = time.Now().UTC()
		answer.CreatedBy = userAuth.Id.Hex()
		answer.UpdatedDate = time.Now().UTC()
		answer.UpdatedBy = userAuth.Id.Hex()
		_, err = db.Collection(model.CollectionAnswer).InsertOne(db.GetContext(), answer)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeSuccess,
				"message": constant.MessageSuccess010,
				"data":    answer,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": err.Error(),
			})
			return
		}
	}
}

func GetAnswerByQuestionId() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, errAuth := middlewares.VerifyToken(c)
		if errAuth != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": errAuth.Error(),
			})
			return
		}
		/*Parameter c.GetRawData*/
		questionId := c.Query("questionId")

		if common.IsEmpty(questionId) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageError002,
			})
			return
		}

		cur, logError := db.Collection(model.CollectionAnswer).Find(db.GetContext(), bson.M{"question_id": questionId})

		if logError != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": logError.Error(),
			})
			return
		}
		var listAnswer []model.Answer
		for cur.Next(db.GetContext()) {
			var answer model.Answer
			err := cur.Decode(&answer)
			if err != nil {
				println("Error on decoding the document ", err.Error())
			}
			listAnswer = append(listAnswer, answer)
		}

		if len(listAnswer) > 0 {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeSuccess,
				"message": constant.MessageSuccess011,
				"data":    listAnswer,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeSuccess,
				"message": constant.MessageSuccess013,
				"data":    listAnswer,
			})
		}
	}
}
