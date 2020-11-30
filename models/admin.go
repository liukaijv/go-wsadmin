package models

import (
	"gorm.io/gorm"
	"time"
)

type Admin struct {
	gorm.Model
	Username    string     `json:"username"`
	Password    string     `json:"password"`
	Token       string     `json:"token"`
	TokenExpire *time.Time `json:"token_expire"`
}
