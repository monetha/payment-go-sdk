PKGS ?= $(shell glide novendor)

.PHONY: dependencies
dependencies:
	@echo "Installing Glide and locked dependencies..."
	glide --version || go get -u -f github.com/Masterminds/glide
	glide install
	@echo "Installing goimports..."
	go install ./vendor/golang.org/x/tools/cmd/goimports
	@echo "Installing golint..."
	go install ./vendor/golang.org/x/lint/golint
	@echo "Installing staticcheck..."
	go install ./vendor/honnef.co/go/tools/cmd/staticcheck

.PHONY: lint
lint:
	@echo "Checking formatting..."
	@gofiles=$$(go list -f {{.Dir}} $(PKGS) | grep -v mock) && [ -z "$$gofiles" ] || unformatted=$$(for d in $$gofiles; do goimports -l $$d/*.go; done) && [ -z "$$unformatted" ] || (echo >&2 "Go files must be formatted with goimports. Following files has problem:\n$$unformatted" && false)
	@echo "Checking vet..."
	@go vet $(PKGS)
	@echo "Checking staticcheck..."
	@staticcheck $(PKGS)
	@echo "Checking lint..."
	@$(foreach dir,$(PKGS),golint $(dir);)