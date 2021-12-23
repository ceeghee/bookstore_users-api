package app

import "github.com/ceeghee/bookstore_users-api/controllers"

func MapUrls() {
	router.GET("/ping", controllers.Ping)

	router.GET("/users/:user_id", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)
}
