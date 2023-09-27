# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean

# Binary name for your Go program
BINARY_NAME=eurovote

all: build run

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

run:
	$(GORUN) .

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

fclean: clean run