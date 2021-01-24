package v1

import (
	"backend/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

var INDEX_URL = "/service/v1/themes"

func Index(c *gin.Context) {
	theme := models.Theme{}

	c.JSON(200, gin.H{
		"themes": theme.GetAll(),
	})
}

func Create(c *gin.Context) {
	theme := models.Theme{}

	err := c.ShouldBindJSON(&theme)
	if err != nil {
		fmt.Println(err)
	}

	theme.Create()

	c.Redirect(302,
		INDEX_URL)
}

func Read(c *gin.Context) {
	theme := models.Theme{}

	c.JSON(200, gin.H{
		"theme": theme.Read(c.Param("id")),
	})
}

func Update(c *gin.Context) {
	theme := models.Theme{}

	err := c.ShouldBindJSON(&theme)
	if err != nil {
		fmt.Println(err)
	}
	theme.Update(c.Param("id"))
	c.Redirect(302,
		INDEX_URL)
}

func Delete(c *gin.Context) {
	theme := models.Theme{}

	theme.Delete(c.Param("id"))
	c.Redirect(302,
		INDEX_URL)
}
