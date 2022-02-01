package main

import (
	"fmt"
	"rest-example/api"
	"rest-example/api/controllers"
	mid "rest-example/api/middlewares"
	"rest-example/config"
	"rest-example/models"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Hello World")
	// Inisialisasi Database
	db := config.InitDB()
	// create echo http
	e := echo.New()
	// initial user model
	userModel := models.NewUserModel(db)
	// initial user controller
	newUserController := controllers.NewController(userModel)
	//register API path and controller
	api.RegisterPath(e, newUserController)
	// log middleware
	mid.LogMiddlewares(e)
	// Start the server and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
