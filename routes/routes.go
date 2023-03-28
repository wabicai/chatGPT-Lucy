package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	. "github.com/wabicai/chatgpt-lucy/app/http/controllers"
	"github.com/wabicai/chatgpt-lucy/app/middlewares"
	"github.com/wabicai/chatgpt-lucy/config"
)

var chatController = NewChatController()

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
	// router.POST("/completion", chatController.Completion)
}
