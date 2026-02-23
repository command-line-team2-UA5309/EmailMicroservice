FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o email-service .

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/email-service .
EXPOSE 8080
ENTRYPOINT [ "./email-service" ]
