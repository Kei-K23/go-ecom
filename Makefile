build:
	@go build -o bin/go-ecom cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go-ecom