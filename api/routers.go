package api

import (
	"rest-example/api/controllers"
	"rest-example/api/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, uController *controllers.Controller) {
	// ---------------------
	// --- REGISTER USER ---
	// ---------------------
	e.POST("/users/register", uController.RegisterController)
	e.POST("/users/login", uController.LoginController)

	// Midleware Auth
	r := e.Group("")
	r.Use(middleware.JWT([]byte(middlewares.SECRET_JWT)))
	r.PUT("/users", uController.UpdateController)
	r.DELETE("/users/:id", uController.DeleteController)
	r.GET("/users", uController.GetUserController)
}
