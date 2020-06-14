package db

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckUser(id, pw string) bool {
	collection := Client().Database("user").Collection("user")
	defer Client().Disconnect(context.TODO())

	// Struct for containing user info in DB
	var result User

	// Check whether user exists
	filter := bson.D{{"id", id}}
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

func Login(c echo.Context) (err error) {
	// Request body
	u := new(User)
	if err = c.Bind(u); err != nil {
		return echo.ErrInternalServerError
	}

	if CheckUser(u.ID, u.PW) {
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
	} else if !CheckUser(u.ID, u.PW) {
		return echo.ErrUnauthorized
	} else {
		return echo.ErrInternalServerError
	}
}
