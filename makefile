ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: build
build:
	go build -v .

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

migrate-up: 
	migrate -path migrations -database "postgres://$(POSTGRES_USER):${POSTGRES_PASSWORD}@localhost/orders-service-db?sslmode=disable" up

.DEFAULT_GOAL := build