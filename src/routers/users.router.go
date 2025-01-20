package routers

import (
	usersController "crud-golang/src/controllers"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.Engine) {
	api := r.Group("/users")

	api.POST("/", usersController.Create)
	api.PUT("/:id", usersController.Update)
	api.GET("/", usersController.FindAllPaginated)
	api.GET("/:id", usersController.FindById)
	api.DELETE("/:id", usersController.Delete)
}
