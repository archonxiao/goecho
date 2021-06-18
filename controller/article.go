package controller

import (
	"github.com/labstack/echo/v4"
	"goecho/dto"
	"goecho/service"
	"net/http"
	"strconv"
)

func GetArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusPreconditionFailed)
	}
	article := service.GetArticle(id)
	return c.JSON(http.StatusOK, article)
}

func GetArticles(c echo.Context) error {
	articles := service.GetArticles()
	return c.JSON(http.StatusOK, articles)
}

func SaveArticle(c echo.Context) error {
	body := new(dto.ArticleRequestBody)
	if err := c.Bind(body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	//if err := c.Validate(&body); err != nil {
	//	return err
	//}
	article := service.SaveArticle(body.Title, body.Content)
	return c.JSON(http.StatusOK, article)
}

func UpdateArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusPreconditionFailed)
	}
	body := new(dto.ArticleRequestBody)
	if err := c.Bind(body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	article := service.UpdateArticle(id, body.Title, body.Content)
	return c.JSON(http.StatusOK, article)
}

func DeleteArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusPreconditionFailed)
	}
	article := service.DeleteArticle(id)
	return c.JSON(http.StatusOK, article)
}
