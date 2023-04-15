package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	CREATE_COMMENT_URL               = "/backend/api/v1/comments"
	READ_UPDATE_DELETE_COMMENT_URL   = "/backend/api/v1/comments/1"
	READ_ALL_BY_POST_FOR_COMMENT_URL = "/backend/api/v1/posts/1/comments"

	CREATE_COMMENT_BODY = `{
		"content": "test_content",
		"post_id": 1
	}`

	UPDATE_COMMENT_BODY = `{
		"content": "test_updated_content",
		"post_id": 1
	}`
)

func TestCreateCommentRoute(t *testing.T) {
	dbClient, postHandler, commentHandler := setupHandler()
	setPost(dbClient)
	defer dbClient.Close()

	router := setupRouter(postHandler, commentHandler)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := strings.NewReader(CREATE_COMMENT_BODY)
	c.Request, _ = http.NewRequest("POST", CREATE_COMMENT_URL, body)
	c.Request.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, 200)
	assert.JSONEq(t, w.Body.String(), `{
		"id": 1,
		"content": "test_content",
		"post_id": 1,
		"commented_at": "2023-01-01T00:00:00+09:00"
	}`)
}

func TestReadCommentRoute(t *testing.T) {
	dbClient, postHandler, commentHandler := setupHandler()
	setPost(dbClient)
	setComment(dbClient)
	defer dbClient.Close()

	router := setupRouter(postHandler, commentHandler)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", READ_UPDATE_DELETE_COMMENT_URL, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, w.Body.String(), `{
		"id": 1,
		"content": "test_content",
		"post_id": 1,
		"commented_at":"2023-01-01T00:00:00+09:00"
	}`)
}

func TestUpdateCommentRoute(t *testing.T) {
	dbClient, postHandler, commentHandler := setupHandler()
	setPost(dbClient)
	setComment(dbClient)
	defer dbClient.Close()

	router := setupRouter(postHandler, commentHandler)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := strings.NewReader(UPDATE_COMMENT_BODY)
	c.Request, _ = http.NewRequest("PUT", READ_UPDATE_DELETE_COMMENT_URL, body)
	c.Request.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, c.Request)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, w.Body.String(), `{
		"id": 1,
		"content": "test_updated_content",
		"post_id": 1,
		"commented_at":"2023-01-01T00:00:00+09:00"
	}`)
}

func TestDeleteCommentRoute(t *testing.T) {
	dbClient, postHandler, commentHandler := setupHandler()
	setPost(dbClient)
	setComment(dbClient)
	defer dbClient.Close()

	router := setupRouter(postHandler, commentHandler)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("DELETE", READ_UPDATE_DELETE_COMMENT_URL, nil)
	c.Request.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, 200)
	assert.JSONEq(t, w.Body.String(), `{
		"id": 1,
		"content": "test_content",
		"post_id": 1,
		"commented_at":"2023-01-01T00:00:00+09:00"
	}`)

	got := dbClient.Comment.Query().CountX(context.Background())
	want := 0
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReadAllByPost(t *testing.T) {
	dbClient, postHandler, commentHandler := setupHandler()
	setPost(dbClient)
	setComment(dbClient)
	setComment(dbClient)
	setComment(dbClient)
	defer dbClient.Close()

	router := setupRouter(postHandler, commentHandler)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", READ_ALL_BY_POST_FOR_COMMENT_URL, nil)
	c.Request.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, 200)
	assert.JSONEq(t, w.Body.String(), `[
		{
			"id": 1,
			"content": "test_content",
			"post_id": 1,
			"commented_at":"2023-01-01T00:00:00+09:00"
		},
		{
			"id": 2,
			"content": "test_content",
			"post_id": 1,
			"commented_at":"2023-01-01T00:00:00+09:00"
		},
		{
			"id": 3,
			"content": "test_content",
			"post_id": 1,
			"commented_at":"2023-01-01T00:00:00+09:00"
		}
	]`)
}
