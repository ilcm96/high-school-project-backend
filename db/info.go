package db

import (
	"go.mongodb.org/mongo-driver/bson"

	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Info(c echo.Context) (err error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	collection := Client().Database("user").Collection("user")
	defer Client().Disconnect(context.TODO())

	var result User

	filter := bson.M{"id": id}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	name := result.Name

	return c.JSON(http.StatusCreated, map[string]string{"id": id, "name": name})
}
