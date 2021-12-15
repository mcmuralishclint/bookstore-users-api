package app

import (
	ping_controller "github.com/mcmuralishclint/bookstore-users-api/controllers/ping"
	users_controller "github.com/mcmuralishclint/bookstore-users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping_controller.Ping)

	router.POST("/users", users_controller.CreateUser)
	router.GET("/users/:user_id", users_controller.GetUser)
	//router.GET("/users/search", controllers.SearchUser)
}
