install_dependencies:
	go get -u go.uber.org/zap
	go get -u github.com/natefinch/lumberjack
	go get -u github.com/rakyll/gotest

setup_develop:
	cp ./scripts/pre-commit.sh ./.git/hooks/pre-commit

setup_all: install_dependencies setup_develop

test_all:
	rm -rf ./coverage
	mkdir -p ./coverage
	gotest -v -cover -coverprofile ./coverage/cover.out ./...
	go tool cover -html=./coverage/cover.out -o ./coverage/cover.html
