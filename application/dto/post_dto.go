package dto

import "time"

type CreatePostInput struct {
	Title   string
	Content string
}

type CreatePostOutput struct {
	ID       int
	Title    string
	Content  string
	PostedAt time.Time
}

type GetPostInput struct {
	ID int
}

type GetPostOutput struct {
	ID       int
	Title    string
	Content  string
	PostedAt time.Time
}

type PostItem struct {
	ID       int
	Title    string
	Content  string
	PostedAt time.Time
}

type ListPostOutput struct {
	Posts []PostItem
}

type UpdatePostInput struct {
	ID      int
	Title   string
	Content string
}

type UpdatePostOutput struct {
	ID       int
	Title    string
	Content  string
	PostedAt time.Time
}

type DeletePostInput struct {
	ID int
}
