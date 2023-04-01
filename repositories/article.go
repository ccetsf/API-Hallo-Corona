package repo

import (
	"hallo-corona/models"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateArticle(article models.Article) (models.Article, error)
	GetArticle(id int) (models.Article, error)
	FindArticles() ([]models.Article, error)
	DeleteArticle(article models.Article, id int) (models.Article, error)
	MyArticles(userId int) ([]models.Article, error)
	UpdateArticle(article models.Article) (models.Article, error)
}

func RepositoryArticle(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) CreateArticle(article models.Article) (models.Article, error) {
	err := repo.db.Create(&article).Error
	return article, err
}

func (repo *repository) GetArticle(id int) (models.Article, error) {
	var article models.Article
	err := repo.db.Preload("User").First(&article, id).Error
	return article, err
}

func (repo *repository) FindArticles() ([]models.Article, error) {
	var articles []models.Article
	err := repo.db.Preload("User").Find(&articles).Error
	return articles, err
}

func (repo *repository) DeleteArticle(article models.Article, id int) (models.Article, error) {
	err := repo.db.Delete(&article, id).Scan(&article).Error
	return article, err
}

func (repo *repository) MyArticles(userId int) ([]models.Article, error) {
	var articles []models.Article
	err := repo.db.Preload("User").Where("user_id = ?", userId).Find(&articles).Error
	return articles, err
}

func (repo *repository) UpdateArticle(article models.Article) (models.Article, error) {
	err := repo.db.Save(&article).Error
	return article, err
}
