package db

import (
	"gorm.io/gorm"
	"youtube/internal/model"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase(db *gorm.DB) *Database {
	return &Database{db: db}
}

func (d Database) Migrate() error {
	return d.db.AutoMigrate(
		&model.User{},
		&model.ApiKey{},
	)
}

func (d Database) Populate() error {
	return nil
}

func (d Database) CheckApiKey(key string) (*model.User, error) {
	user := &model.User{}
	result := d.db.InnerJoins("ApiKey").Where("access_key = ?", key).Limit(1).Find(user)
	if result.RowsAffected == 0 {
		return nil, result.Error
	}
	return user, result.Error
}
