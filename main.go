package main

import (
	"fmt"
	"hallo-corona/database"
	"hallo-corona/routes"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {

	// load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		panic("Failed to load .env file")
	}

	//declare e as a new object router
	e := echo.New()

	//database initialized and run migrations
	database.DatabaseInit()

	//set static folder
	e.Static("/uploads", "./uploads")

	//set route group
	routes.RouteInit(e.Group("/api/v1"))

	//running server
	port := os.Getenv("PORT")
	fmt.Println("Server running on port " + port)
	e.Logger.Fatal(e.Start("localhost:" + port))

}
