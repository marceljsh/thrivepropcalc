all: build run

build:
	go build -o bin/propcalc cmd/main.go

run:
	./bin/propcalc

clean:
	rm -rf bin/propcalc

.PHONY: all build run clean