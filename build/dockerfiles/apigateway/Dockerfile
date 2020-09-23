FROM golang:1.15.2-alpine3.12 as builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/io-1/kuiper/
COPY . .
# RUN go get -v -d ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/main cmd/apigateway/main.go

FROM alpine:latest
RUN apk update && apk add --no-cache ca-certificates

# for swagger
COPY api/ /api
COPY --from=builder /go/bin/main /go/bin/main
ENTRYPOINT ["/go/bin/main"] 
