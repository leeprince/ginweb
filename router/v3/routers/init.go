package routers

import "github.com/gin-gonic/gin"

// 约定自动注册路由的类型
type Router func(*gin.Engine)

// 自动注册路由的切片
var routers = []Router{}

// 注册路由
func RegisterRouter(router ...Router)  {
	routers = append(routers, router...)
}

// 初始化路由
func IniterRouter() *gin.Engine  {
	r := gin.Default()
	
	// 初始化路由
	for _, router := range routers {
		router(r)
	}
	return r
}