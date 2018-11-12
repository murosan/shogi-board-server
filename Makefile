install_dependencies:
	go get -u go.uber.org/zap
	go get -u github.com/natefinch/lumberjack

develop:
	cp ./scripts/pre-commit.sh ./.git/hooks/pre-commit

all: install_dependencies develop
