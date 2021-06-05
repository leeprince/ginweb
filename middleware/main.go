package main

import (
	"fmt"
	"ginweb-middleware/middleware"
	"github.com/gin-gonic/gin"
)

/*
1. 当在中间件或 handler 中启动新的 Goroutine 时，不能使用原始的上下文，必须使用只读副本。
 */
func main() {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()
	
	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
	
	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())
	
	/********************* 自定义中间件 */
	r.Use(middleware.InitApp())
	r.Use(middleware.Logic())
	/********************* 自定义中间件 - end*/
	
	/***************** 局部路由 */
	// 你可以为每个路由添加任意数量的中间件。
	r.GET("/benchmark", middleware.Auth(), func(c *gin.Context) {
			fmt.Println("/benchmark ...")
			c.String(200, "/benchmark")
		})
	/***************** 局部路由 - end */
	
	/****************** 认证路由组及嵌套路由组 */
	/* v1 := r.Group("/v1", middleware.Auth()) 与 以下两行代码效果完全一直
	v1 := r.Group("/v1")
	v1.Use(middleware.Auth())
	*/
	v1 := r.Group("/v1", middleware.Auth())
	{ // 仅让代码更加规范
		v1.POST("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "/v1/login",
			})
		})
		v1.POST("/submit", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "/v1/submit",
			})
		})
		
		// 嵌套路由组
		testing := v1.Group("testing")
		testing.GET("/analytics", func(c *gin.Context) {
			panic("手动 panic 测试 gin.Recovery() 中间件： /analytics panic")
		})
	}
	/****************** 认证路由组及嵌套路由 - end */
	
	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}
