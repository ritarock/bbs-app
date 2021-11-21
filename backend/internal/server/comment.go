package server

import (
	"net/http"
	"ritarock/bbs-app/backend/internal/model"
	"ritarock/bbs-app/backend/internal/server/types"

	"github.com/labstack/echo/v4"
)

func readComments(c echo.Context) error {
	topicId := c.Param("topic_id")
	comments := model.ReadComments(topicId)
	var response struct {
		Code int `json:"code"`
		Data []types.Comment
	}
	response.Code = 200
	for _, comment := range comments {
		response.Data = append(response.Data, types.Comment(comment))
	}
	return c.JSON(http.StatusOK, response)
}

func creatComments(c echo.Context) error {
	topicId := c.Param("topic_id")
	cmt := new(types.Comment)
	if err := c.Bind(cmt); err != nil {
		return err
	}
	bindComment := &types.Comment{
		TopicId: cmt.TopicId,
		Body:    cmt.Body,
	}

	comment := model.Comment(*bindComment)
	id := comment.Create(topicId)

	var response struct {
		Code int `json:"code"`
		Data []types.Comment
	}
	response.Code = 200
	response.Data = append(response.Data, types.Comment{
		Id:      id,
		TopicId: comment.TopicId,
		Body:    comment.Body,
	})

	return c.JSON(http.StatusOK, response)
}

func readComment(c echo.Context) error {
	id := c.Param("id")
	cmt := model.Comment{
		Id: id,
	}
	comment := cmt.Read()

	var response struct {
		Code int `json:"code"`
		Data []types.Comment
	}
	response.Code = 200
	response.Data = append(response.Data, types.Comment(comment))

	return c.JSON(http.StatusOK, response)
}

func updateComment(c echo.Context) error {
	id := c.Param("id")
	cmtModel := model.Comment{
		Id: id,
	}
	comment := cmtModel.Read()
	cmt := new(types.Comment)
	if err := c.Bind(cmt); err != nil {
		return err
	}
	bindComment := &types.Comment{
		Id:      id,
		TopicId: comment.TopicId,
		Body:    cmt.Body,
	}
	comment = model.Comment(*bindComment)
	comment.Update()

	var response struct {
		Code int `json:"code"`
		Data []types.Comment
	}
	response.Code = 200
	response.Data = append(response.Data, types.Comment(comment))

	return c.JSON(http.StatusOK, response)
}

func deleteComment(c echo.Context) error {
	id := c.Param("id")
	comment := model.Comment{
		Id: id,
	}
	comment.Delete()
	return readComment(c)
}
