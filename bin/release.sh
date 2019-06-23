#!/bin/bash

read -p "Version (ex. 1.0.0) > " VERSION

PROJECT_ROOT=$(cd $(dirname $0)/..; pwd)
RELEASE=$PROJECT_ROOT/releases

rm -rf $RELEASE
mkdir $RELEASE

MAC_DIR="sbserver-$VERSION-macOS"
WINDOWS_DIR="sbserver-$VERSION-Windows-x64"

export GOARCH=amd64

function build() {
  local dir=$1
  local conf_dir="$RELEASE/$dir/config"
  local bin_name="sbserver"
  if [[ $dir = $WINDOWS_DIR ]]; then
    bin_name+=".exe"
  fi

  cd $PROJECT_ROOT

  mkdir $RELEASE/$dir $conf_dir
  go build -o $RELEASE/$dir/$bin_name
  cp ./config/app.example.yml $conf_dir/app.config.yml

  cd $RELEASE
  tar cvzf $dir.tar.gz $dir
  rm -rf $dir
}

GOOS=darwin build $MAC_DIR
GOOS=windows build $WINDOWS_DIR
