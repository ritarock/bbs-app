package models

type Theme struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
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

	var theme = Theme{Title: t.Title, Detail: t.Detail}
	db.Create(&theme)
}

func (t Theme) Read(id string) interface{} {
	db := GormConnect()
	defer db.Close()

	result := db.First(&t, id)

	return result.Value
}

func (t Theme) Update(id string) {
	db := GormConnect()
	defer db.Close()

	var theme Theme
	db.First(&theme, id)
	theme.Title = t.Title
	theme.Detail = t.Detail
	db.Save(&theme)
}

func (t Theme) Delete(id string) {
	db := GormConnect()
	defer db.Close()

	db.Delete(&t, id)
}
