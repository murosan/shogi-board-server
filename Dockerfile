FROM golang:1.11

ENV GO111MODULE=on TZ=Asia/Tokyo

WORKDIR /go/src/shogi-board-server

COPY . .
COPY ./config/lotate /etc/logrotate.d/shogi-board-server

RUN mkdir -p /go/src/shogi-board-server/log
CMD ["go", "run", "main.go", "-port", "8081", "-appConfig", "./config/app_docker.yml", "-logConfig", "./config/log_docker.yml"]

EXPOSE 8081
