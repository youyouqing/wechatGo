package core

import (
	"github.com/wechatGo/controller"
)

type router struct {
	handler map[string]HandlerFunc
}

// 手动路由配置方法
func init() {
	app := ShareAppInstance()
	app.GET("/", &controller.UserController{})
}

func NewRouter() *router {
	return &router{handler:make(map[string]HandlerFunc)}
}

// 添加路由
func (this *router) AddRouter(method string, pathUrl string, handler HandlerFunc)  {
	key := method + "-" + pathUrl
	if len(key) < 2 {
		panic("addRouter null key")
	}
	// 更新handler
	this.handler[key] = handler
}