package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	SenderID   uint
	ReceiverID uint
	Body       string
}

type Contact struct {
	gorm.Model
	UserID        uint `json:"user_id"`
	ContactUserID uint `json:"contact_user_id"`
}
