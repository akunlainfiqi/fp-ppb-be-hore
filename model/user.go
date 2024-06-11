package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                string         `json:"id" gorm:"primarykey"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Username          string         `json:"username" gorm:"unique"`
	Password          string         `json:"password"`
	ProfilePictureUrl string         `json:"profile_picture_url"`
}
