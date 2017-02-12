TEST_UNIT?=$(shell go list ./... | grep -v vendor)
export PATH := $(GOPATH)/bin:$(PATH)

all: vendor build

install:
	@echo "Installing dependencies"
	go get github.com/tools/godep
	go get github.com/Sirupsen/logrus
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega

vendor:
	@echo "Create/update vendor and godeps dir"
	godep save ./...

build:
	@echo "building..."
	go build -o build/bin/server

fmt:
	@echo "Formatting all go code..."
	go fmt `go list ./... | grep -v vendor`

test: fmt
	@echo "Running the tests"
	@go test $(TEST_UNIT) -cover
	@go vet $(TEST_UNIT) ; if [ $$? -eq 1 ]; then \
		echo "ERROR: Vet found problems in the code."; \
		exit 1; \
		fi

clean:
	@echo "Cleaning binaries..."
	@rm -rf $(GOPATH)/bin $(GOPATH)/pkg build/bin

.PHONY: all install vendor build fmt test clean