package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"unique"`
	Phone     string `gorm:"unique,index"`
	SessionID string `gorm:"index"`
	Code      string
}
