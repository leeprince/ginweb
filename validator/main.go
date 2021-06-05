package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// 根据请求发送的json的数据解析 =》 rpc通信
type Goods struct {
	Id         int
	GoodsName  string `form:"name" json:"name" binding:"required"`
	
	// 自定义验证器：goodsCheckPrice,多个验证器中间不能有空格
	GoodsPrice int    `form:"price" json:"price" binding:"required,goodsCheckPrice"`
	GoodsNum   int    `form:"num" json:"num" binding:"required,gt=5"`
}

func main() {
	r := gin.Default()
	
	/***** 注册自定义验证器 */
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("goodsCheckPrice", goodsCheckPrice)
	}
	
	r.POST("/goods", func(c *gin.Context) {
		var json Goods
		if err := c.ShouldBind(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(), // 右括号新起一行时，当前行必须存在英文逗号
			})
			return
		}

		c.JSON(200, gin.H{
			"code":  "ok",
			"goods": json,
		})
	})
	
	// 默认端口是 8080
	r.Run()
}

/** 定义自定义验证器 */
var goodsCheckPrice validator.Func = func(fl validator.FieldLevel) bool {
	value, _ := fl.Field().Interface().(int) // 是根据验证的字段进行定义
	fmt.Println("field value : ", value)
	
	// 自定义的验证规则 ...
	if value > 100 {
		return false
	}
	
	return true
}
