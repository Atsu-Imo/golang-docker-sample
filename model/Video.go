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
	FindByChannelID(channelID []string, page int, limit int) []Video
}

func NewVideoRepository(db *gorm.DB) *videoRepository {
	return &videoRepository{db: db}
}

func (r *videoRepository) FindAll() []Video {
	var videos []Video
	r.db.Find(&videos)
	return videos
}

func (r *videoRepository) FindByChannelID(channelID []string, page int, limit int) []Video {
	var videos []Video
	if channelID == nil {
		r.db.Find(&videos)
	}
	state := r.db.Where("Channel_ID in (?)", channelID)
	state.Offset(page * limit).Limit(limit).Find(&videos)
	return videos
}
