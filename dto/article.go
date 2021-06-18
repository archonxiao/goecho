package dto

import "gorm.io/datatypes"

type ArticleRequestBody struct {
	Title   string `json:"title" form:"title" name:"title"`
	Content datatypes.JSON `json:"content" form:"content" name:"content"`
}