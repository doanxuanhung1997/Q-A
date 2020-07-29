package router

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) {
	tag := app.Group("/vote")
	{
		tag.POST("/vote-question", controllers.VoteQuestion())
		tag.POST("/vote-answer", controllers.VoteAnswer())
	}
}
