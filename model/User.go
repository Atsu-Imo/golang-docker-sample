package model

import "github.com/jinzhu/gorm"

// User users
type User struct {
	gorm.Model
	username string `gorm:"varchar(256) unique not null"`
}