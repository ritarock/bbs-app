package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ritarock/bbs-app/db"
)

const (
	ServerError = http.StatusInternalServerError
	ServerOK    = http.StatusOK
)

func errorResponse(c *gin.Context, err error) {
	c.JSON(ServerError, map[string]any{
		"code":    ServerError,
		"message": err.Error(),
	})
}

func setupRouter(
	postHandler *postHandler,
	commentHandler *commendHandler,
) *gin.Engine {
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

	postV1 := r.Group("/backend/api/v1")
	{
		postV1.POST("/posts", postHandler.create)
		postV1.GET("/posts/:id", postHandler.readById)
		postV1.PUT("/posts/:id", postHandler.update)
		postV1.DELETE("/posts/:id", postHandler.delete)
		postV1.GET("/posts", postHandler.readAll)
	}

	commentV1 := r.Group("/backend/api/v1")
	{
		commentV1.POST("/comments", commentHandler.create)
		commentV1.GET("/comments/:id", commentHandler.readById)
		commentV1.PUT("/comments/:id", commentHandler.update)
		commentV1.DELETE("/comments/:id", commentHandler.delete)
	}

	r.GET("/backend/api/v1/topics/:id/comments", commentHandler.readAllByPost)

	return r
}

func Run() {
	dbClient, err := db.Connection()
	if err != nil {
		panic(err)
	}
	defer dbClient.Close()

	postHandler := newPostHandler(dbClient)
	commentHandler := newCommendHandler(dbClient)

	r := setupRouter(postHandler, commentHandler)

	r.Run(":8080")
}
