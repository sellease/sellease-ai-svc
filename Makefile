# VARIABLES
# =======================================================================================
# GO

include .env
export $(shell sed 's/=.*//' .env)

GOCMD=go
GOTEST=$(GOCMD) test ./...
# Machine OS
OS := $(shell uname)
WD := $(shell pwd)
OS_LOWERCASE := $(shell uname | tr '[:upper:]' '[:lower:]')

# INSTALL TARGETS
# =======================================================================================
swagger-install: # Install Go-Swagger
ifeq ($(OS), Darwin)
	brew tap go-swagger/go-swagger
	brew install go-swagger
else
	sudo apt-get install jq
	download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')
	curl -o /usr/local/bin/swagger -L'#' "$download_url"
	chmod +x /usr/local/bin/swagger
endif
	swagger version

precommit-hook:  # Golang Pre-Commit Hook Installation ##https://pre-commit.com/#cli
ifeq ($(OS), Darwin)
	brew install pre-commit
else
	sudo pip install pre-commit
endif
	pre-commit --version
	pre-commit install

golangci-lint: # Installing Magic golangci-lint
ifeq ($(OS), Darwin)
	# Run MacOS commands
	brew install golangci-lint
	brew upgrade golangci-lint
else
	# check for Linux and run other commands
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2
endif
	golangci-lint --version
	golangci-lint linters -E bodyclose
	golangci-lint linters -E gocyclo
	golangci-lint linters -E gocritic
	golangci-lint linters -E goimports
	golangci-lint linters -E goconst
	golangci-lint linters -E sqlclosecheck
	golangci-lint linters -E lll
	golangci-lint linters -E funlen
	golangci-lint linters -E godot
	golangci-lint linters -E exportloopref
	golangci-lint linters -D scopelint

# GIT PRECOMMIT INSTALL
install-precommit-hook: precommit-hook golangci-lint # Install pre-commit hook and linter

# GOTOOLS

vet: # Vet examines Go source code and reports suspicious constructs
	${GOCMD} vet

fmt: # Gofmt is a tool that automatically formats Go source code
	gofmt

test: # GO Test
	$(GOTEST)

cover: # Go Test Coverage
	${GOCMD} test -coverprofile=coverage.out ./... && ${GOCMD} tool cover -html=coverage.out

tidy: # Update Modules and Dependency Consistency
	${GOCMD} mod tidy

build: # Builds the project
	${GOCMD} build main.go

run: # Builds the project
	${GOCMD} run main.go

lint: # Lint the files
	golangci-lint run --skip-dirs docs

# SWAGGER

check-swagger:
	which swagger || (go get -u github.com/go-swagger/go-swagger/internal/swagger)

gen-swagger:  # Generate Swagger API Documentation
	swagger generate spec -o $(WD)/docs/swagger.json  --scan-models
	swagger generate spec -o $(WD)/swagger-ui/swagger.json  --scan-models

serve-swagger: check-swagger  # Serve Swagger API Documentation
	swagger validate $(WD)/docs/swagger.json
	swagger serve -F=swagger $(WD)/docs/swagger.json

swagger: gen-swagger serve-swagger	# Generate & Serve Swagger API Documentation


# MOCK
mockery-install:
ifeq ($(OS), Darwin)
    # Run MacOS commands
    brew install mockery
    brew upgrade mockery
else
	sudo apt-get update
	sudo apt-get -y install mockery
endif
	mockery --version

mockery-create-repository-sample: # Generate mock files for all repository functions in sample package
	mockery -dir=internal/repository/sample -all --output=internal/repository/sample/mocks

# MIGRATION
add-migration:
ifeq ($(filename),)
	$(error Please provide migration file name in the form of 'add-migration filename=foo_bar' type=go|sql")
else ifeq ($(type),)
	$(error Please provide migration file name in the form of 'add-migration filename=foo_bar' type=go|sql")
else
	goose -dir database/migrations/ create $(filename) $(type)
endif

migrate-up:
	goose -dir database/migrations/ postgres "postgresql://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable&search_path=$(POSTGRES_SCHEMA)" up

migrate-down:
	goose -dir database/migrations/ postgres "postgresql://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable&search_path=$(POSTGRES_SCHEMA)" down

migrate-status:
	goose -dir database/migrations/ postgres "postgresql://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable&search_path=$(POSTGRES_SCHEMA)" status
