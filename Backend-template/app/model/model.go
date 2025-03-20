package model

import "gorm.io/gorm"

type Cat struct {
	gorm.Model
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
