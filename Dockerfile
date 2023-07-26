FROM golang:1.20

WORKDIR /usr/src/app

RUN go mod init example/goapi

RUN go get github.com/gofiber/fiber/v2
# RUN go install github.com/cosmtrek/air@latest

COPY . .

# RUN go build ./main.go
RUN go build -v -o /usr/local/bin/app ./...