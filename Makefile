include .env

hello:
	echo "Hello"
run:
	go run main.go
watch:
	gow run main.go
test-env:
	echo ${NODE_ENV}
create-migration:
	migrate create -ext sql -dir database/migrations $(name)
migrate-up:
	migrate -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -path database/migrations up
migrate-down:
	migrate -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -path database/migrations down