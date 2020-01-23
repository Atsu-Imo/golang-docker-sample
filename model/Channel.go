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

type channelRepository struct {
	db *gorm.DB
}
// ChannelRepository Channel関連の操作を行う
type ChannelRepository interface {
	FindAll() []Channel
	FindByChannelID(channelID string) Channel
}
// NewChannelRepository 新規のChannelRepositoryを返却する
func NewChannelRepository(db *gorm.DB) *channelRepository{
	return &channelRepository{db: db}
}

// FindAll すべてのチャンネル情報を返却する
func (r *channelRepository) FindAll() []Channel {
	var channels []Channel
	r.db.Find(&channels)
	return channels
}

// FindByChannelID channel_idで指定したchannelを返却する
func (r *channelRepository) FindByChannelID(channelID string) Channel {
	var channel Channel
	r.db.Where("channel_ID = ?", channelID).Find(&channel)
	return channel
}
