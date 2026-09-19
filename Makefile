.PHONY: migrate-new migrate-hash run
migrate-new:
	@test -n "$(NAME)" || (echo "Usage: migrate-new NAME=add_something"; exit 1)
	GOWORK=off atlas migrate new $(NAME)
migrate-hash:
	@atlas migrate hash
migrate-apply:
	@atlas migrate apply --dir "file://migrations/" --url "sqlite://data/dmc.db"
run:
	@go run ./cmd/api/main.go
build:
	@go build ./...