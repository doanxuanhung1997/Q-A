package router

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) {
	answer := app.Group("/answer")
	{
		answer.POST("/create", controllers.CreateAnswer())
		answer.GET("/get-answer-by-question-id", controllers.GetAnswerByQuestionId())
	}
}
