BINARY=go_with_tests
.DEFAULT_GOAL := run

build:
	@go build -o bin/${BINARY}

run:build
	@./bin/${BINARY}

clean:
	rm -rf bin

test:
	go test -v -cover ./...


.PHONY: test