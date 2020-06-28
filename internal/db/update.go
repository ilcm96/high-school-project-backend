package db

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ilcm96/high-school-auth-backend/internal/model"

	"github.com/ilcm96/high-school-auth-backend/internal/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func structToMap(update *model.User) map[string]string {
	// Make map
	var updateMap map[string]string
	j, _ := json.Marshal(update)
	_ = json.Unmarshal(j, &updateMap)

	// Delete if value does not exists
	for key, val := range updateMap {
		if val == "" {
			delete(updateMap, key)
		}
	}
	return updateMap
}

func UpdateUser(c echo.Context) (err error) {
	collection := Client().Database("user").Collection("user")
	defer Client().Disconnect(context.TODO())

	// Get ID from token
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := claims["id"].(string)

	// Request Body
	u := new(model.User)
	u.ID = ID
	if err = c.Bind(u); err != nil {
		return echo.ErrInternalServerError
	}

	// Hash PW
	u.PW = util.HashPW(u.PW)

	// Update user info
	filter := bson.M{"id": u.ID}
	update := bson.M{
		"$set": structToMap(u),
	}
	updateResult, _ := collection.UpdateOne(context.TODO(), filter, update)

	if updateResult.ModifiedCount != 0 {
		return c.JSON(http.StatusCreated, map[string]string{"message": "Update success"})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Provide info for update"})
	}
}
