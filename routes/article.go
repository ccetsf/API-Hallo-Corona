package routes

import (
	"hallo-corona/database"
	"hallo-corona/handlers"
	"hallo-corona/pkg/middleware"
	repo "hallo-corona/repositories"

	"github.com/labstack/echo"
)

func ArticleRoutes(e *echo.Group) {
	articleRepo := repo.RepositoryArticle(database.DB)
	h := handlers.HandlerArticle(articleRepo)

	e.POST("/article", middleware.Auth(middleware.UploadFile(h.CreateArticle)))
	e.GET("/article/:id", h.GetArticle)
	e.GET("/articles", h.FindArticles)
	e.DELETE("/article/:id", middleware.Auth(h.DeleteArticle))
	e.GET("/my-articles", middleware.Auth(h.MyArticles))
	e.PATCH("/article/:id", middleware.Auth(middleware.UploadFile(h.UpdateArticle)))
}
