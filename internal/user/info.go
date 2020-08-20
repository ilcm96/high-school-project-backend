package user

import (
	"github.com/ilcm96/high-school-auth-backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Info(client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(string)

		collection := client.Database("user").Collection("user")
		defer client.Disconnect(context.TODO())

		var result model.User

		filter := bson.M{"id": id}
		err := collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			return echo.ErrInternalServerError
		}
		name := result.Name

		return c.JSON(http.StatusCreated, map[string]string{"id": id, "name": name})
	}
}
