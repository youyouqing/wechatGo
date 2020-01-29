package core

import (
	"github.com/wechatGo/controller"
	"github.com/wechatGo/utils"
)

type router struct {
	handlerFuncs map[string]HandlerFunc
	handlerControllers map[string]HandlerController
}

// 手动路由配置方法
func init() {
	app := ShareAppInstance()

	// 控制器路由
	app.AddGetControllerRouter("/", &controller.UserController{})

	// 请求回调路由（优先）
	app.AddGetFuncRouter("/dzc", func(ctx *utils.Context) {
		println(ctx.Query("name"))
	})
}

func NewRouter() *router {
	return &router{handlerControllers:make(map[string]HandlerController),handlerFuncs:make(map[string]HandlerFunc)}
}

// 添加路由
func (this *router) AddControllerRouter(method string, pathUrl string, handler HandlerController)  {
	key := method + "-" + pathUrl
	if len(key) < 2 {
		panic("addRouter null key")
	}
	// 更新handler
	this.handlerControllers[key] = handler
}

func (this *router) AddFuncRouter(method string, pathUrl string, handler HandlerFunc)  {
	key := method + "-" + pathUrl
	if len(key) < 2 {
		panic("addRouter null key")
	}
	// 更新handler
	this.handlerFuncs[key] = handler
}