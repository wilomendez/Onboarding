BINARY=notify-aws-lambda
CHECKSUMFILE=SHA256SUM
all: build

deps:
	@go mod tidy

test:
	@go test -v -cover ./...

lint:
	@go vet ./...

build: test
	@export GO111MODULE=on
	# FIX: Add CGO_ENABLED=0 when building go files
	@env CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ${BINARY}
	@shasum -a256 ${BINARY} > ${CHECKSUMFILE}
	@echo "binary file ${BINARY}"

clean:
	@rm -frv {${BINARY}.*,${CHECKSUMFILE},${BINARY}}

compress pack deploy: build
	@zip ${BINARY}.zip ${BINARY}
	@echo "${BINARY}.zip"

.PHONY: build