package main

import (
	"net/http"

	"github.com/ilcm96/high-school-project-backend/internal/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Main page
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "github.com/ilcm96/high-school-project-backend")
	})

	// Sign-up
	e.POST("/sign-up", db.CreateUser)

	// Login
	e.POST("/login", db.Login)

	// Update User Info
	e.PUT("/update", db.UpdateUser, middleware.JWT([]byte("jwt-secret")))

	// Delete user
	e.DELETE("/delete", db.DeleteUser, middleware.JWT([]byte("jwt-secret")))

	// User info
	e.POST("/info", db.Info, middleware.JWT([]byte("jwt-secret")))

	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":1323"))
}
