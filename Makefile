install_development:
	cp ./scripts/pre-commit.sh ./.git/hooks/pre-commit
	go get -u github.com/rakyll/gotest
	go get -u golang.org/x/lint/golint
	go get -u honnef.co/go/tools/cmd/staticcheck

build:
	go build -o ./main

test_all:
	rm -rf ./coverage
	mkdir -p ./coverage
	gotest -v -cover -coverprofile ./coverage/cover.out ./...
	go tool cover -html=./coverage/cover.out -o ./coverage/cover.html
