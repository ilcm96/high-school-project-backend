package user

import (
	"context"
	"net/http"
	"time"

	"github.com/ilcm96/high-school-auth-backend/internal/model"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckUser(id string, pw string, client *mongo.Client) bool {
	// Set collection
	collection := client.Database("user").Collection("user")
	defer client.Disconnect(context.TODO())

	// Struct for containing user info in DB
	var result model.User

	// Check whether user exists
	filter := bson.M{"id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return false
	}

	// Compare to password stored in DB
	err = bcrypt.CompareHashAndPassword([]byte(result.PW), []byte(pw))
	if err == nil {
		return true
	} else {
		return false
	}
}

func Login(client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Convert request body to User struct
		u := new(model.User)
		if err := c.Bind(u); err != nil {
			return echo.ErrInternalServerError
		}

		if CheckUser(u.ID, u.PW, client) {
			// New jwt token
			token := jwt.New(jwt.SigningMethodHS256)

			// Claims
			claims := token.Claims.(jwt.MapClaims)
			claims["id"] = u.ID
			claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

			// Token signed with secret key
			tokenString, _ := token.SignedString([]byte("jwt-secret"))

			return c.JSON(http.StatusCreated, map[string]string{
				"token": tokenString,
			})
		} else if !CheckUser(u.ID, u.PW, client) {
			return echo.ErrUnauthorized
		} else {
			return echo.ErrInternalServerError
		}
	}
}
