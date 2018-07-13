.PHONY: run build clean test help default

BIN_NAME:=ctree



VERSION:=${shell git rev-parse HEAD}


default: run


run:
	go run main.go


build:
	go build -o bin/${BIN_NAME}


version:
	@echo ${VERSION}

	
