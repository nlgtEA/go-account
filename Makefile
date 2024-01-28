run:
	@ templ generate
	@ npm run build
	@ go run cmd/main.go

migrateup:
	goose -dir './db/schema' postgres "$(DB)" up

migratedown:
	goose -dir './db/schema' postgres $(DB) down

migration:
	goose -dir './db/schema' create $(name) sql

sqlc:
	sqlc generate
