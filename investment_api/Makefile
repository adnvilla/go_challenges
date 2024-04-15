SHELL := /bin/bash

default: build

build: 
	go build -o bin/investment_api main.go 

run:
	./bin/investment_api

test:
	go test ./... -cover

mock_gen:	
	mockery