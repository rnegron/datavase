.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY: fmt

lint: fmt
	staticcheck ./...
.PHONY: lint

vet: lint
	go vet ./...
.PHONY:	vet

build: vet
	go build cmd/dv.go
.PHONY:	build
