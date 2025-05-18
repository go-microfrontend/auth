-- name: CreateUser :one
INSERT INTO users (email, password_hash)
VALUES ($1, $2)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users 
WHERE id = $1;

-- name: UpsertRoleByID :one
INSERT INTO roles (user_id, role)
VALUES ($1, $2)
ON CONFLICT (user_id, role)
DO UPDATE SET role = EXCLUDED.role
RETURNING *;

-- name: GetUserRoles :many
SELECT role FROM roles
WHERE user_id = $1;
