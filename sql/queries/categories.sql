-- name: ListCategories :many
SELECT * FROM categories;

-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1;

-- name: CreateCategory :exec
INSERT INTO categories (id, name, description)
VALUES ($1, $2, $3);

-- name: UpdateCategory :exec
UPDATE categories SET name = $2, description = $3
WHERE id = $1;
