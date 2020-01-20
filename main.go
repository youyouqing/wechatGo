package main

import (
	"./controller"
	"./service"
	"fmt"
	"net/http"
)

func main()  {

	config := service.GetEnvConfig()
	//读取env配置 print test
	for key,value:= range config{
		fmt.Printf("%s=>%s\n",key,value)
	}

	var routeHandle = new(controller.RouteHandle)
	//托管回调 必须在服务开启之前
	RegistRoute(routeHandle)
	//监听服务
	err := http.ListenAndServe("0.0.0.0:" + config["server_port"],routeHandle)
	if err != nil {
		fmt.Println("create http server error" + err.Error())
	}

}

// 注册各个控制器路由
func RegistRoute(handle *controller.RouteHandle)  {
	new(controller.UserController).Router(handle)
}
