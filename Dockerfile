FROM golang:1.26.8-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o dmc ./cmd/api/main.go

FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/dmc .
COPY --from=builder /app/data/dmc.db /app/data/dmc.db
EXPOSE 3000
CMD ["./dmc"]