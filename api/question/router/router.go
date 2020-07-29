package router

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) {
	question := app.Group("/question")
	{
		question.POST("/create", controllers.CreateQuestion())
		question.GET("/view", controllers.GetQuestion())
		question.GET("/get-list-question-by-tag", controllers.GetListQuestionByTag())
	}
}
