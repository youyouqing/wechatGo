package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/robfig/cron/v3"
	//. "wechatGin.cthai.cn/middlewares"
	. "wechatGin.cthai.cn/service"
)

func main() {

	router := gin.Default()

	//微信api相关
	wechatApiGroup := router.Group("/wechat")
	{
		wechatApiGroup.GET("/users", func(context *gin.Context) {
			context.String(200, "dzc")
		})
	}
	// 任务api相关
	taskApiGroup := router.Group("/task")
	{
		// 加入鉴权中间件
		//taskApiGroup.Use(AuthToken)
		taskApiGroup.GET("/add", HandleTaskAdd)
		taskApiGroup.GET("/addAndStart", HandleTaskAddAndStart)
		taskApiGroup.GET("/start/:id", HandleTaskStart)
		taskApiGroup.GET("/stop/:id", HandleTaskStop)
	}

	config := ShareConfigInstance(false)
	serverPort := config.GetConfigFromKey("server_port")
	router.Run(serverPort)
}
