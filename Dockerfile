FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk --no-cache add build-base

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=1
RUN go build -o app ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/app .

COPY --from=builder /app/cmd/server/.env .

EXPOSE 8080

CMD ["./app"]
