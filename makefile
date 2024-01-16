BINARY=life

.DEFAULT: ${BINARY}

${BINARY}:
	go build -o ${BINARY}

build: ${BINARY}

run: build
	./${BINARY} --board-size-x=12 --board-size-y=12 --cycles=10

# snapshot-build:
# 	goreleaser build --single-target --snapshot --rm-dist -o osv-detector

test:
	go test ./...

coverage-test:
	go test ./... -cover

lint:
# 	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2 run ./... --max-same-issues 0
	gofmt -s -d */**.go
