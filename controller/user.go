package controller

import (
	"github.com/wechatGo/utils"
)

type UserController struct {
	BaseController
}

func (this * UserController)Index(context *utils.Context)  {
	println(2223333)
}

func (this *UserController)TestHandle(contexts interface{})  {
	println(222)
}
