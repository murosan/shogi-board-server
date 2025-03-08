GIT_DIR = $(shell git rev-parse --git-dir)
GOPATH = $(shell go env GOPATH)
GOBIN = $(GOPATH)/bin

TEST_ARGS = . ./app/...

install:
	go install

install-tools:
	go mod download
	go install honnef.co/go/tools/cmd/staticcheck@2025.1.1

setup:
	cp ./bin/pre-commit.sh $(GIT_DIR)/hooks/pre-commit
	git config commit.template .commit-template

clean:
	go mod tidy
	go clean

build:
	go build -ldflags="-s -w" -trimpath -o ./main

test:
	go test $(TEST_ARGS)

test-c:
	rm -rf ./coverage
	mkdir -p ./coverage
	gotest -v -cover -coverprofile ./coverage/cover.out $(TEST_ARGS)
	go tool cover -html=./coverage/cover.out -o ./coverage/cover.html
	open ./coverage/cover.html

lint: install-tools
	go fmt $(TEST_ARGS)
	go vet $(TEST_ARGS)
	# honnef.co/go/tools/cmd/staticcheck
	$(GOBIN)/staticcheck $(TEST_ARGS)
	# $(GOBIN)/golint -set_exit_status $(TEST_ARGS)

ci: clean install-tools
	go vet $(TEST_ARGS)
	$(GOBIN)/staticcheck $(TEST_ARGS)
	# $(GOBIN)/golint -set_exit_status $(TEST_ARGS)
	go test -race $(TEST_ARGS)

release:
	./bin/release.sh
