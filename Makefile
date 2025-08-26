APP_NAME := auth_service
CMD_DIR := cmd/server

init:
	@if [ ! -f go.mod ]; then \
		echo "Initializing Go module..."; \
		go mod init github.com/michaelrodriguess/$(APP_NAME); \
	else \
		echo "go.mod already exists, skipping init"; \
	fi

tidy:
	@echo "Tidying dependencies..."
	go mod tidy

deps: tidy

build:
	@echo "Building $(APP_NAME)..."
	go build -o $(APP_NAME) ./$(CMD_DIR)

doc:
	@echo "Generating documentation..."
	swag init -g $(CMD_DIR)/main.go

run:
	@echo "Running $(APP_NAME)..."
	go run ./$(CMD_DIR)/main.go

clean:
	@echo "Cleaning compiled files..."
	rm -f $(APP_NAME)

reload: tidy run

all: init tidy build run

.PHONY: init tidy deps build doc run clean reload all
