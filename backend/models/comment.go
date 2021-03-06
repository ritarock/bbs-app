package models

import "strconv"

type Comment struct {
	Id   uint   `json:"id"`
	Body string `json:"body"`

	ThemeId int `json:"theme_id"`
}

func (c Comment) GetAll() []Comment {
	db := GormConnect()
	defer db.Close()

	var comments []Comment
	db.Find(&comments)
	return comments
}

func (c Comment) GetByThemeId(theme_id string) []Comment {
	db := GormConnect()
	defer db.Close()

	to_int_theme_id, _ := strconv.Atoi(theme_id)
	var comments []Comment
	db.Where("theme_id = ?", to_int_theme_id).Find(&comments)
	return comments
}

func (c Comment) Create(theme_id string) {
	db := GormConnect()
	defer db.Close()

	to_int_theme_id, _ := strconv.Atoi(theme_id)
	var comment = Comment{Body: c.Body, ThemeId: to_int_theme_id}
	db.Create(&comment)
}
