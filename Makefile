.DEFAULT_GOAL := build

# Build app
build:
	docker-compose build
.PHONY: build

# Start app
start:
	docker-compose up -d
.PHONY: start

# Stop app
stop:
	docker-compose down
.PHONY: stop

#  Execute tests
test:
	docker-compose exec api go test ./...
.PHONY: test
