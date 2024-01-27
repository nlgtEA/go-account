run:
	@ templ generate
	@ npm run build
	@ go run cmd/main.go
