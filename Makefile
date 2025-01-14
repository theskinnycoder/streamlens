# Database migration commands using Goose
DB_CONNECTION ?= "postgres://postgres:postgres@localhost:5432/streamlens-dev"
MIGRATIONS_DIR ?= "internal/db/migrations"

# Start the database
start-db:
	docker compose up

# Stop the database
stop-db:
	docker compose down

# Clean up the database
clean-db:
	docker compose down -v

# Create a new migration file
migrate-create:
	@read -p "Enter migration name: " name; \
	goose -dir $(MIGRATIONS_DIR) create $$name sql

# Apply all pending migrations
migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_CONNECTION) up

# Roll back the most recent migration
migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_CONNECTION) down

# Roll back all migrations
migrate-reset:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_CONNECTION) reset

# Check current migration status
migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_CONNECTION) status

# Generate SQLC code
generate-sqlc:
	sqlc generate

# Run the auth service
run-auth:
	air -c cmd/auth/.air.toml

# Build the auth service
build-auth:
	go build -o bin/auth cmd/auth/main.go
