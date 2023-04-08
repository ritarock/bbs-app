package server

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ritarock/bbs-app/ent"
	"github.com/ritarock/bbs-app/ent/post"
)

type Post struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	PostedAt time.Time `json:"posted_at"`
}

type postHandler struct {
	db *ent.Client
}

func newPostHandler(db *ent.Client) *postHandler {
	return &postHandler{db}
}

func (ph *postHandler) create(c *gin.Context) {
	post := Post{}
	if err := c.ShouldBind(&post); err != nil {
		errorResponse(c, err)
	}
	created, err := ph.db.
		Post.
		Create().
		SetTitle(post.Title).
		SetContent(post.Content).
		SetPostedAt(timeNow).
		Save(context.Background())
	if err != nil {
		errorResponse(c, err)
	}

	c.JSON(ServerOK, Post{
		Id:       created.ID,
		Title:    created.Title,
		Content:  created.Content,
		PostedAt: created.PostedAt,
	})
}

func (ph *postHandler) readById(c *gin.Context) {
	urlId := c.Param("id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		errorResponse(c, err)
	}
	read := ph.db.
		Post.
		Query().
		Where(post.ID(id)).
		OnlyX(context.Background())
	c.JSON(ServerOK, Post{
		Id:       read.ID,
		Title:    read.Title,
		Content:  read.Content,
		PostedAt: read.PostedAt,
	})
}

func (ph *postHandler) update(c *gin.Context) {
	urlId := c.Param("id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		errorResponse(c, err)
	}

	post := Post{}
	if err := c.ShouldBind(&post); err != nil {
		errorResponse(c, err)
	}

	updated, err := ph.db.
		Post.
		UpdateOneID(id).
		SetTitle(post.Title).
		SetContent(post.Content).
		SetPostedAt(timeNow).
		Save(context.Background())
	if err != nil {
		errorResponse(c, err)
	}
	c.JSON(ServerOK, Post{
		Id:       updated.ID,
		Title:    updated.Title,
		Content:  updated.Content,
		PostedAt: updated.PostedAt,
	})
}

func (ph *postHandler) delete(c *gin.Context) {
	urlId := c.Param("id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		errorResponse(c, err)
	}
	ctx := context.Background()

	read := ph.db.
		Post.
		Query().
		Where(post.ID(id)).
		OnlyX(ctx)

	err = ph.db.
		Post.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		errorResponse(c, err)
	}
	c.JSON(ServerOK, Post{
		Id:       read.ID,
		Title:    read.Title,
		Content:  read.Content,
		PostedAt: read.PostedAt,
	})
}

func (ph *postHandler) readAll(c *gin.Context) {
	searchedAll, err := ph.db.
		Post.
		Query().
		All(context.Background())
	if err != nil {
		errorResponse(c, err)
	}
	response := []Post{}
	for _, searched := range searchedAll {
		response = append(response, Post{
			Id:       searched.ID,
			Title:    searched.Title,
			Content:  searched.Content,
			PostedAt: searched.PostedAt,
		})
	}
	c.JSON(ServerOK, response)
}
