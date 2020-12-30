package models

import (
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type Theme struct {
	gorm.Model
	Name     string
	Detail   string
	Comments []Comment
}

func (t Theme) GetAll() []Theme {
	db := GormConnect()
	defer db.Close()

	var themes []Theme
	db.Find(&themes)
	return themes
}

func (t Theme) Create() {
	db := GormConnect()
	defer db.Close()

	var theme = Theme{Name: t.Name, Detail: t.Detail}
	db.Create(&theme)
}

func (t Theme) Get(id string) Theme {
	db := GormConnect()
	defer db.Close()

	var theme Theme
	db.Where("id = ?", id).First(&theme)
	return theme
}

func (t Theme) Update(id string) {
	db := GormConnect()
	defer db.Close()

	var theme Theme
	db.Where("id = ?", id).First(&theme)
	theme.Name = t.Name
	theme.Detail = t.Detail
	db.Save(&theme)
}

func (t Theme) Delete(id string) {
	db := GormConnect()
	defer db.Close()

	var theme Theme
	db.Where("id = ?", id).First(&theme)
	db.Delete(&theme)
}
