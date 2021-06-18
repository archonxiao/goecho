package dao

import (
	"goecho/model"
	"gorm.io/datatypes"
)

func GetArticle(id int) model.Article {
	var article model.Article
	db := GetInstance()
	db.First(&article, id)
	article.Read = 1
	db.Save(&article)
	//var value string
	//db.Raw("SELECT content->>'$.k11' FROM articles WHERE id = ?", article.ID).Scan(&value)
	return article
}

func GetArticles() []model.Article {
	var articles []model.Article
	db := GetInstance()
	db.Find(&articles)
	return articles
}

func SaveArticle(title string, content datatypes.JSON) model.Article {
	db := GetInstance()
	article := model.Article{
		Title: title,
		Content: content,
	}
	db.Create(&article)
	return article
}

func UpdateArticle(id int, title string, content datatypes.JSON) model.Article {
	var article model.Article
	db := GetInstance()
	db.First(&article, id)
	if len(title) > 0 {
		article.Title = title
	}
	if len(content) > 0 {
		article.Content = content
	}
	db.Save(&article)
	return article
}

func DeleteArticle(id int) model.Article {
	var article model.Article
	db := GetInstance()
	db.Delete(&article, id)
	return article
}