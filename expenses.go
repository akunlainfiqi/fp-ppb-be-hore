package main

import (
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"mobile.mabuk.cyou/model"
)

func getAllExpenses(c echo.Context) error {
	var expenses []model.Expense
	q := db.Find(&expenses)
	if q.Error != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(200, echo.Map{
		"status": "success",
		"data":   expenses,
	})
}

func getExpensesByUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	userId := claims.UserId

	var expenses []model.Expense
	q := db.Where("user_id = ?", userId).Find(&expenses)
	if q.Error != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(200, echo.Map{
		"status": "success",
		"data":   expenses,
	})
}

func getLatestUpdatedUserExpenses(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	userId := claims.UserId

	var expenses model.Expense
	q := db.Where("user_id = ?", userId).Order("updated_at desc").First(&expenses)
	if q.Error != nil {
		if q.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{
				"status":  "error",
				"message": "No expenses found",
				"data":    q.Error.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "error",
			"message": "Unhandled Error",
			"data":    q.Error.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"status": "success",
		"data":   expenses,
	})
}

func insertBatchExpense(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	userId := claims.UserId

	var expenses []model.Expense
	if err := c.Bind(&expenses); err != nil {
		log.Print(err)
		return echo.ErrBadRequest
	}

	z := db.Delete(&model.Expense{}, "user_id = ?", userId)
	if z.Error != nil {
		log.Print(z.Error.Error())
		return echo.ErrInternalServerError
	}

	for i := range expenses {
		expenses[i].UserId = userId
	}

	q := db.Create(&expenses)
	if q.Error != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(200, echo.Map{
		"status": "success",
		"data":   expenses,
	})
}
