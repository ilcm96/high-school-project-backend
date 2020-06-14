package db

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUser(c echo.Context) (err error) {
	collection := Client().Database("user").Collection("user")
	defer Client().Disconnect(context.TODO())

	// Get ID from token
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	ID := claims["id"].(string)

	// Request Body
	u := new(User)

	// Add ID from token
	u.ID = ID
	if err = c.Bind(u); err != nil {
		return echo.ErrInternalServerError
	}

	if u.PW != "" && u.Name == "" { // Update user PW
		filter := bson.D{{"id", u.ID}}
		update := bson.D{
			{"$set", bson.D{
				{"name", u.Name},
			}},
		}
		_, _ = collection.UpdateOne(context.TODO(), filter, update)

		return c.JSON(http.StatusCreated, map[string]string{"message": "Update success"})
	} else if u.PW == "" && u.Name != "" { // Update user name
		hashedPW := hashPW(u.PW)

		filter := bson.D{{"id", u.ID}}
		update := bson.D{
			{"$set", bson.D{
				{"pw", hashedPW},
			}},
		}
		_, _ = collection.UpdateOne(context.TODO(), filter, update)

		return c.JSON(http.StatusCreated, map[string]string{"message": "Update success"})
	} else if u.Name != "" && u.PW != "" { // Update user PW and name
		hashedPW := hashPW(u.PW)

		filter := bson.D{{"id", u.ID}}
		update := bson.D{
			{"$set", bson.D{
				{"pw", hashedPW},
				{"name", u.Name},
			}},
		}
		_, _ = collection.UpdateOne(context.TODO(), filter, update)

		return c.JSON(http.StatusCreated, map[string]string{"message": "Update success"})
	} else {
		_ = Client().Disconnect(context.TODO())
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Provide PW or name or both"})
	}
}
