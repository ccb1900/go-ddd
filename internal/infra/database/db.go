package database

import (
	"goddd/pkg/config"

	"gorm.io/gorm"
)


func NewDB(cfg config.AppConfig) *gorm.DB {
	return &gorm.DB{}
}