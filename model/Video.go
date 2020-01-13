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

// Repository DBとの接続を管理
type Repository struct {
	DB *gorm.DB
}

// FindAll すべてのデータを返す
func (p *Repository) FindAll() []Video {
	var videos []Video
	p.DB.Find(&videos)
	return videos
}