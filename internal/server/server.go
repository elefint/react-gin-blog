package server

import (
	"react-gin-blog/internal/database"
	"react-gin-blog/internal/store"
)

func Start() {
	store.SetDBConnection(database.NewDBOptions())

	router := setRouter()

	// Start listening and serving requests
	router.Run(":8080")
}
