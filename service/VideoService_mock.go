// +build mock

package service

import (
	"github.com/Atsu-Imo/golang-docker-sample/model"
	"github.com/jinzhu/gorm"
)
type VideoRepository struct {
	db *gorm.DB
}

func (m * VideoRepository) FindAll() []model.Video {
	return nil;
}