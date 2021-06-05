package goods

import (
	"ginweb-router-v3/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	routers.RegisterRouter(Routers)
}

func Routers(e *gin.Engine)  {
	e.GET("/getGoods", GetGoods)
}