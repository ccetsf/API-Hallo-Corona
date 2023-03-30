package routes

import (
	"hallo-corona/database"
	"hallo-corona/handlers"
	"hallo-corona/pkg/middleware"
	repo "hallo-corona/repositories"

	"github.com/labstack/echo"
)

func AuthRoutes(e *echo.Group) {
	authRepo := repo.RepositoryAuth(database.DB)
	h := handlers.HandlerAuth(authRepo)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
	e.GET("/check-auth", middleware.Auth(h.CheckAuth))
}
