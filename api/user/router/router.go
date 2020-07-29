package router

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(app *gin.Engine) {
	user := app.Group("/user")
	{
		user.POST("/login", controllers.Login())
		user.POST("/create", controllers.CreateUser())
		//user.POST("/logout", controllers.Logout())
	}
}
