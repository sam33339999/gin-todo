package common

import "gorm.io/gorm"

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}
