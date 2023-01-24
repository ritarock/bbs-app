package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ritarock/bbs-app/db"
)

func Run() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	dbClient, err := db.Connection()
	if err != nil {
		panic(err)
	}
	defer dbClient.Close()

	topicHandler := newTopicHander(dbClient)
	commentHandler := newCommentHandler(dbClient)

	topicV1 := r.Group("/backend/api/v1")
	{
		topicV1.POST("/topics", topicHandler.create)
		topicV1.GET("/topics/:id", topicHandler.readById)
		topicV1.PUT("/topics/:id", topicHandler.update)
		topicV1.DELETE("/topics/:id", topicHandler.delete)
		topicV1.GET("/topics", topicHandler.readAll)
	}

	commentV1 := r.Group("/backend/api/v1")
	{
		commentV1.POST("/comments", commentHandler.create)
		commentV1.GET("/comments/:id", commentHandler.readById)
		commentV1.PUT("/comments/:id", commentHandler.update)
		commentV1.DELETE("/comments/:id", commentHandler.delete)
	}

	r.GET("/backend/api/v1/topics/:id/comments", commentHandler.readAllByTopic)

	r.Run()
}
