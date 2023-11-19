BIN_NAME=goweb
BIN_DIR=bin
COVERAGE_DIR=coverage

# Misc
.DEFAULT_GOAL = help
.PHONY        = help test build run coverage live clean

## —— Makefile ————————————————————————————————————————————————————————————————
help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

## —— Tests ———————————————————————————————————————————————————————————————————
test: ## Run tests
	go clean -testcache
	go test ./...

coverage: ## Run tests with coverage
	@mkdir -p ${COVERAGE_DIR}
	go test -coverprofile=${COVERAGE_DIR}/coverage.out ./...
	go tool cover -html=${COVERAGE_DIR}/coverage.out

build: ## Build the binary file
	@mkdir -p ${BIN_DIR}
	go build -o ${BIN_DIR}/${BIN_NAME} .

run: build ## Run the binary file
	./${BIN_DIR}/${BIN_NAME}

live: ## Run the binary file with live reload
	air

clean: ## Remove previous build
	go clean
	rm -Rf ./${BIN_DIR}
	rm -Rf ./${COVERAGE_DIR}
