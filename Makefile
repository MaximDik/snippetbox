.PHONY: build

build:
	go build -v ./cmd/web/main.go

start:
	./main



.DEFAULT_GOAL :=  build