package main

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"mobile.mabuk.cyou/model"
)

func me(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	name := claims.Name

	var users *model.User
	q := db.Where("username = ?", name).First(&users)
	if q.Error != nil {
		return echo.ErrUnauthorized
	}

	return c.JSON(200, echo.Map{
		"status": "success",
		"data":   users,
	})
}