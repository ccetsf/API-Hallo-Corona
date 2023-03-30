package repo

import (
	"hallo-corona/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(username string) (models.User, error)
	CheckAuth(Id int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (repo *repository) Register(user models.User) (models.User, error) {
	err := repo.db.Create(&user).Error
	return user, err
}

func (repo *repository) Login(username string) (models.User, error) {
	var user models.User
	err := repo.db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (repo *repository) CheckAuth(Id int) (models.User, error) {
	var user models.User
	err := repo.db.Where("id = ?", Id).First(&user).Error
	return user, err
}
