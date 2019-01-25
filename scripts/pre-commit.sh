#!/bin/bash -x

# https://tip.golang.org/misc/git/pre-commit

gofiles=$(git diff --cached --name-only --diff-filter=ACM | grep '\.go$')
[ -z "$gofiles" ] && exit 0

set -e

unformatted=$(gofmt -l $gofiles)
[ -z "$unformatted" ] && exit 0

for fn in $unformatted; do
  gofmt -w $PWD/$fn
  git add $PWD/$fn
done

go vet ./...
staticcheck ./...
golint -set_exit_status ./...
