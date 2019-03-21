GIT_DIR=$(shell git rev-parse --git-dir)

setup_development:
	cp ./scripts/pre-commit.sh $(GIT_DIR)/hooks/pre-commit
	git config commit.template .commit-template

build:
	go build -o ./main

test_all:
	rm -rf ./coverage
	mkdir -p ./coverage
	gotest -v -cover -coverprofile ./coverage/cover.out . ./app/...
	go tool cover -html=./coverage/cover.out -o ./coverage/cover.html

check_code_style:
	go vet . ./app/...
	staticcheck . ./app/...
	golint . ./app/...

# clean up go modules
clean_modules:
	go mod tidy
