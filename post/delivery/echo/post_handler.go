package echo

import (
	"net/http"
	"ritarock/bbs-app/domain"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	PUsecase domain.PostUsecase
}

func NewPostHandler(e *echo.Echo, us domain.PostUsecase) {
	handler := &PostHandler{
		PUsecase: us,
	}
	e.POST("/backend/api/v1/posts", handler.Create)
	e.GET("/backend/api/v1/posts", handler.GetAll)
	e.GET("/backend/api/v1/posts/:id", handler.GetById)
}

func (p *PostHandler) Create(c echo.Context) error {
	var post domain.Post
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	if err := p.PUsecase.Create(ctx, &post); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, post)
}

func (p *PostHandler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	ctx := c.Request().Context()

	post, err := p.PUsecase.GetById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	return c.JSON(http.StatusOK, post)
}

func (p *PostHandler) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	posts, err := p.PUsecase.GetAll(ctx)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, posts)
}
