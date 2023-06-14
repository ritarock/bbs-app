package http

import (
	"encoding/json"
	"net/http"
	"ritarock/bbs-app/domain"
	"strconv"
	"strings"
	"time"
)

type CommentHandler struct {
	CUsecase domain.CommentUsecase
}

type Ping struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

func NewCommentHandler(h *http.ServeMux, us domain.CommentUsecase) {
	handler := &CommentHandler{
		CUsecase: us,
	}
	h.HandleFunc("/backend/api/v1/comments/ping", handler.PingHandler)
	h.HandleFunc("/backend/api/v1/comments", handler.CommentHandler)
	h.HandleFunc("/backend/api/v1/comments/", handler.CommentsHandler)
}

func (c *CommentHandler) PingHandler(w http.ResponseWriter, r *http.Request) {
	ping := Ping{Status: http.StatusOK, Result: "ok"}

	res, err := json.Marshal(ping)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (c *CommentHandler) CommentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		c.Create(w, r)
	case "GET":
		c.GetAllByPost(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (c *CommentHandler) CommentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.GetAllByPost(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (c *CommentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var comment domain.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err := c.CUsecase.Create(ctx, &comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comment.CommentedAt = time.Now()
	res, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (c *CommentHandler) GetAllByPost(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/backend/api/v1/comments/")
	id, err := strconv.Atoi(sub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	comments, err := c.CUsecase.GetAllByPost(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	res, err := json.Marshal(comments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
