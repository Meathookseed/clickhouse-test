RUN=go run main.go

migrate:
	${RUN} migrations:execute
.PHONY: migrate

pre-commit-check:
	go mod tidy
	golangci-lint run -v
.PHONY: pre-commit-check

lint:
	golangci-lint run -v --fix
.PHONY: lint

load-test:
	bombardier -c 200 -n 10000 -m POST -f test.json  0.0.0.0:9191/game-event
.PHONY: bombardier-load-test