package todo

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/ilcm96/high-school-auth-backend/internal/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateTodo(client *mongo.Client) echo.HandlerFunc {
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
		if todo.Id == "" || todo.Todo == "" {
			return echo.ErrBadRequest
		}

		// Update todo
		filter := bson.M{"_id": todo.Id}
		update := bson.M{
			"$set": bson.M{"todo": todo.Todo},
		}
		updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			return echo.ErrInternalServerError
		}
		if updateResult.ModifiedCount != 0 {
			return c.JSON(http.StatusCreated, map[string]string{"message": "Update success"})
		} else {
			return echo.ErrInternalServerError
		}

	}
}
