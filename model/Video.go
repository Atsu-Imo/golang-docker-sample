package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Video struct {
	gorm.Model
	videoID     string
	channelID   string
	title       string
	url         string
	length      string
	publishedAt time.Time
}
