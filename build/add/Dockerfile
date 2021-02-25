FROM golang:1.13 as gozero

COPY / /go/src/gozero.io/gozero

WORKDIR /go/src/gozero.io/gozero
RUN CGO_ENABLED=0 GO111MODULE=on GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor go build -i -ldflags '-w -s' -o zero-add rpc/add/add.go

FROM alpine:3.9
RUN apk add --update ca-certificates && update-ca-certificates && apk add --no-cache tzdata
COPY --from=gozero /go/src/gozero.io/gozero/zero-add /usr/local/bin/
CMD ["sh"]