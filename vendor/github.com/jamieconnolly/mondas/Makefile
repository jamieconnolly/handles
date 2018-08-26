DIRS := $(shell ls -d */ | grep -v vendor/)
PACKAGES := $(shell go list ./... | grep -v /vendor/)
SOURCES := $(wildcard $(addsuffix *.go, $(DIRS)) *.go)

all: clean test

clean:
	@rm -f .coverprofile *.coverprofile */.coverprofile */*.coverprofile

deps:
	@go get
	@go get -u github.com/golang/lint/golint
	@go mod tidy

fmt:
	@gofmt -s -l $(SOURCES) | awk '{print $$1 ": file is not formatted correctly"} END{if(NR>0) {exit 1}}' 2>&1; \
	if [ $$? -eq 1 ]; then \
		echo "!!! ERROR: Gofmt found unformatted files"; \
		exit 1; \
	fi

lint: fmt vet
	@echo $(PACKAGES) | xargs -n 1 $(GOPATH)/bin/golint | awk '{print} END{if(NR>0) {exit 1}}' 2>&1; \
	if [ $$? -eq 1 ]; then \
		echo "!!! ERROR: Golint found stylistic issues"; \
		exit 1; \
	fi

test: lint
	@go list -f '"go test -v -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}}"' $(PACKAGES) | xargs -n 1 bash -c

vet:
	@go tool vet $(SOURCES) 2>&1; \
	if [ $$? -eq 1 ]; then \
		echo "!!! ERROR: Vet found suspicious constructs"; \
		exit 1; \
	fi

.PHONY: all clean deps fmt lint test vet
