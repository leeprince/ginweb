package routers

import "github.com/gin-gonic/gin"

// 约定自动注册路由的类型
type Router func(*gin.Engine)

// 自动注册路由的切片
var routers = []Router{}

func RegisterRouter(router ...Router)  {
	routers = append(routers, router...)
}

func IniterRouter() *gin.Engine  {
	r := gin.Default()
	for _, router := range routers {
		router(r)
	}
	return r
}