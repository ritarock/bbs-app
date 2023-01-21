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

type topicHandler struct {
	db *ent.Client
}

func newTopicHander(db *ent.Client) *topicHandler {
	return &topicHandler{db}
}

func (handler *topicHandler) create(c *gin.Context) {
	topic := v1.Topic{}
	if err := c.ShouldBind(&topic); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	created, err := handler.db.
		Topic.
		Create().
		SetName(topic.Name).
		SetDetail(topic.Detail).
		Save(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, v1.Topic{
		Id:     int32(created.ID),
		Name:   created.Name,
		Detail: created.Detail,
	})
}

func (handler *topicHandler) readById(c *gin.Context) {
	urlId := c.Param("id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	read := handler.db.
		Topic.
		Query().
		Where(topic.ID(id)).
		OnlyX(context.Background())
	c.JSON(http.StatusOK, v1.Topic{
		Id:     int32(read.ID),
		Name:   read.Name,
		Detail: read.Detail,
	})
}

func (handler *topicHandler) update(c *gin.Context) {
	urlId := c.Param("id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	topicObj := v1.Topic{}
	if err := c.ShouldBind(&topicObj); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	updated, err := handler.db.
		Topic.
		UpdateOneID(id).
		SetName(topicObj.Name).
		SetDetail(topicObj.Detail).
		SetUpdatedAt(time.Now()).
		Save(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, v1.Topic{
		Id:     int32(updated.ID),
		Name:   updated.Name,
		Detail: updated.Detail,
	})
}

func (handler *topicHandler) delete(c *gin.Context) {
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
		Topic.
		Query().
		Where(topic.ID(id)).
		OnlyX(ctx)

	_, err = handler.db.
		Comment.
		Delete().
		Where(comment.TopicID(read.ID)).
		Exec(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	deleted := handler.db.
		Topic.
		DeleteOneID(id).
		Exec(ctx)

	if deleted != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, v1.Topic{
		Id:     int32(read.ID),
		Name:   read.Name,
		Detail: read.Detail,
	})
}

func (handler *topicHandler) readAll(c *gin.Context) {
	searchedAll := handler.db.
		Topic.
		Query().
		AllX(context.Background())
	resp := []v1.Topic{}
	for _, searched := range searchedAll {
		resp = append(resp, v1.Topic{
			Id:     int32(searched.ID),
			Name:   searched.Name,
			Detail: searched.Detail,
		})
	}
	c.JSON(http.StatusOK, resp)
}
