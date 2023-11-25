package delivery

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ritarock/bbs-app/domain"
	timeutil "github.com/ritarock/bbs-app/internal/time_util"
)

type commentHandler struct {
	commentUsecase domain.CommentUsecase
}

func NewCommentHandler(us domain.CommentUsecase) *commentHandler {
	return &commentHandler{
		commentUsecase: us,
	}
}

func (ch *commentHandler) Create(c echo.Context) error {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound)
	}

	var comment domain.Comment
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	comment.PostId = postId
	ctx := c.Request().Context()
	comment.CommentedAt = timeutil.Now(ctx)

	if err := comment.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := ch.commentUsecase.Create(ctx, &comment); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, comment)
}

func (ch *commentHandler) GetAllByPostId(c echo.Context) error {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound)
	}

	ctx := c.Request().Context()
	comments, err := ch.commentUsecase.GetByPostId(ctx, postId)
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound)
	}

	return c.JSON(http.StatusOK, comments)
}
