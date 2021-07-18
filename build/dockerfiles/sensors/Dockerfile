FROM golang:1.15.2-alpine3.12 as builder
RUN apk update && apk add --no-cache git make
WORKDIR $GOPATH/src/github.com/io-1/kuiper/
COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/main cmd/sensors/main.go
RUN make build-bin-sensors
RUN mv bin/sensors /go/bin/main

FROM alpine:latest
RUN apk update && apk add --no-cache ca-certificates
COPY --from=builder /go/bin/main /go/bin/main
ENTRYPOINT ["/go/bin/main"] 
