-- name: InsertPost :execresult
INSERT INTO posts (
    title, content, posted_at
) VALUES (
    ?, ?, ?
);

-- name: SelectPost :one
SELECT * FROM posts
WHERE id = ?
;

-- name: SelectPosts :many
SELECT * FROM posts
ORDER BY id
;

-- name: UpdatePost :exec
UPDATE posts
SET title = ?, content = ?
WHERE id = ?
;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = ?
;

-- name: InsertComment :execresult
INSERT INTO comments (
    post_id, body, commented_at
) VALUES (
    ?, ?, ?
);

-- name: SelectComment :one
SELECT * FROM comments
WHERE id = ?
;

-- name: SelectCommentsByPostId :many
SELECT * FROM comments
WHERE post_id = ?
ORDER BY id
;

-- name: UpdateComment :exec
UPDATE comments
SET body = ?
WHERE id = ?
;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = ?
