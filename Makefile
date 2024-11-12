BINARY_NAME=thrivepropcalc
SOURCE_DIR=cmd
INPUT_FILE=input.txt
DOCKER_IMAGE=thrivepropcalc

all: build

build:
	@echo "Building the Go application..."
	go build -o bin/$(BINARY_NAME) $(SOURCE_DIR)/main.go

run: build
	@echo "Running the Go application..."
	./bin/$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	rm -f bin/$(BINARY_NAME)

docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	@echo "Running Docker container..."
	docker run --rm $(DOCKER_IMAGE)

.PHONY: all build run clean test docker-build docker-run