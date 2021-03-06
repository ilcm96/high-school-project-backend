package user

import (
	"context"
	"net/http"

	"github.com/ilcm96/high-school-auth-backend/internal/model"
	"github.com/ilcm96/high-school-auth-backend/internal/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Set collection
		collection := client.Database("user").Collection("user")

		// Request Body
		u := new(model.User)
		if err := c.Bind(u); err != nil {
			return echo.ErrInternalServerError
		}

		// Check whether ID already exists
		filter := bson.M{"id": u.ID}
		err := collection.FindOne(context.TODO(), filter).Err()
		if err == nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID already exists"})
		}

		// Hash PW with bcrypt and replace plain PW to hashed PW
		u.PW = util.HashPW(u.PW)

		// Insert new user to DB
		_, err = collection.InsertOne(context.TODO(), u)
		if err != nil {
			return echo.ErrInternalServerError
		} else {
			return c.JSON(http.StatusCreated, map[string]string{"message": "Sign-up success"})
		}
	}
}
