package main

import (
	"net/http"

	"github.com/ilcm96/high-school-auth-backend/internal/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Main page
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "github.com/ilcm96/high-school-auth-backend")
	})

	// Sign-up
	e.POST("/sign-up", user.CreateUser)

	// Login
	e.POST("/login", user.Login)

	// Update User Info
	e.PUT("/update", user.UpdateUser, middleware.JWT([]byte("jwt-secret")))

	// Delete user
	e.DELETE("/delete", user.DeleteUser, middleware.JWT([]byte("jwt-secret")))

	// User info
	e.POST("/info", user.Info, middleware.JWT([]byte("jwt-secret")))

	// Log time, ip, host, method, uri, response status, error, and latency
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}","ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}",` +
			`"latency":"${latency_human}"}` + "\n",
	}))
	e.Logger.Fatal(e.Start(":1323"))
}
