
# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# Create the new confirm target.
.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/test: runs go test with default values
.PHONY: run/test
run/test:
	go test -timeout 300s -v -count=1 -race ./...

## run/update: runs go get -u && go mod tidy
.PHONY: run/update
run/update:
	go get -u ./...
	go mod tidy

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	#staticcheck ./...  # go install honnef.co/go/tools/cmd/staticcheck@latest
	@echo 'Running tests...'
	go test -race -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/lib: build the package
.PHONY: build/lib
build/lib: audit
	@echo 'Building cmd/...'
	go build 

## build/dlv-debug: build the application with dlv gcflags
.PHONY: build/dlv-debug
build/dlv-debug: audit
	@echo "Building for delve debug..."
	@go build \
	-ldflags=-compressdwarf=false \
	-gcflags=all=-d=checkptr \
	-gcflags="all=-N -l" \
	-o ./bin/cmd ./cmd
