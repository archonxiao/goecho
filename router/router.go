package router

import (
	"github.com/labstack/echo/v4"
	"goecho/controller"
	"net/http"
)

func Init(e *echo.Echo) *echo.Echo {
	e.GET("/articles", controller.GetArticles)
	e.GET("/articles/:id", controller.GetArticle)
	e.POST("/articles", controller.SaveArticle)
	e.PUT("/articles/:id", controller.UpdateArticle)
	e.DELETE("/articles/:id", controller.DeleteArticle)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Golang Echo example")
	})
	return e
}
