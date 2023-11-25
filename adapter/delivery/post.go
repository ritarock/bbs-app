package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ritarock/bbs-app/domain"
	timeutil "github.com/ritarock/bbs-app/internal/time_util"
)

type postHandler struct {
	postUsecase domain.PostUsecase
}

func NewPostHandler(us domain.PostUsecase) *postHandler {
	return &postHandler{
		postUsecase: us,
	}
}

func (p *postHandler) Create(c echo.Context) error {
	var post domain.Post
	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	ctx := c.Request().Context()

	post.PostedAt = timeutil.Now(ctx)
	if err := post.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := p.postUsecase.Create(ctx, &post); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, post)
}

func (p *postHandler) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	posts, err := p.postUsecase.GetAll(ctx)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, posts)
}

func (p *postHandler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound)
	}

	ctx := c.Request().Context()
	post, err := p.postUsecase.GetById(ctx, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound)
	}

	return c.JSON(http.StatusOK, post)
}
