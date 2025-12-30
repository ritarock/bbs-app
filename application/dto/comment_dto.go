package dto

import "time"

type CreateCommentInput struct {
	PostID int
	Body   string
}

type CreateCommentOutput struct {
	ID          int
	PostID      int
	Body        string
	CommentedAt time.Time
}

type GetCommentInput struct {
	ID int
}

type GetCommentOutput struct {
	ID          int
	PostID      int
	Body        string
	CommentedAt time.Time
}

type CommentItem struct {
	ID          int
	PostID      int
	Body        string
	CommentedAt time.Time
}

type ListCommentsInput struct {
	PostID int
}

type ListCommentsOutput struct {
	Comments []CommentItem
}

type UpdateCommentInput struct {
	ID   int
	Body string
}

type UpdateCommentOutput struct {
	ID          int
	PostID      int
	Body        string
	CommentedAt time.Time
}

type DeleteCommentInput struct {
	ID int
}
