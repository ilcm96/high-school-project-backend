package web

import (
	"github.com/ilcm96/crud-jwt/db"
	"go.mongodb.org/mongo-driver/bson"

	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID   string `json:"id"`
	PW   string `json:"pw"`
	Name string `json:"name"`
}

func Info(c echo.Context) (err error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	collection := db.Client().Database("user").Collection("user")
	defer db.Client().Disconnect(context.TODO())

	var result User

	filter := bson.D{{"id", id}}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	name := result.Name

	return c.JSON(http.StatusCreated, map[string]string{"id": id, "name": name})
}
