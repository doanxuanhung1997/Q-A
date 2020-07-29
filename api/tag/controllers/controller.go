package controllers

import (
	"../../../db"
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

func CreateTag() gin.HandlerFunc {
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
		inputTag := model.Tag{}
		err := json.Unmarshal(rawBody, &inputTag)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageError001,
			})
			return
		}

		var tag model.Tag
		tag.Id = primitive.NewObjectID()
		tag.Tag = inputTag.Tag
		tag.Image = inputTag.Image
		tag.CreatedDate = time.Now().UTC()
		tag.CreatedBy = userAuth.Id.Hex()
		tag.UpdatedDate = time.Now().UTC()
		tag.UpdatedBy = userAuth.Id.Hex()
		_, err = db.Collection(model.CollectionTag).InsertOne(db.GetContext(), tag)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeSuccess,
				"message": constant.MessageSuccess005,
				"data":    tag,
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

func GetListTag() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, errAuth := middlewares.VerifyToken(c)
		if errAuth != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": errAuth.Error(),
			})
			return
		}
		cur, logError := db.Collection(model.CollectionTag).Find(db.GetContext(), bson.M{})

		if logError != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": logError.Error(),
			})
			return
		}
		var listTag []model.Tag
		for cur.Next(db.GetContext()) {
			var tag model.Tag
			err := cur.Decode(&tag)
			if err != nil {
				println("Error on decoding the document ", err.Error())
			}
			listTag = append(listTag, tag)
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    constant.CodeSuccess,
			"message": constant.MessageSuccess006,
			"data":    listTag,
		})
		return
	}
}
