# ---------- Stage 1: Build ----------
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o calculator

# ---------- Stage 2: Run ----------
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/calculator .

CMD ["./calculator"]