run:
	@go run cmd/app/main.go

build:
	@go build cmd/app/main.go

migrate:
	@go run cmd/migration/main.go
