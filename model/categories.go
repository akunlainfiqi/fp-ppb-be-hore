package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	Id             int            `gorm:"primarykey,autoIncrement"`
	UserId         string         `json:"user_id"`
	CategoryId     int            `json:"_id"`
	Title          string         `json:"title"`
	IconCodePoint  int            `json:"icon_code_point"`
	CategoriesType int            `json:"categories_type"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
