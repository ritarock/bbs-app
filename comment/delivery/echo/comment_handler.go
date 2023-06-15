package echo

import (
	"net/http"
	"ritarock/bbs-app/domain"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	CUsecase domain.CommentUsecase
}

func NewPostHandler(e *echo.Echo, us domain.CommentUsecase) {
	handler := &CommentHandler{
		CUsecase: us,
	}
	e.POST("/backend/api/v1/comments", handler.Create)
	e.GET("/backend/api/v1/comments/:id", handler.GetAllByPost)
}

func (co *CommentHandler) Create(c echo.Context) error {
	var comment domain.Comment
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	if err := co.CUsecase.Create(ctx, &comment); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, comment)
}

func (co *CommentHandler) GetAllByPost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	ctx := c.Request().Context()
	comments, err := co.CUsecase.GetAllByPost(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, comments)
}
