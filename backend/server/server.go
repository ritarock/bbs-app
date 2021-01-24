package server

import (
	controller "backend/controllers/api/v1"
	"backend/models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func init() {
	engine = gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		MaxAge: 24 * time.Hour,
	}))
	models.DbInit()
}

func Run() {
	router := engine.Group("service")
	{
		api := router.Group("v1")
		{
			api.GET("/themes", controller.Index)
			api.POST("/themes", controller.Create)
			api.GET("/themes/:id", controller.Read)
			api.PUT("/themes/:id", controller.Update)
			api.DELETE("/themes/:id", controller.Delete)
		}
	}

	engine.Run()
}
