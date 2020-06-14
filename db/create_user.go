package db

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(c echo.Context) (err error) {
	collection := Client().Database("user").Collection("user")

	defer Client().Disconnect(context.TODO())

	// Request Body
	u := new(User)
	if err = c.Bind(u); err != nil {
		return echo.ErrInternalServerError
	}

	// Check whether ID already exists
	filter := bson.D{{"id", u.ID}}
	err = collection.FindOne(context.TODO(), filter).Err()
	if err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID already exists"})
	}

	// Hash PW with bcrypt and replace plain PW to hashed PW
	u.PW = hashPW(u.PW)

	// Insert new user to DB
	_, err = collection.InsertOne(context.TODO(), u)
	if err != nil {
		return echo.ErrInternalServerError
	} else {
		return c.JSON(http.StatusCreated, map[string]string{"message": "Sign-up success"})
	}
}
