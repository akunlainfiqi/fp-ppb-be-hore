package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mobile.mabuk.cyou/model"
)

var db *gorm.DB

var (
	DB_HOST = "localhost"
	DB_PORT = "5432"
	DB_USER = "postgres"
	DB_PASS = "passowrd"
	DB_NAME = "ppb"
)

func init() {
	if v := os.Getenv("DB_HOST"); v != "" {
		DB_HOST = v
	}
	if v := os.Getenv("DB_PORT"); v != "" {
		DB_PORT = v
	}
	if v := os.Getenv("DB_USER"); v != "" {
		DB_USER = v
	}
	if v := os.Getenv("DB_PASS"); v != "" {
		DB_PASS = v
	}
	if v := os.Getenv("DB_NAME"); v != "" {
		DB_NAME = v
	}
}
func main() {
	e := echo.New()
	e.Debug = true
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderContentEncoding,
		},
	}))

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{}, &model.Expense{}, &model.Category{}, &model.Budget{})

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/login", login)
	e.POST("/register", register)

	// Restricted group
	r := e.Group("/v1")

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.JwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	r.Use(echojwt.WithConfig(config))

	r.GET("/me", me)
	r.GET("/expenses", getExpensesByUser)
	r.GET("/expenses/latest", getLatestUpdatedUserExpenses)
	r.POST("/expenses", insertBatchExpense)

	r.GET("/categories", getCategoryByUser)
	r.POST("/categories", insertBatchCategory)

	r.GET("/budgets", getBudgetByUser)
	r.POST("/budgets", insertBatchBudget)

	admin := r.Group("/admin")
	admin.GET("/expenses", getAllExpenses)

	e.Logger.Fatal(e.Start(":8080"))
}
