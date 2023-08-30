package database

import "github.com/shobky/giggs/model"

func Sync() {
	DB.AutoMigrate(&model.User{})
}
