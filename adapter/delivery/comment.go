package delivery

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ritarock/bbs-app/domain"
	"github.com/ritarock/bbs-app/port"
)

type commentHandler struct {
	commentUsecase port.CommentUsecase
}

func NewCommentHandler(us port.CommentUsecase) *commentHandler {
	return &commentHandler{
		commentUsecase: us,
	}
}

func (cm *commentHandler) Create(c echo.Context) error {
	var comment domain.Comment
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	comment.CommentedAt = time.Now()
	comment.PostID = postID

	if err := comment.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := cm.commentUsecase.Create(ctx, &comment); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, comment)
}

func (cm *commentHandler) GetAll(c echo.Context) error {
	postID, err := strconv.Atoi(c.Param("post_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	comments, err := cm.commentUsecase.GetAll(ctx, postID)
	if err != nil {
		if err == domain.ErrNotFound {
			return c.JSON(http.StatusNotFound, err)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, comments)
}
