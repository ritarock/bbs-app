package server

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ritarock/bbs-app/ent"
	"github.com/ritarock/bbs-app/ent/comment"
	"github.com/ritarock/bbs-app/ent/topic"
	v1 "github.com/ritarock/bbs-app/gen/v1/go"
)

type commendHandler struct {
	db *ent.Client
}

func newCommentHandler(db *ent.Client) *commendHandler {
	return &commendHandler{db}
}

func (handler *commendHandler) create(c *gin.Context) {
	commentObj := v1.Comment{}
	if err := c.ShouldBind(&commentObj); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	ctx := context.Background()
	topicId := handler.db.
		Topic.
		Query().
		Where(topic.ID(int(commentObj.TopicId))).
		Select(topic.FieldID).
		IntX(ctx)
	created, err := handler.db.
		Comment.
		Create().
		SetTopicID(topicId).
		SetBody(commentObj.Body).
		Save(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, v1.Comment{
		Id:      int32(created.ID),
		Body:    created.Body,
		TopicId: int32(topicId),
	})
}

func (handler *commendHandler) readById(c *gin.Context) {
	urlId := c.Param("id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	read := handler.db.
		Comment.
		Query().
		Where(comment.ID(id)).
		OnlyX(context.Background())
	c.JSON(http.StatusOK, v1.Comment{
		Id:      int32(read.ID),
		Body:    read.Body,
		TopicId: int32(read.TopicID),
	})
}

func (handler *commendHandler) update(c *gin.Context) {
	urlId := c.Param("id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	commentObj := v1.Comment{}
	if err := c.ShouldBind(&commentObj); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	ctx := context.Background()
	topicId := handler.db.
		Topic.
		Query().
		Where(topic.ID(int(commentObj.TopicId))).
		Select(topic.FieldID).
		IntX(ctx)

	updated, err := handler.db.
		Comment.
		UpdateOneID(id).
		SetTopicID(topicId).
		SetBody(commentObj.Body).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, v1.Comment{
		Id:      int32(updated.ID),
		Body:    updated.Body,
		TopicId: int32(updated.TopicID),
	})
}

func (handler *commendHandler) delete(c *gin.Context) {
	urlId := c.Param("id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	ctx := context.Background()
	read := handler.db.
		Comment.
		Query().
		Where(comment.ID(id)).
		OnlyX(ctx)
	deleted := handler.db.
		Comment.
		DeleteOneID(id).
		Exec(ctx)
	if deleted != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, v1.Comment{
		Id:      int32(read.ID),
		Body:    read.Body,
		TopicId: int32(read.TopicID),
	})
}

func (handler *commendHandler) readAllByTopic(c *gin.Context) {
	urlId := c.Param("id")
	topicId, err := strconv.Atoi(urlId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	searchedAll := handler.db.
		Comment.
		Query().
		Where(comment.TopicIDEQ(topicId)).
		AllX(context.Background())
	resp := []v1.Comment{}
	for _, searched := range searchedAll {
		resp = append(resp, v1.Comment{
			Id:      int32(searched.ID),
			Body:    searched.Body,
			TopicId: int32(topicId),
		})
	}
	c.JSON(http.StatusOK, resp)
}
