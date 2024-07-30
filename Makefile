.PHONY: build

build:
	go build -v ./cmd/web/

start:
	./web -addr="127.0.0.1:9999"



.DEFAULT_GOAL :=  build