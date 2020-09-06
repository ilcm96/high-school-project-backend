package todo

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/ilcm96/high-school-auth-backend/internal/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateTodo(client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get id from jwt
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(string)

		// Set collection
		collection := client.Database("todo").Collection(id)

		// Convert request body to Todo struct
		todo := new(model.Todo)
		if err := c.Bind(todo); err != nil {
			return echo.ErrInternalServerError
		}
		todo.Id = primitive.NewObjectID().Hex()

		// Insert Note to DB
		_, err := collection.InsertOne(context.TODO(), todo)
		if err != nil {
			return echo.ErrInternalServerError
		} else {
			return c.JSON(http.StatusCreated, map[string]string{"message": "Todo created"})
		}
	}
}
