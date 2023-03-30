package repo

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}
