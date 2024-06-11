package model

import "gorm.io/gorm"

type Expense struct {
	UserId     string  `json:"user_id"`
	ExpenseId  int     `json:"expense_id"`
	Title      string  `json:"title"`
	Amount     float64 `json:"amount" gorm:"type:numeric(20,2)"`
	Date       string  `json:"date"`
	TypeId     int     `json:"type_id"`
	CategoryId string  `json:"category_id"`
	gorm.Model
}
