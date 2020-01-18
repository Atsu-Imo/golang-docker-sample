package model

import (
	"github.com/jinzhu/gorm"
	
	//postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Connect DBへの接続を返す
// TODO すでに接続が存在した場合にはどうなるのか？確認 
func Connect(dbInfo string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", dbInfo)	
	return db, err
}