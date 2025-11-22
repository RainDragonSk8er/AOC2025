.PHONY: build install scaffold docker-build

build:
	go build -o bin/tracker ./cmd/tracker
	go build -o bin/scaffold ./cmd/scaffold

scaffold:
	@if [ -z "$(DAY)" ]; then echo "Usage: make scaffold DAY=X"; exit 1; fi
	go run ./cmd/scaffold -day $(DAY)

docker-build:
	docker build -f docker/Dockerfile -t aoc-tracker .

test:
	go test ./pkg/...
