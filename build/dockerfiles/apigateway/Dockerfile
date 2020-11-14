FROM golang:1.15.2-alpine3.12 as builder
RUN apk update && apk add --no-cache git make
WORKDIR $GOPATH/src/github.com/io-1/kuiper/
COPY . .
RUN make build-bin-apigateway
RUN mv bin/apigateway /go/bin/main

FROM alpine:latest
RUN apk update && apk add --no-cache ca-certificates

# for swagger
COPY api/ /api
COPY --from=builder /go/bin/main /go/bin/main
ENTRYPOINT ["/go/bin/main"] 
