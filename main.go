package main

import (
	"./api"
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

	//httpHandle 托管到自定义api  TODO实现路由
	http.HandleFunc("/",api.ApiHandle)
	//监听服务
	err := http.ListenAndServe("0.0.0.0:" + config["server_port"],nil)

	if err != nil {
		fmt.Println("create http server error" + err.Error())
	}

}
