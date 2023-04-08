package server

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ritarock/bbs-app/ent"
	"github.com/ritarock/bbs-app/ent/comment"
	"github.com/ritarock/bbs-app/ent/post"
)

type Comment struct {
	Id          int       `json:"id"`
	Content     string    `json:"content"`
	CommentedAt time.Time `json:"commented_at"`
	PostId      int       `json:"post_id"`
}

type commendHandler struct {
	db *ent.Client
}

func newCommendHandler(db *ent.Client) *commendHandler {
	return &commendHandler{db}
}

func (ch *commendHandler) create(c *gin.Context) {
	comment := Comment{}
	if err := c.ShouldBind(&comment); err != nil {
		errorResponse(c, err)
	}
	ctx := context.Background()
	postId := ch.db.
		Post.
		Query().
		Where(post.ID(comment.PostId)).
		Select(post.FieldID).
		IntX(ctx)
	created, err := ch.db.
		Comment.
		Create().
		SetPostID(postId).
		SetContent(comment.Content).
		Save(ctx)
	if err != nil {
		errorResponse(c, err)
	}
	c.JSON(ServerOK, Comment{
		Id:          created.ID,
		Content:     created.Content,
		CommentedAt: created.CommentedAt,
		PostId:      created.PostID,
	})
}

func (ch *commendHandler) readById(c *gin.Context) {
	urlId := c.Param("id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		errorResponse(c, err)
	}
	read := ch.db.
		Comment.
		Query().
		Where(comment.ID(id)).
		OnlyX(context.Background())
	c.JSON(ServerOK, Comment{
		Id:          read.ID,
		Content:     read.Content,
		CommentedAt: read.CommentedAt,
		PostId:      read.PostID,
	})
}

func (ch *commendHandler) update(c *gin.Context) {
	urlId := c.Param("id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		errorResponse(c, err)
	}

	comment := Comment{}
	if err := c.ShouldBind(&comment); err != nil {
		errorResponse(c, err)
	}

	ctx := context.Background()

	postId := ch.db.
		Post.
		Query().
		Where(post.ID(comment.PostId)).
		Select(post.FieldID).
		IntX(ctx)

	updated, err := ch.db.
		Comment.
		UpdateOneID(id).
		SetContent(comment.Content).
		SetCommentedAt(comment.CommentedAt).
		SetPostID(postId).
		Save(ctx)
	if err != nil {
		errorResponse(c, err)
	}
	c.JSON(ServerOK, Comment{
		Id:          updated.ID,
		Content:     updated.Content,
		CommentedAt: updated.CommentedAt,
		PostId:      updated.PostID,
	})
}

func (ch *commendHandler) delete(c *gin.Context) {
	urlId := c.Param("id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		errorResponse(c, err)
	}
	ctx := context.Background()

	read := ch.db.
		Comment.
		Query().
		Where(comment.ID(id)).
		OnlyX(ctx)

	err = ch.db.
		Comment.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		errorResponse(c, err)
	}
	c.JSON(ServerOK, Comment{
		Id:          read.ID,
		Content:     read.Content,
		CommentedAt: read.CommentedAt,
		PostId:      read.PostID,
	})
}

func (ch *commendHandler) readAllByPost(c *gin.Context) {
	urlId := c.Param("id")
	topicId, err := strconv.Atoi(urlId)
	if err != nil {
		errorResponse(c, err)
	}
	searchedAll := ch.db.
		Comment.
		Query().
		Where(comment.PostIDEQ(topicId)).
		AllX(context.Background())
	response := []Comment{}
	for _, searched := range searchedAll {
		response = append(response, Comment{
			Id:          searched.ID,
			Content:     searched.Content,
			CommentedAt: searched.CommentedAt,
			PostId:      searched.PostID,
		})
	}
	c.JSON(ServerOK, response)
}
