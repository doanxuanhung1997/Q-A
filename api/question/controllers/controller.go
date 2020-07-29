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

func CreateQuestion() gin.HandlerFunc {
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
		inputQuestion := model.Question{}
		err := json.Unmarshal(rawBody, &inputQuestion)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageError001,
			})
			return
		}

		var question model.Question
		question.Id = primitive.NewObjectID()
		question.Content = inputQuestion.Content
		question.TagId = inputQuestion.TagId
		question.Image = inputQuestion.Image
		question.CreatedDate = time.Now().UTC()
		question.CreatedBy = userAuth.Id.Hex()
		question.UpdatedDate = time.Now().UTC()
		question.UpdatedBy = userAuth.Id.Hex()
		_, err = db.Collection(model.CollectionQuestion).InsertOne(db.GetContext(), question)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeSuccess,
				"message": constant.MessageSuccess007,
				"data":    question,
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

func GetQuestion() gin.HandlerFunc {
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
		objectId, _ := primitive.ObjectIDFromHex(questionId)
		var question model.Question
		err := db.Collection(model.CollectionQuestion).FindOne(db.GetContext(), bson.M{"_id": objectId}).Decode(&question)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeSuccess,
				"message": constant.MessageSuccess008,
				"data":    question,
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

func GetListQuestionByTag() gin.HandlerFunc {
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
		tagId := c.Query("tagId")

		if common.IsEmpty(tagId) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageError003,
			})
			return
		}

		cur, logError := db.Collection(model.CollectionQuestion).Find(db.GetContext(), bson.M{"tag": tagId})

		if logError != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": logError.Error(),
			})
			return
		}
		var listQuestion []model.Question
		for cur.Next(db.GetContext()) {
			var question model.Question
			err := cur.Decode(&question)
			if err != nil {
				println("Error on decoding the document ", err.Error())
			}
			listQuestion = append(listQuestion, question)
		}

		if len(listQuestion) > 0 {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeSuccess,
				"message": constant.MessageSuccess009,
				"data":    listQuestion,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageSuccess009,
				"data":    listQuestion,
			})
			return
		}
	}
}
