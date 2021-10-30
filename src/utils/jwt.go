package utils

import (
	"errors"
	"fmt"
	"net/http"
	"src/config"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaims struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Id    string `json:"id"`
	jwt.StandardClaims
}

var (
	JWTconfig middleware.JWTConfig = middleware.JWTConfig{
		TokenLookup: "header:Authorization",
		Claims:      &JWTCustomClaims{},
	}
)

func GenerateToken(c echo.Context, nama, email, id string) (string, error) {
	// Set custom claims
	claims := &JWTCustomClaims{
		nama,
		email,
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.GetConfigs(c).Secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func GetJWTData(c echo.Context, header http.Header) (jwt.MapClaims, error) {
	// var data jwtCustomClaims
	var authData string = header["Authorization"][0]
	var token string = strings.TrimPrefix(authData, "bearer ")

	finalToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetConfigs(c).Secret), nil
	})
	if err != nil {
		return nil, errors.New("unexpected error on getting jwt data")
	}

	claims := finalToken.Claims.(jwt.MapClaims)
	return claims, nil
}
