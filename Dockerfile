FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod download \
    && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o main -ldflags="-s -w" .

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .
COPY config.yaml ./

CMD ["./main"]
