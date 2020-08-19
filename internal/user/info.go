package user

import (
	"github.com/ilcm96/high-school-auth-backend/internal/db"
	"github.com/ilcm96/high-school-auth-backend/internal/model"
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

	collection := db.Client().Database("user").Collection("user")
	defer db.Client().Disconnect(context.TODO())

	var result model.User

	filter := bson.M{"id": id}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	name := result.Name

	return c.JSON(http.StatusCreated, map[string]string{"id": id, "name": name})
}
