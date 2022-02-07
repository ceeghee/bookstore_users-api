package app

import (
	"github.com/ceeghee/bookstore_users-api/controllers"
	"github.com/gin-gonic/gin"
)

var (
	userRoutes *gin.RouterGroup
)

func MapUrls() {
	router.GET("/ping", controllers.Ping)

	router.GET("/users/:user_id", controllers.Get)
	router.POST("/users", controllers.CreateUser)
}

func GroupUrls() {
	v1 := router.Group("/v1")
	userRoutes = v1.Group("/users")
	MapUserRoutes()
}
