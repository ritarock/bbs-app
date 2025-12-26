-- name: InsertPost :execresult
INSERT INTO posts (
    title, content, created_at
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
