run:
	@go run main.go -profile=profiles/localhost.json

test:
	@go test ./...