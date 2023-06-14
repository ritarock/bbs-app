package http

import (
	"encoding/json"
	"net/http"
	"ritarock/bbs-app/domain"
	"strconv"
	"strings"
	"time"
)

type PostHandler struct {
	PUsecase domain.PostUsecase
}

type Ping struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

func NewPostHandler(h *http.ServeMux, us domain.PostUsecase) {
	handler := &PostHandler{
		PUsecase: us,
	}
	h.HandleFunc("/backend/api/v1/posts/ping", handler.PingHandler)
	h.HandleFunc("/backend/api/v1/posts", handler.PostHandler)
	h.HandleFunc("/backend/api/v1/posts/", handler.PostIdHandler)
}

func (p *PostHandler) PingHandler(w http.ResponseWriter, r *http.Request) {
	ping := Ping{Status: http.StatusOK, Result: "ok"}

	res, err := json.Marshal(ping)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (p *PostHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		p.Create(w, r)
	case "GET":
		p.GetAll(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (p *PostHandler) PostIdHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		p.GetByID(w, r)
	case "PUT":
		p.Update(w, r)
	case "DELETE":
		p.Delete(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (p *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var post domain.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err := p.PUsecase.Create(ctx, &post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post.PostedAt = time.Now()
	res, err := json.Marshal(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (p *PostHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/backend/api/v1/posts/")
	id, err := strconv.Atoi(sub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	post, err := p.PUsecase.GetById(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	res, err := json.Marshal(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (p *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts, err := p.PUsecase.GetAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	res, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (p *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/backend/api/v1/posts/")
	id, err := strconv.Atoi(sub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	post, err := p.PUsecase.GetById(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var dp domain.Post
	if err := json.NewDecoder(r.Body).Decode(&dp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post.Title = dp.Title
	post.Content = dp.Content

	err = p.PUsecase.Update(ctx, &post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	res, err := json.Marshal(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (p *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/backend/api/v1/posts/")
	id, err := strconv.Atoi(sub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	err = p.PUsecase.Delete(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
