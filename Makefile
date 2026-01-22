MODULE_NAME := zntfaker
VERSION ?= $(shell cat VERSION 2>/dev/null || echo "v0.0.0")

.PHONY: lint test bench build tag

lint: ## Run linter
	@echo "\n*** Lint started ***"
	golangci-lint run ./...
	@echo "*** Lint finished ***"

test: ## Run tests
	@echo "\n*** Test started ***"
	go test -v --race ./...
	@echo "*** Test finished ***"

bench:
	@echo "\n*** Benchmark started ***"
	go test -bench=. -benchmem -benchtime=3s ./...
	@echo "*** Benchmark finished ***"

build: ## Build the module
	@echo "\n*** Building probe $(MODULE_NAME) $(VERSION) ***"
	go build -v ./...
	@echo "*** Build finished ***"

tag: ## Create git tag if version changed
	@echo "\n*** Checking version ***"
	@if [ "$(VERSION)" = "v0.0.0" ]; then \
		echo "❌ ERROR: VERSION is empty in VERSION file"; \
		exit 1; \
	fi
	git checkout production
	git pull origin production
	@LAST_TAG=$$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0"); \
	if [ "$$LAST_TAG" = "$(VERSION)" ]; then \
		echo "❌ ERROR: Version $(VERSION) already exists on production"; \
		exit 1; \
	else \
		echo "✅ Previous: $$LAST_TAG. Trying to push current: $(VERSION)"; \
		git tag $(VERSION); \
		git push origin $(VERSION); \
		echo "*** Tag $(VERSION) pushed ***"; \
	fi
