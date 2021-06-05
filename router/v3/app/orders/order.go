package orders

import "github.com/gin-gonic/gin"

func GetOrders(c *gin.Context)  {
	c.String(200, "--GetOrders--")
}