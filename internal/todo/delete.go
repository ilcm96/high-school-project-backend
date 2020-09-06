package todo

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteTodo(client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get id from jwt
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(string)

		// Set collection
		collection := client.Database("todo").Collection(id)

		// Delete todo
		toDoId := c.Param("id")
		filter := bson.M{"_id": toDoId}
		deleteResult, err := collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			return echo.ErrInternalServerError
		}
		if deleteResult.DeletedCount != 0 {
			return c.JSON(http.StatusCreated, map[string]string{"message": "Delete success"})
		} else {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Todo not exists"})
		}
	}
}
