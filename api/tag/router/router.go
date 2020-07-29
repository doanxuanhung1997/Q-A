package router

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) {
	tag := app.Group("/tag")
	{
		tag.POST("/create", controllers.CreateTag())
		tag.GET("/list-tag", controllers.GetListTag())
		//client.DELETE("/story/:id", controllers.Delete)
	}
}
