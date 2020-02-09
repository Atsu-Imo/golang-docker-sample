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

type videoRepository struct {
	db *gorm.DB
}

type VideoRepository interface {
	FindAll() []Video
	FindByChannelID(channelID []string) []Video
}

func NewVideoRepository(db *gorm.DB) *videoRepository {
	return &videoRepository{db: db}
}

func (r *videoRepository) FindAll() []Video {
	var videos []Video
	r.db.Find(&videos)
	return videos
}

func (r *videoRepository) FindByChannelID(channelID []string) []Video {
	var videos []Video
	if channelID == nil {
		r.db.Find(&videos)
	}
	r.db.Where("Channel_ID in (?)", channelID).Find(&videos)
	return videos
}
