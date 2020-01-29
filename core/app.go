package core

import (
	"fmt"
	"github.com/wechatGo/utils"
	"net/http"
	"reflect"
	"sync"
)

type app struct {
	router * router
}

// 控制器路由
type HandlerController interface {}
// 请求回调路由
type HandlerFunc func(ctx *utils.Context)


var mu sync.Mutex
var AppInstance *app

func ShareAppInstance() * app  {
	mu.Lock()
	defer mu.Unlock()

	if AppInstance == nil {
		AppInstance = &app{router:NewRouter()}
	}
	return AppInstance
}

// 执行入口
func (this * app)Run()  {
	config := ShareConfigInstance()
	//fmt.Printf("%v",config)
	serverPort := config.GetConfigFromKey("server_port")
	http.ListenAndServe(serverPort, this)
}

// 请求入口
func (this * app) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 请求方法+pathinfo为key  handleFunc为value的map闭包
	method := req.Method
	path := req.URL.Path
	// 过滤谷歌浏览器icon请求
	if path == "/favicon.ico"{
		return
	}
	key := method + "-" + path
	if handlerMap, ok := this.router.handlerFuncs[key]; ok {
		handlerMap(utils.NewContext(w,req))
	}else if handlerMap, ok := this.router.handlerControllers[key]; ok {
		vf := reflect.ValueOf(handlerMap)
		if path == "/" {
			path = "Index"
		}
		method := vf.MethodByName(path)
		//AppPrint(method)
		var args []reflect.Value
		// 拼接数组参数进入方法
		args = append(args, reflect.ValueOf(utils.NewContext(w,req)))
		// 调用方法
		method.Call(args)
	}else {
		fmt.Fprintf(w, "ROUTER NOT FOUND: %s\n", req.URL)
	}
}

// 添加控制器路由
func (this * app) AddControllerRouter(method string, pathUrl string, handler HandlerController) {
	this.router.AddControllerRouter(method, pathUrl, handler)
}

// 添加回调路由
func (this * app) AddFuncRouter(method string, pathUrl string, handler HandlerFunc) {
	this.router.AddFuncRouter(method, pathUrl, handler)
}

// get控制器请求
func (this * app) AddGetControllerRouter(pathUrl string, handler HandlerController) {
	this.router.AddControllerRouter("GET", pathUrl, handler)
}

// post控制器请求
func (this * app) AddPostControllerRouter(pathUrl string, handler HandlerController) {
	this.router.AddControllerRouter("POST", pathUrl, handler)
}

// get回调请求
func (this * app) AddGetFuncRouter(pathUrl string, handler HandlerFunc) {
	this.router.AddFuncRouter("GET", pathUrl, handler)
}

// post回调请求
func (this * app) AddPostFuncRouter(pathUrl string, handler HandlerFunc) {
	this.router.AddFuncRouter("POST", pathUrl, handler)
}







