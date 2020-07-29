package main

import (
	routerAnswer "./api/answer/router"
	routerQuestion "./api/question/router"
	routerTag "./api/tag/router"
	routerUser "./api/user/router"
	routerVote "./api/voted/router"
	"./db"
	"./helpers/config"
	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"
)

func main() {
	// environment
	arg := "dev"
	config.Loads("./" + arg + ".env")
	config.SetEnv(config.EnvData)
	env := config.GetEnvValue()
	app := gin.Default()
	routerUser.InitRouter(app)
	routerAnswer.InitRouter(app)
	routerQuestion.InitRouter(app)
	routerTag.InitRouter(app)
	routerVote.InitRouter(app)
	if err := db.InitDb(); err != nil {
		panic(err)
	}
	app.Use(ginsession.New())
	app.Run(env.Server.Host + ":" + env.Server.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
