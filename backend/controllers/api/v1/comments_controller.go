package v1

import (
	"backend/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func IndexComments(c *gin.Context) {
	comments := models.Comment{}

	c.JSON(200, gin.H{
		"comments": comments.GetAll(),
	})
}

func CreateComments(c *gin.Context) {
	comment := models.Comment{}

	err := c.ShouldBindJSON(&comment)
	if err != nil {
		fmt.Println(err)
	}
	comment.Create(c.Param("theme_id"))

	c.Redirect(302,
		"/service/v1/themes/"+c.Param("theme_id"))
}
