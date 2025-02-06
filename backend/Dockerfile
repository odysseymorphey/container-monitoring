FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o server ./cmd/api/main.go
RUN ls -l /app/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

ENTRYPOINT ["./server"]