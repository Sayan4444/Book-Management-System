FROM golang:1.23.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/main /app/cmd/main

FROM scratch

WORKDIR /app

COPY --from=builder /app/bin/main /app/bin/main

EXPOSE 4000

CMD ["/app/bin/main"]