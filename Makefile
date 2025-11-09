.PHONY: build run test clean help

# Variables
BINARY_NAME=calculator-api
CMD_PATH=./cmd/server

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the application
	go build -o $(BINARY_NAME) $(CMD_PATH)

run: ## Run the application
	go run $(CMD_PATH)/main.go

test: ## Run tests
	go test -v ./...

test-cover: ## Run tests with coverage
	go test -cover ./...

clean: ## Remove build artifacts
	rm -f $(BINARY_NAME)
	go clean

deps: ## Download dependencies
	go mod download
	go mod tidy

fmt: ## Format code
	go fmt ./...

lint: ## Run linter
	golangci-lint run

docker-build: ## Build Docker image
	docker build -t $(BINARY_NAME):latest .

docker-run: ## Run Docker container
	docker run -p 8080:8080 $(BINARY_NAME):latest
