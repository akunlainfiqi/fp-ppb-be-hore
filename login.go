package main

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"mobile.mabuk.cyou/model"
)

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return echo.ErrBadRequest
	}
	// Throws unauthorized error
	var user model.User
	userQuery := db.Where("username = ? AND password = ?", username, password).First(&user)
	if userQuery.Error != nil {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &model.JwtCustomClaims{
		Name:   username,
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   t,
	})
}
