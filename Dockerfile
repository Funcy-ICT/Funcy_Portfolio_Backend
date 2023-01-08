FROM golang:1.18-alpine as dev
RUN apk update && \
    apk upgrade && \
    apk add bash git && \
    apk add make gcc g++ &&\
    rm -rf /var/cache/apk/*

RUN go install github.com/cespare/reflex@latest

FROM golang:1.18-alpine
COPY --from=dev /go/bin/reflex /go/bin/reflex

RUN apk update && \
    apk upgrade && \
    rm -rf /var/cache/apk/

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download


COPY . .

RUN go build -o /backend-api cmd/main.go