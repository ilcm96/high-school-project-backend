package db

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteUser(c echo.Context) (err error) {
	collection := Client().Database("user").Collection("user")
	defer Client().Disconnect(context.TODO())

	// Get user ID from token
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := claims["id"].(string)

	// Delete user
	filter := bson.M{"id": ID}
	deleteResult, _ := collection.DeleteOne(context.TODO(), filter)

	if deleteResult.DeletedCount != 0 {
		return c.JSON(http.StatusCreated, map[string]string{"message": "Delete success"})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "User not exists"})
	}
}
