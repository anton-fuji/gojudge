.PHONY: build fmt

OUTPUT_DIR = ./bin

build:
	go build -o gojudge


fmt:
	go fmt ./...