package goods

import "github.com/gin-gonic/gin"

func GetGoods(c *gin.Context)  {
	c.String(200, "--GetGoods--")
}