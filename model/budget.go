package model

type Budget struct {
	Id         int     `gorm:"primarykey,autoIncrement"`
	UserId     string  `json:"user_id"`
	BudgetId   int     `json:"_id"`
	Amount     float64 `json:"amount" gorm:"type:numeric(20,2)"`
	CategoryId string  `json:"category_id"`
	Date       string  `json:"date"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	DeletedAt  string  `json:"deleted_at"`
}
