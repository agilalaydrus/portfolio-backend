package models

import (
	"gorm.io/gorm"
	"time"
)

type Contact struct {
	gorm.Model
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Phone       string     `json:"phone"`
	LinkedIn    string     `json:"linkedin"`
	Message     string     `json:"message"`
	IsRead      bool       `json:"is_read"`
	RespondedAt *time.Time `json:"responded_at"`
}
