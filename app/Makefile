# Variáveis
APP := ./app/cmd/http
MAIN := $(APP)
DOCS := ./app/docs

# Target padrão
default: start

# --------------------------------------------------
# Executar
# --------------------------------------------------
.PHONY: start
start:
	go run $(MAIN)

.PHONY: develop
develop:
	air -c .air.toml

# --------------------------------------------------
# Testes
# --------------------------------------------------
.PHONY: test
test:
	go test ./... -v -cover -coverprofile=coverage.out && ./scripts/filter_coverage.sh coverage.out coverage_filtered.out && go tool cover -html=coverage_filtered.out

.PHONY: test-verbose
test-verbose:
	go test ./... -v

# --------------------------------------------------
# Lint e format
# --------------------------------------------------
.PHONY: lint
lint:
	golangci-lint run

.PHONY: fmt
fmt:
	go fmt ./...

# --------------------------------------------------
# Build
# --------------------------------------------------
.PHONY: build
build:
	go build -o bin/server $(MAIN)

.PHONY: run-bin
run-bin: build
	./bin/server

# --------------------------------------------------
# Swagger
# --------------------------------------------------
.PHONY: swag
swag:
	swag init -g $(MAIN)/main.go -o $(DOCS)

# --------------------------------------------------
# Clean
# --------------------------------------------------
.PHONY: clean
clean:
	rm -f bin/server

# --------------------------------------------------
# Mocks
# --------------------------------------------------
.PHONY: mocks
generate-mocks:
	mockery --config=.mockery.yaml