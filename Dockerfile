FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o email-service .

FROM alpine:3.23
WORKDIR /app
COPY --from=builder /app/email-service .
EXPOSE 8080
ENTRYPOINT [ "./email-service" ]
