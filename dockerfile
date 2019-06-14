FROM golang:alpine as builder

RUN apk add --update --no-cache git

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 cd server && go build -o ../main

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main /app/

EXPOSE 8080

CMD ["/app/main"]
