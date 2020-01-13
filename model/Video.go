package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Video ビデオテーブル
type Video struct {
	gorm.Model
	VideoID     string
	ChannelID   string
	Title       string
	URL         string
	Length      string
	PublishedAt time.Time
}
