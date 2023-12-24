# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean

# Binary name for your Go program
BINARY_DIR=bin

help:
	@echo "Available targets:"
	@echo "  all        : Build and run the program"
	@echo "  build      : Build the program"
	@echo "  run        : Run the program"
	@echo "  clean      : Clean the build files"
	@echo "  fclean     : Clean the build files and run the program"
	@echo "  start      : Start the Docker containers"
	@echo "  logs       : Show the logs of the Docker containers"
	@echo "  restart    : Restart the Docker containers"
	@echo "  logs-real-time : Show the real-time logs of the Docker containers"
	@echo "  logs-service   : Show the logs of a specific service in the Docker containers"

.PHONY: all build run clean fclean start logs restart logs-real-time logs-service help



all: build run

build:
	$(GOBUILD) -o $(BINARY_DIR) -v ./...

run:
	./bin/api

clean:
	$(GOCLEAN)
	rm -f $(BINARY_DIR)

fclean: clean run

start:
	docker-compose up -d

logs:
	docker-compose logs

restart:
	docker-compose restart

logs-real-time:
	docker-compose logs -f

logs-service:
	docker-compose logs $(service)


