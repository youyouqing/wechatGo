wechatGo
----
**0、优化代码**
**1、添加路由router**
```
// 手动路由配置方法
func init() {
	app := ShareAppInstance()
    //默认走控制器的Index()方法
	app.GET("/", &controller.UserController{})
    //执行控制器UserController的Demo()方法
    app.GET("/Demo", &controller.UserController{})
    
}
```
**2、上下文整理context**
`每一个请求到控制器中都会带该次请求的上下文进来`
```
func (this * UserController)Index(context *utils.Context)  {

    
}
```
