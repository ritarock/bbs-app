package server

import (
	"net/http"
	"ritarock/bbs-app/backend/internal/model"
	"ritarock/bbs-app/backend/internal/server/types"

	"github.com/labstack/echo/v4"
)

func readTopics(c echo.Context) error {
	topics := model.ReadTopics()
	var response struct {
		Code int `json:"code"`
		Data []types.Topic
	}
	response.Code = 200
	for _, topic := range topics {
		response.Data = append(response.Data, types.Topic(topic))
	}
	return c.JSON(http.StatusOK, response)
}

func createTopics(c echo.Context) error {
	t := new(types.Topic)
	if err := c.Bind(t); err != nil {
		return err
	}
	bindTopic := &types.Topic{
		Title:  t.Title,
		Detail: t.Detail,
	}

	topic := model.Topic(*bindTopic)
	id := topic.Create()

	var response struct {
		Code int `json:"code"`
		Data []types.Topic
	}
	response.Code = 200
	response.Data = append(response.Data, types.Topic{
		Id:     id,
		Title:  topic.Title,
		Detail: topic.Detail,
	})

	return c.JSON(http.StatusOK, response)
}

func readTopic(c echo.Context) error {
	id := c.Param("id")
	topic := model.Topic{
		Id: id,
	}
	topic.Read()
	var response struct {
		Code int `json:"code"`
		Data []types.Topic
	}
	response.Code = 200
	response.Data = append(response.Data, types.Topic(topic))

	return c.JSON(http.StatusOK, response)
}

func updateTopic(c echo.Context) error {
	id := c.Param("id")
	t := new(types.Topic)
	if err := c.Bind(t); err != nil {
		return err
	}
	bindTopic := &types.Topic{
		Id:     id,
		Title:  t.Title,
		Detail: t.Detail,
	}
	topic := model.Topic(*bindTopic)
	topic.Update()

	var response struct {
		Code int `json:"code"`
		Data []types.Topic
	}
	response.Code = 200
	response.Data = append(response.Data, types.Topic(topic))

	return c.JSON(http.StatusOK, response)
}

func deleteTopic(c echo.Context) error {
	id := c.Param("id")
	topic := model.Topic{
		Id: id,
	}
	topic.Delete()
	return readTopics(c)
}
