FROM golang:1.24.2-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main .

FROM alpine:3.21.3
WORKDIR /root/
COPY --from=builder /main .
EXPOSE 8080
CMD ["./main"]