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

	// User
	e.POST("/sign-up", user.CreateUser)                                        // Sign-up
	e.POST("/login", user.Login)                                               // Login
	e.PUT("/update", user.UpdateUser, middleware.JWT([]byte("jwt-secret")))    // Update User Info
	e.DELETE("/delete", user.DeleteUser, middleware.JWT([]byte("jwt-secret"))) // Delete user
	e.POST("/info", user.Info, middleware.JWT([]byte("jwt-secret")))           // User info

	// Log time, ip, host, method, uri, response status, error, and latency
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}","ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}",` +
			`"latency":"${latency_human}"}` + "\n",
	}))
	e.Logger.Fatal(e.Start(":1323"))
}
