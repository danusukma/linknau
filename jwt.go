// 4. Implement a secure authentication and authorization mechanism using JWT authentication
// lib: github.com/dgrijalva/jwt-go and github.com/labstack/echo/v4

// - Generate JWT:
//   curl --location 'http://localhost:8080/login' \
// --form 'username="user"' \
// --form 'password="pass"'

//- Validate JWT:
//	curl --location 'http://localhost:8080/protected' \
// 	--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRhbnUiLCJleHAiOjE3MjY3NjMxNjh9.BvOlN6v4r-pFbCbsHotGrcHOP4lcfwFJTyHr89BHalM'

package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var jwtKey = []byte("secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func generateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username != "danu" || password != "123456" {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	token, err := generateJWT(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Could not generate token")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func protected(c echo.Context) error {
	tokenStr := c.Request().Header.Get("Authorization")

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return c.JSON(http.StatusUnauthorized, "Invalid token")
	}

	return c.JSON(http.StatusOK, "Validate token success, welcome "+claims.Username)
}

func main() {
	e := echo.New()

	e.POST("/login", login)
	e.GET("/protected", protected)

	e.Start(":8080")
}
