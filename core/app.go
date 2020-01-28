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



type HandlerFunc interface {

}

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
	if handler, ok := this.router.handler[key]; ok {
		//AppPrint(*handler)
		vf := reflect.ValueOf(handler)
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
		//handler(newContext(w,req))
	} else {
		// TODO pathinfo router
		//vf := reflect.ValueOf(handler)
		//if path == "/" {
		//	path = "Index"
		//}else {
		//	pathArr := strings.Split(path,"/")
		//	if len(pathArr) < 2 {
		//		fmt.Fprintf(w, "METHOD ROUTER NOT FOUND: %s\n", req.URL)
		//		return
		//	}
		//	path = pathArr[1]
		//}
		//println(path)
		//method := vf.MethodByName(path)
		//// 没有映射出值  说明方法不存在
		//if method == (reflect.Value{}) {
		//	fmt.Fprintf(w, "METHOD NOT FOUND: %s\n", req.URL)
		//}else {
			fmt.Fprintf(w, "ROUTER NOT FOUND: %s\n", req.URL)
		//}
	}
}

// 添加路由
func (this * app) AddRouter(method string, pathUrl string, handler HandlerFunc) {
	this.router.AddRouter(method, pathUrl, handler)
}

// get请求
func (this * app) GET(pathUrl string, handler HandlerFunc) {
	this.router.AddRouter("GET", pathUrl, handler)
}

// post请求
func (this * app) POST(pathUrl string, handler HandlerFunc) {
	this.router.AddRouter("POST", pathUrl, handler)
}







