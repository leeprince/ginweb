package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
	/*********************** 路由参数 */
	// 此 handler 将匹配 /user/john 但不会匹配 /user/ 或者 /user
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	
	// 此 handler 将匹配 /user/john/ 和 /user/john/send
	// 如果没有其他路由匹配 /user/john，它将重定向到 /user/john/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	/*********************** 路由参数 */
	
	/************************ 路由组 */
	
	// 简单的路由组: v1
	v1 := r.Group("/v1")
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
	}
	
	// 简单的路由组: v2
	v2 := r.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "/v2/login",
			})
		})
		v2.POST("/submit", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "/v2/submit",
			})
		})
	}
	/************************ 路由组 - end */
	
	/*********************** 请求参数 */
	r.POST("/login", func(c *gin.Context) {
		name := c.PostForm("name")
		password := c.PostForm("password")
		code := c.Query("code")
		c.JSON(200, gin.H{
			"name":     name,
			"password": password,
			"code":     code,
		})
	})
	/*********************** 请求参数 -end */
	
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	// r.Run(":8100") // 监听并在 0.0.0.0:8100 上启动服务
}
