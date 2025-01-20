package routers

import "github.com/gin-gonic/gin"

func InitServerRoutes() {
	r := gin.Default()
	InitUserRoutes(r)
	r.Run()
}
