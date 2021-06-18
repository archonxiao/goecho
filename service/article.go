package service

import (
	"goecho/dao"
	"goecho/model"
	"gorm.io/datatypes"
)

func GetArticle(id int) model.Article {
	return dao.GetArticle(id)
}

func GetArticles() []model.Article {
	return dao.GetArticles()
}

func SaveArticle(title string, content datatypes.JSON) model.Article {
	return dao.SaveArticle(title, content)
}

func UpdateArticle(id int, title string, content datatypes.JSON) model.Article {
	return dao.UpdateArticle(id, title, content)
}

func DeleteArticle(id int) model.Article {
	return dao.DeleteArticle(id)
}