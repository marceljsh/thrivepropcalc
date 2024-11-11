all: build run

build:
	go build -o bin/thrivepropcalc cmd/main.go

run:
	./bin/thrivepropcalc

clean:
	rm -rf bin/thrivepropcalc

.PHONY: all build run clean