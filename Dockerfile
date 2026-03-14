FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o wallet-demo ./cmd/wallet-demo

FROM debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates tzdata curl && rm -rf /var/lib/apt/lists/*

ENV TZ=Asia/Singapore
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app

COPY --from=builder /app/wallet-demo .
COPY --from=builder /app/deploy ./deploy

EXPOSE 8080

CMD ["./wallet-demo"]