FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o pinger ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/pinger .

ENTRYPOINT ["./pinger"]