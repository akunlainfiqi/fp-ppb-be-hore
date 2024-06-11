package main

import (
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"mobile.mabuk.cyou/model"
)

func getAllCategories(c echo.Context) error {
	var categories []model.Category
	q := db.Find(&categories)
	if q.Error != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(200, echo.Map{
		"status": "success",
		"data":   categories,
	})
}

func getCategoryByUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	userId := claims.UserId

	var categories []model.Category
	q := db.Where("user_id = ?", userId).Find(&categories)
	if q.Error != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(200, echo.Map{
		"status": "success",
		"data":   categories,
	})
}

func insertBatchCategory(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	userId := claims.UserId

	var categories []model.Category
	if err := c.Bind(&categories); err != nil {
		log.Print(err)
		return echo.ErrBadRequest
	}

	z := db.Delete(&model.Category{}, "user_id = ?", userId)
	if z.Error != nil {
		log.Print(z.Error.Error())
		return echo.ErrInternalServerError
	}

	for i := range categories {
		categories[i].UserId = userId
	}

	q := db.Create(&categories)
	if q.Error != nil {
		log.Print(q.Error.Error())
		return echo.ErrInternalServerError
	}

	return c.JSON(200, echo.Map{
		"status": "success",
		"data":   categories,
	})
}
