-- name: CreateTenant :one
INSERT INTO tenants (name)
VALUES ($1)
RETURNING id, name, created_at, updated_at;

-- name: GetTenantByID :one
SELECT id, name, created_at, updated_at
FROM tenants
WHERE id = $1;

-- name: GetTenantByName :one
SELECT id, name, created_at, updated_at
FROM tenants
WHERE name = $1;

-- name: CreateUser :one
INSERT INTO users (tenant_id, email, hashed_password)
VALUES ($1, $2, $3)
RETURNING id, tenant_id, email, created_at, updated_at;

-- name: AssignUserRole :one
INSERT INTO user_roles (user_id, role)
VALUES ($1, $2)
RETURNING id, user_id, role, created_at, updated_at;

-- name: GetUserByEmail :one
SELECT id, tenant_id, email, hashed_password, created_at, updated_at
FROM users
WHERE email = $1;

-- name: GetUserRoles :many
SELECT id, user_id, role, created_at, updated_at
FROM user_roles
WHERE user_id = $1;
