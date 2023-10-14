
-- name: CreateCourses :exec
INSERT INTO courses (id, name, description, category_id, price)
VALUES ($1, $2, $3, $4, $5);

-- name: ListCourses :many
SELECT c.*, ca.name as category_name FROM courses c
JOIN categories ca ON ca.id = c.category_id;

-- name: GetCourse :one
SELECT c.*, ca.name as category_name FROM courses c
JOIN categories ca ON ca.id = c.category_id
WHERE c.id = $1;
