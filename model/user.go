package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email         string `gorm:"unique;not null" json:"email"`
	EmailVerified bool   `gorm:"default:false" json:"email_verified"`
	Password      string `json:"password"`
	GivenName     string `gorm:"not null" json:"given_name"`
	FamilyName    string `gorm:"not null" json:"family_name"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Provider      string `gorm:"default:local" json:"provider"`
}
