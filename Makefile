# DCC SPIRE Bridge Makefile

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=spire-dcc-attestor

all: build

build:
	@echo "Building SPIRE DCC Causal Attestor..."
	cd src && $(GOBUILD) -v -o ../bin/$(BINARY_NAME) .

test-integration:
	@echo "Running Logic Verification (Python)..."
	python3 tests/verify_spire.py

clean:
	rm -f bin/$(BINARY_NAME)

.PHONY: all build test-integration clean
