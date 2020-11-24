package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
