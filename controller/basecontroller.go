package controller

import (
	"net/http"
)

type RouteHandle struct {
}

type BaseController struct {
	W http.ResponseWriter
	R *http.Request
}

var callbackFuncs = make(map[string]func(http.ResponseWriter, *http.Request))

// map保存回调
func (handle *RouteHandle) Router(relativePath string, handler func(http.ResponseWriter, *http.Request)) {
	callbackFuncs[relativePath] = handler
}

// httpHandle必须要实现的接口方法，请求入口
func (handle *RouteHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// 过滤谷歌浏览器默认ico请求
	if r.URL.Path == "/favicon.ico" {
		return
	}
	path := r.URL.Path

	if call, ok := callbackFuncs[path]; ok {
		// 实际上转发的逻辑
		call(w, r)
		return
	}
	http.Error(w, "error URL:"+r.URL.String(), http.StatusNotFound)
}
