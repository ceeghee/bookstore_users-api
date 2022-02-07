package app

import "github.com/ceeghee/bookstore_users-api/controllers"

func MapUserRoutes() {
	userRoutes.GET("/:user_id", controllers.Get)
	userRoutes.POST("/users", controllers.CreateUser)
	userRoutes.POST("/login", controllers.Login)
}
