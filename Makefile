GIT_DIR=$(shell git rev-parse --git-dir)

setup:
	cp ./scripts/pre-commit.sh $(GIT_DIR)/hooks/pre-commit
	git config commit.template .commit-template

build:
	go build -o ./main

test:
	gotest . ./app/...

test-c:
	rm -rf ./coverage
	mkdir -p ./coverage
	gotest -v -cover -coverprofile ./coverage/cover.out . ./app/...
	go tool cover -html=./coverage/cover.out -o ./coverage/cover.html
	open ./coverage/cover.html

check:
	go fmt . ./app/...
	go vet . ./app/...
	golint . ./app/...
	# honnef.co/go/tools/cmd/staticcheck
	staticcheck . ./app/...

# clean up go modules
clean_modules:
	go mod tidy
