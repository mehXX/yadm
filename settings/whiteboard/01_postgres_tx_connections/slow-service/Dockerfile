FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod  ./

COPY . .
RUN go build -o slow-service .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/slow-service .

EXPOSE 8080

CMD ["./slow-service"]