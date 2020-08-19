package user

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/ilcm96/high-school-auth-backend/internal/db"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteUser(c echo.Context) (err error) {
	// Set collection
	collection := db.Client().Database("user").Collection("user")
	defer db.Client().Disconnect(context.TODO())

	// Get user ID from token
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	// Delete user
	filter := bson.M{"id": id}
	deleteResult, _ := collection.DeleteOne(context.TODO(), filter)

	if deleteResult.DeletedCount != 0 {
		return c.JSON(http.StatusCreated, map[string]string{"message": "Delete success"})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "User not exists"})
	}
}
