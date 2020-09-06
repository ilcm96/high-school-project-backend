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

func GetAllTodo(client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get id from jwt
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(string)

		// Set collection
		collection := client.Database("todo").Collection(id)

		// Get all todo from DB
		var respBody []*model.Todo
		cur, err := collection.Find(context.TODO(), bson.D{{}})
		if err != nil {
			return echo.ErrInternalServerError
		}
		for cur.Next(context.TODO()) {
			var element model.Todo
			if err := cur.Decode(&element); err != nil {
				return echo.ErrInternalServerError
			}
			respBody = append(respBody, &element)
		}
		cur.Close(context.TODO())

		if respBody != nil {
			return c.JSON(http.StatusOK, map[string][]*model.Todo{"message": respBody})
		}
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Todo not exists"})
	}
}

func GetTodo(client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get id from jwt
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["id"].(string)

		// Set collection
		collection := client.Database("todo").Collection(userId)

		// Get single todo
		var result model.Todo
		toDoId := c.Param("id")
		filter := bson.M{"_id": toDoId}
		err := collection.FindOne(context.TODO(), filter).Decode(&result)
		if err == mongo.ErrNoDocuments {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Todo not exists"})
		} else if err != mongo.ErrNoDocuments && err != nil {
			return echo.ErrInternalServerError
		}
		return c.JSON(http.StatusOK, result)
	}
}
