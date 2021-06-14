BINARY=create-contract-aws-lambda
CHECKSUMFILE=SHA256SUM
all: build

deps:
	@go mod tidy

test:
	@go test -v -cover ./...

lint:
	@go vet ./...

build: test
	@env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ${BINARY}
	@shasum -a256 ${BINARY} > ${CHECKSUMFILE}
	@echo "binary file ${BINARY}"

clean:
	@rm -frv {${BINARY}.*,${CHECKSUMFILE},${BINARY}}

compress pack deploy: build
	@zip ${BINARY}.zip ${BINARY}
	@echo "${BINARY}.zip"

.PHONY: build