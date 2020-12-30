package server

import (
	model "backend/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

const VERSION = "v0"

func Run() {
	engine := gin.Default()
	model.DbInit()

	router := engine.Group("service")
	{
		api := router.Group(VERSION)
		{
			api.GET("/themes", func(c *gin.Context) {
				theme := model.Theme{}
				c.JSON(200, gin.H{
					"themes": theme.GetAll(),
				})
			})

			api.POST("/themes", func(c *gin.Context) {
				theme := model.Theme{}
				err := c.ShouldBindJSON(&theme)
				if err != nil {
					fmt.Println(err)
				}
				theme.Create()
				c.Redirect(302,
					"/service/"+VERSION+"/themes")
			})

			api.GET("/themes/:id", func(c *gin.Context) {
				var t model.Theme
				var comment model.Comment
				c.JSON(200, gin.H{
					"theme":   t.Get(c.Param("id")),
					"comment": comment.GetByThemeId(c.Param("id")),
				})
			})

			api.PUT("/themes/:id", func(c *gin.Context) {
				theme := model.Theme{}
				err := c.ShouldBindJSON(&theme)
				if err != nil {
					fmt.Println(err)
				}
				theme.Update(c.Param("id"))
				c.Redirect(302,
					"/service/"+VERSION+"/themes/"+c.Param("id"))
			})

			api.DELETE("/themes/:id", func(c *gin.Context) {
				theme := model.Theme{}
				theme.Delete(c.Param("id"))
				c.Redirect(302,
					"/service/"+VERSION+"/themes")
			})

			api.POST("/themes/:theme_id/comments", func(c *gin.Context) {
				comment := model.Comment{}
				err := c.ShouldBindJSON(&comment)
				if err != nil {
					fmt.Println(err)
				}
				comment.Create(c.Param("theme_id"))
				c.Redirect(302,
					"/service/"+VERSION+"/themes/"+c.Param("theme_id"))
			})
		}
	}
	engine.Run()
}
