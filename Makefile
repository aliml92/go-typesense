# Add the 'list' target
.PHONY: list
list:
	@echo "Available targets:"
	@awk -F: '/^[a-zA-Z0-9_-]+:.*?##/ {printf "%-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.PHONY: test
test: ## Runs all units tests.
	go test -v -race ./typesense

.PHONY: integration-test 
integration-test: ## Runs all intergration tests.
	go test -v -race -tags=integration ./test/...

.PHONY: test-coverage
test-coverage: ## Runs all unit tests + gathers code coverage.
	go test -v -race -coverprofile coverage.txt ./typesense

.PHONY: test-coverage-html
test-coverage-html: test-coverage ## Runs all unit tests + gathers code coverage + displays them in your default browser
	go tool cover -html=coverage.txt

.PHONY: lint
lint: ## Runs golangci-lint to check for  
	golangci-lint run
