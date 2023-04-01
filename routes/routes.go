package routes

import "github.com/labstack/echo"

func RouteInit(e *echo.Group) {
	AuthRoutes(e)
	ArticleRoutes(e)
}
