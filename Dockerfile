FROM golang:1.23.0

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/main ./cmd/main

EXPOSE 4000

CMD ["./.bin/main"]