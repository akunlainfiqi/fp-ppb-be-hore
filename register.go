package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"mobile.mabuk.cyou/model"
)

func register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirm_password")

	if username == "" || password == "" || confirmPassword == "" {
		return echo.ErrBadRequest
	}

	if password != confirmPassword {
		return echo.ErrBadRequest
	}
	uid, err := uuid.NewRandom()
	if err != nil {
		return echo.ErrInternalServerError
	}

	user := model.User{
		ID:       uid.String(),
		Username: username,
		Password: password,
	}

	db.Create(&user)

	return c.JSON(200, echo.Map{
		"status": "success",
		"data":   user,
	})
}
