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

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*Parameter c.GetRawData*/
		rawBody, _ := c.GetRawData()
		inputUser := model.User{}
		err := json.Unmarshal(rawBody, &inputUser)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageError001,
			})
			return
		}

		// validate email
		if common.IsEmpty(inputUser.Email) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorEmail001,
			})
			return
		}
		if !common.CheckValidationEmail(inputUser.Email) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorEmail002,
			})
			return
		}
		if !common.CheckLength(inputUser.Email, constant.MinLengthEmail, constant.MaxLengthEmail) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorEmail003,
			})
			return
		}

		if common.CheckEmailExit(inputUser.Email) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorEmail004,
			})
			return
		}

		// validate password
		if common.IsEmpty(inputUser.Password) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorPassword001,
			})
			return
		}
		if !common.CheckLength(inputUser.Password, constant.MinLengthPassword, constant.MaxLengthPassword) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorPassword002,
			})
			return
		}

		/*validate first name*/
		if common.IsEmpty(inputUser.FirstName) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorFirstName001,
			})
			return
		}
		if !common.CheckLength(inputUser.FirstName, constant.MinLengthFirstName, constant.MaxLengthFirstName) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorFirstName002,
			})
			return
		}
		if !common.CheckSpecialCharacters(inputUser.FirstName) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorFirstName003,
			})
			return
		}

		/*Validate last name*/
		if common.IsEmpty(inputUser.LastName) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorLastName001,
			})
			return
		}
		if !common.CheckLength(inputUser.LastName, constant.MinLengthLastName, constant.MaxLengthLastName) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorLastName002,
			})
			return
		}
		if !common.CheckSpecialCharacters(inputUser.LastName) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorLastName003,
			})
			return
		}

		// validate birthday
		if common.IsEmpty(inputUser.DateOfBirth) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorBirthday001,
			})
			return
		}

		if !common.CheckFormatDate(inputUser.DateOfBirth) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorBirthday002,
			})
			return
		}

		// validate phone number
		if common.IsEmpty(inputUser.PhoneNumber) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorPhoneNumber001,
			})
			return
		}
		if !common.CheckIsNumber(inputUser.PhoneNumber) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorPhoneNumber002,
			})
			return
		}

		if !common.CheckLength(inputUser.PhoneNumber, constant.MinLengthPhoneNumber, constant.MaxLengthPhoneNumber) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorPhoneNumber003,
			})
			return
		}

		/*Validate address*/
		if common.IsEmpty(inputUser.Address) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorAddress001,
			})
			return
		}
		if !common.CheckLength(inputUser.Address, constant.MinLengthAddress, constant.MaxLengthAddress) {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageErrorAddress002,
			})
			return
		}

		var user model.User
		user.Id = primitive.NewObjectID()
		user.Email = inputUser.Email
		user.Username = inputUser.Email
		user.FirstName = inputUser.FirstName
		user.LastName = inputUser.LastName
		user.Address = inputUser.Address
		user.PhoneNumber = inputUser.PhoneNumber
		user.DateOfBirth = inputUser.DateOfBirth
		user.RegisterDate = time.Now().UTC()
		user.RegisterCode = common.GenerateTokenString(constant.MaxLenRegisterCode)
		user.Password = common.HashPassword(inputUser.Password)
		user.CreatedDate = time.Now().UTC()
		user.UpdatedDate = time.Now().UTC()
		_, err = db.Collection(model.CollectionUser).InsertOne(db.GetContext(), user)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeSuccess,
				"message": constant.MessageSuccess001,
				"data":    user,
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

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*Parameter c.GetRawData*/
		rawBody, _ := c.GetRawData()
		inputUser := model.User{}
		err := json.Unmarshal(rawBody, &inputUser)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": constant.MessageError001,
			})
			return
		}

		email := inputUser.Username
		password := inputUser.Password

		//Hash password
		hashedPass := common.HashPassword(password)

		var user model.User
		errLogin := db.Collection(model.CollectionUser).FindOne(db.GetContext(), bson.M{"username": email, "password": hashedPass}).Decode(&user)
		if errLogin == nil {
			/*Login success*/
			var token, _ = middlewares.GenerateJWT(user.Username, user.Id.Hex())
			data := model.ResponseLogin{}
			data.Id = user.Id.Hex()
			data.Token = token
			data.Email = user.Email
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeSuccess,
				"message": constant.MessageSuccess002,
				"data":    data,
			})
		} else {
			/*Login fail*/
			c.JSON(http.StatusOK, gin.H{
				"code":    constant.CodeFail,
				"message": errLogin.Error(),
			})
		}
	}
}
//
//func Logout() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		au, err := middlewares.ExtractTokenMetadata(c.Request)
//		if err != nil {
//			c.JSON(http.StatusUnauthorized, "unauthorized")
//			return
//		}
//		deleted, delErr := middlewares.DeleteAuth(c, au.AccessUuid)
//		if delErr != nil || deleted == 0 { //if any goes wrong
//			c.JSON(http.StatusUnauthorized, "unauthorized")
//			return
//		}
//		c.JSON(http.StatusOK, "Successfully logged out")
//	}
//}
