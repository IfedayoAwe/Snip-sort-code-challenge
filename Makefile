# Go compiler
GOCMD := go
GOTEST := $(GOCMD) test
GOBUILD := $(GOCMD) build
GOFMT := $(GOCMD) fmt
BINARY_NAME := products

# List sorting methods
SORTERS := "Price, SalesViewRatio, CreationDate"

.PHONY: all run test fmt sorters

all: fmt build

# Run the program
run:
	$(GOCMD) run .

# Run tests
test:
	$(GOTEST) -v ./...

# Format code
fmt:
	$(GOFMT) ./...

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_NAME) main.go

# Show sorting methods
sorters:
	@echo "Available sorting methods: $(SORTERS)"
