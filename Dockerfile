FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o biller-microservice ./cmd/

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/biller-microservice .
COPY app.env .
COPY db/migration ./db/migration

EXPOSE 8080

ENTRYPOINT [ "/app/biller-microservice" ]