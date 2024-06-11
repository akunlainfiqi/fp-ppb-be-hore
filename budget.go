package main

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"mobile.mabuk.cyou/model"
)

func getBudgetByUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	userId := claims.UserId

	z := db.Delete(&model.Budget{}, "user_id = ?", userId)
	if z.Error != nil {
		return echo.ErrInternalServerError
	}

	var budgets []model.Budget
	q := db.Where("user_id = ?", userId).Find(&budgets)
	if q.Error != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(200, echo.Map{
		"status": "success",
		"data":   budgets,
	})
}

func insertBatchBudget(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	userId := claims.UserId

	var budgets []model.Budget
	if err := c.Bind(&budgets); err != nil {
		return echo.ErrBadRequest
	}

	for i := range budgets {
		budgets[i].UserId = userId
	}

	q := db.Create(&budgets)
	if q.Error != nil {
		print(q.Error.Error())
		return echo.ErrInternalServerError
	}

	return c.JSON(200, echo.Map{
		"status": "success",
		"data":   budgets,
	})
}
