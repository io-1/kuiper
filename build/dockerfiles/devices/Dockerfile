FROM golang:1.15.2-alpine3.12 as builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/io-1/kuiper/
COPY . .
# RUN go get -v -d ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/main cmd/devices/main.go

FROM alpine:latest
ENV DOCKERIZE_VERSION v0.12.0
RUN apk update && apk add --no-cache ca-certificates curl
RUN curl -sfL https://github.com/powerman/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-`uname -s`-`uname -m` | install /dev/stdin /usr/local/bin/dockerize
COPY --from=builder /go/bin/main /go/bin/main
# ENTRYPOINT ["/go/bin/main"] 
