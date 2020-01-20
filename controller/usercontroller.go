package controller

import "net/http"
import "../util"

type UserController struct {

}

// 注册方法路由，必须实现 否则路由不生效
func (p *UserController) Router(router *RouteHandle) {
	router.Router("/hello", p.hello)
	router.Router("/json", p.json)
}

func (p *UserController)hello(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("hello"))
}


func (p *UserController)json(w http.ResponseWriter,r *http.Request)  {
	util.ResultJson(w,[]string{"2"})
}

