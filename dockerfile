FROM golang:1.26-alpine AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/bot ./cmd/bot


FROM alpine:3.21

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /app/bot /app/bot

ENV TZ=Asia/Ho_Chi_Minh

ENTRYPOINT ["/app/bot"]
