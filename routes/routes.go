package routes

import (
	"chatgpt-lucy/app/http/controllers"
	"chatgpt-lucy/app/middlewares"
	"chatgpt-lucy/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

var chatController = controllers.NewChatController()

// RegisterWebRoutes 注册路由
func RegisterWebRoutes(router *gin.Engine) {

	router.Use(middlewares.Cors())

	cnf := config.LoadConfig()
	fmt.Println(cnf)
	if len(cnf.AuthUser) > 0 {
		router.Use(gin.BasicAuth(gin.Accounts{
			cnf.AuthUser: cnf.AuthPassword,
		}))
	}
	router.GET("/", chatController.Index)
	chatgpt := router.Group("chatgpt")
	{
		chatgpt.POST("/completions", chatController.Completions)
	}
}
