FROM golang:1.13-alpine3.10 as builder

ENV APP_NAME=RxSecure
ENV SRC_DIR=/go/src/github.com/mukesh0513/$APP_NAME

WORKDIR $SRC_DIR

COPY go.mod go.sum $SRC_DIR/

RUN go mod download

ADD . $SRC_DIR/

RUN set -eux && \
    apk add --no-cache git && \
    CGO_ENABLED=0 GOOS=linux go build -o main .

EXPOSE 8081

ENTRYPOINT ["./main"]