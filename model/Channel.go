package model

import "github.com/jinzhu/gorm"

//Channel チャンネルテーブル
type Channel struct {
	gorm.Model
	ChannelID string
	Name string
	Title string
	URL string
	Thumbnail string
	Category int
	Rotation int
}