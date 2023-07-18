FROM golang:1.21rc3-alpine3.18 AS builder

WORKDIR /app

COPY . .

RUN apk add curl
RUN apk add --no-cache make
RUN go mod download
RUN go mod verify
RUN go build -o api cmd/api/*

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz


FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/api .
COPY --from=builder /app/migrate ./migrate
COPY migrations ./migrations
COPY .env .
COPY start.sh .
COPY wait-for.sh .

EXPOSE 8000
CMD ["/app/api"]
ENTRYPOINT [ "/app/start.sh" ]