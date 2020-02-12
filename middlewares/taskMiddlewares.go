package middlewares

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 简单鉴权
func AuthToken(ctx *gin.Context) {
	token := ctx.Query("token")
	if token == "" {
		ctx.JSON(401, gin.H{"code": 401, "data": false, "msg": "token不能为空"})
		ctx.Abort()
		//return
	} else if token != getRealToken(ctx) {
		ctx.JSON(403, gin.H{"code": 403, "data": false, "msg": "token验证失败"})
		ctx.Abort()
		//return
	} else {
		ctx.Next()
	}
}

// 32位小写
func getRealToken(ctx *gin.Context) string {
	// /task/lists/wechatGoWithPHP
	str := ctx.Request.URL.Path + "/wechatGoWithPHP"
	md5Byte := md5.Sum([]byte(str))
	md5Str := fmt.Sprintf("%x", md5Byte)
	return md5Str
}
