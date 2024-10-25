.DEFAULT_GOAL := all

.PHONY: all
all: clean goapp test

.PHONY: goapp
goapp:
	mkdir -p bin
	go build -o bin ./...

.PHONY: clean
clean:
	go clean
	rm -f bin/*

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: bench
bench:
	go test -v ./... -bench=. -run=xxx -benchmem