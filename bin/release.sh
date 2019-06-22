#!/bin/bash

read -p "Version (ex. 1.0.0) > " VERSION

DIR="./releases/sbserver-${VERSION}-macOS"

rm -rf ${DIR} ${DIR}.zip
mkdir -p ${DIR}

go build -o ${DIR}/sbserver-${VERSION}

cp ./bin/start.sh ${DIR}/start
cp ./config/app.example.yml ${DIR}/app.yml

zip -9 -r ${DIR}.zip ${DIR}

rm -rf ${DIR}
