package main

import (
	"divine-pizzeria-backend/config"
	"divine-pizzeria-backend/router"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	//setup database
	db, err := config.OpenConn()
	if err != nil {
		println("Error opening database connection")
		panic(err)
	}

	// Initialize the router
	router.InitRouter(db)
}
