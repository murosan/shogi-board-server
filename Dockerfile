FROM golang:1.18

ENV TZ=Asia/Tokyo

WORKDIR /go/src/shogi-board-server

COPY . .

CMD go run main.go \
      -port 8081 \
      -app_config ./config/app.docker.yml \
      -log_config ./config/log.docker.yml

EXPOSE 8081
