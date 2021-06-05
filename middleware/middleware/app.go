package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// gin框架的中间件
// 定位：全局
// 执行：所有路由前
func InitApp() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("InitApp === >>>> start ====>>>>")
		
		c.Next() // 代表需要执行的实际的route -》请求操作
		
		fmt.Println("InitApp === >>>> end   ====>>>>")
	}
}

// gin框架的中间件
// 定位：全局
// 执行：围绕c.Next() 执行 ； 所有路由前 =》处理请求 =》所有路由后
func Logic() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Logic === >>>> start ====>>>>")

		c.Next() // 代表需要执行的实际的route -》请求操作

		fmt.Println("Logic === >>>> end   ====>>>>")
	}
}

// gin框架的中间件
// 定位：局部 -》因使用而设定
// 执行：围绕c.Next() 执行 ； 所有路由前 =》处理请求 =》所有路由后
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Auth === >>>> start ====>>>>")

		c.Next() // 代表需要执行的实际的route -》请求操作

		fmt.Println("Auth === >>>> end   ====>>>>")
	}
}
