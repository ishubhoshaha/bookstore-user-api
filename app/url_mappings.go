package app

import (
	"github.com/ishubhoshaha/bookstore-user-api/controllers/ping"
	"github.com/ishubhoshaha/bookstore-user-api/controllers/users"
)

func mapURL()  {
	router.GET("/ping", ping.Ping)

	//router.GET("/users/search", users.FindUser)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users/create", users.CreateUser)

}