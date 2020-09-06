package main

import (
	"context"
	"net/http"
	"runtime"

	"github.com/ilcm96/high-school-auth-backend/internal/todo"
	"github.com/ilcm96/high-school-auth-backend/internal/user"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	key := "jwt-secret"
	clientOptions := options.Client().ApplyURI(mongodbURI())
	client, _ := mongo.Connect(context.TODO(), clientOptions)
	defer client.Disconnect(context.TODO())

	e := echo.New()

	// Main page
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "github.com/ilcm96/high-school-auth-backend")
	})

	// User
	e.POST("/sign-up", user.CreateUser(client))                               // Sign-up
	e.POST("/login", user.Login(client))                                      // Login
	e.PUT("/update", user.UpdateUser(client), middleware.JWT([]byte(key)))    // Update User Info
	e.DELETE("/delete", user.DeleteUser(client), middleware.JWT([]byte(key))) // Delete user
	e.POST("/info", user.Info(client), middleware.JWT([]byte(key)))           // User info

	// Todo
	e.GET("/todo", todo.GetAllTodo(client), middleware.JWT([]byte(key)))        // Get all todo
	e.GET("/todo/:id", todo.GetTodo(client), middleware.JWT([]byte(key)))       // Get single todo
	e.POST("/todo", todo.CreateTodo(client), middleware.JWT([]byte(key)))       // Create all todo
	e.PATCH("/todo", todo.UpdateTodo(client), middleware.JWT([]byte(key)))      // Update todo
	e.DELETE("/todo/:id", todo.DeleteTodo(client), middleware.JWT([]byte(key))) // Delete todo

	// Log time, ip, host, method, uri, response status, error, and latency
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}","ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}",` +
			`"latency":"${latency_human}"}` + "\n",
	}))
	e.Logger.Fatal(e.Start(":1323"))
}

func mongodbURI() string {
	os := runtime.GOOS
	if os == "windows" {
		return "mongodb://0.0.0.0:27017"
	}
	return "mongodb://root:root@db.sub02111041190.generalvcn.oraclevcn.com:27017"
}
