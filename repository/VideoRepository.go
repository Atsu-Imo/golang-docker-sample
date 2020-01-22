package repository

import (
	"github.com/Atsu-Imo/golang-docker-sample/model"
	"github.com/jinzhu/gorm"
)

type VideoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(dbInfo string) *VideoRepository {
	db, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	return &VideoRepository{db: db}
}

func (r *VideoRepository) FindAll() []model.Video {
	var videos []model.Video
	r.db.Find(&videos)
	return videos
}