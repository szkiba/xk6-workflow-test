all: lint security test doc build it

lint:
	golangci-lint run ./...

security:
	gosec -quiet ./...
	govulncheck ./...

test:
	go test -cover -race ./...

doc:
	bun x typedoc --out build/docs

build:
	xk6 build --with github.com/szkiba/xk6-workflow-test=.

it: build
	./k6 run test/smoke.test.js

.PHONY: all lint security test doc build it
.SILENT:
