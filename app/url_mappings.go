package app
import (
	"github.com/titiksha05/book_store_users_api/controllers/ping"
	"github.com/titiksha05/book_store_users_api/controllers/users"
)

func mapUrls(){
	//router.Run( ":8080")
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	//router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}