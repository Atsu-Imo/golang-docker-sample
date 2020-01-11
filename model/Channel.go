package model

import "github.com/jinzhu/gorm"

type Channel struct {
	gorm.Model
	channelId string
	name string
	title string
	url string
	thumbnail string
	category int
	rotation int
}