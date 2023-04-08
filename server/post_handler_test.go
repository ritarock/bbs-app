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
	CREATE_AND_READ_ALL_POST_URL = "/backend/api/v1/posts"
	READ_UPDATE_DELETE_POST_URL  = "/backend/api/v1/posts/1"

	CREATE_POST_BODY = `{
		"title": "test_title",
		"content": "test_content"
	}`

	UPDATE_POST_BODY = `{
		"title": "test_updated_title",
		"content": "test_updated_content"
	}`
)

func TestCreatePostRoute(t *testing.T) {
	dbClient, postHandler, commentHandler := setupHandler()
	defer dbClient.Close()

	router := setupRouter(postHandler, commentHandler)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := strings.NewReader(CREATE_POST_BODY)
	c.Request, _ = http.NewRequest("POST", CREATE_AND_READ_ALL_POST_URL, body)
	c.Request.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, 200)
	assert.JSONEq(t, w.Body.String(), `{
		"id": 1,
		"title": "test_title",
		"content": "test_content",
		"posted_at": "2023-01-01T00:00:00+09:00"
	}`)
}

func TestReadPostRoute(t *testing.T) {
	dbClient, postHandler, commentHandler := setupHandler()
	setPost(dbClient)
	defer dbClient.Close()

	router := setupRouter(postHandler, commentHandler)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", READ_UPDATE_DELETE_POST_URL, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, w.Body.String(), `{
		"id": 1,
		"title": "test_title",
		"content": "test_content",
		"posted_at":"2023-01-01T00:00:00+09:00"
	}`)
}

func TestUpdatePostRoute(t *testing.T) {
	dbClient, postHandler, commentHandler := setupHandler()
	setPost(dbClient)
	defer dbClient.Close()

	router := setupRouter(postHandler, commentHandler)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := strings.NewReader(UPDATE_POST_BODY)
	c.Request, _ = http.NewRequest("PUT", READ_UPDATE_DELETE_POST_URL, body)
	c.Request.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, c.Request)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, w.Body.String(), `{
		"id": 1,
		"title": "test_updated_title",
		"content": "test_updated_content",
		"posted_at":"2023-01-01T00:00:00+09:00"
	}`)
}

func TestDeletePostRoute(t *testing.T) {
	dbClient, postHandler, commentHandler := setupHandler()
	setPost(dbClient)
	defer dbClient.Close()

	router := setupRouter(postHandler, commentHandler)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("DELETE", READ_UPDATE_DELETE_POST_URL, nil)
	c.Request.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, 200)
	assert.JSONEq(t, w.Body.String(), `{
		"id": 1,
		"title": "test_title",
		"content": "test_content",
		"posted_at":"2023-01-01T00:00:00+09:00"
	}`)

	got := dbClient.Post.Query().CountX(context.Background())
	want := 0
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReadPostAllRoute(t *testing.T) {
	dbClient, postHandler, commentHandler := setupHandler()
	setPost(dbClient)
	setPost(dbClient)
	setPost(dbClient)
	defer dbClient.Close()

	router := setupRouter(postHandler, commentHandler)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", CREATE_AND_READ_ALL_POST_URL, nil)
	c.Request.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, c.Request)

	assert.Equal(t, w.Code, 200)
	assert.JSONEq(t, w.Body.String(), `[
		{
			"id": 1,
			"title": "test_title",
			"content": "test_content",
			"posted_at":"2023-01-01T00:00:00+09:00"
		},
		{
			"id": 2,
			"title": "test_title",
			"content": "test_content",
			"posted_at":"2023-01-01T00:00:00+09:00"
		},
		{
			"id": 3,
			"title": "test_title",
			"content": "test_content",
			"posted_at":"2023-01-01T00:00:00+09:00"
		}
	]`)
}
