package model

import "gorm.io/gorm"

type ApiKey struct {
	gorm.Model
	UserId    uint
	AccessKey string
}
