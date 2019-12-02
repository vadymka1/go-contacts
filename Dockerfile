# FROM golang:1.12.7-alpine3.10
# RUN mkdir /app
# ADD . /app
# WORKDIR /app
# RUN go build -o main .
# CMD ["/app/main"]

FROM golang:latest AS builder

RUN mkdir -p /go/src/app

WORKDIR /go/src/app

COPY . "/work/src/github.com/go-crud"

RUN apk add --no-cache git

RUN go-wrapper download

RUN go-wrapper install

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /your-app

CMD ["go-wrapper", "run", "-web"]

EXPOSE 3000

