VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"
MAKEFLAGS += --silent
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

.PHONY: get 
get:
	echo "getting go dependencies..."
	# go get -v -d ./...
	# go get -u github.com/stretchr/testify
	# go get -u github.com/jstemmer/go-junit-report
	# go get -u github.com/axw/gocov/...
	# go get -u github.com/AlekSi/gocov-xml
	go mod download -x
	echo "done"

.PHONY: get-ci
get-ci:
	echo "getting ci go dependencies..."
	go get -u github.com/jstemmer/go-junit-report
	go get -u github.com/axw/gocov/...
	go get -u github.com/AlekSi/gocov-xml
	echo "done"

.PHONY: generate
generate:
	echo "generating dependency files..."
	protoc --go-grpc_out=pkg/pb/sensors --go_out=pkg/pb/sensors pkg/pb/sensors/sensors.proto
	protoc --go-grpc_out=pkg/pb/users --go_out=pkg/pb/users pkg/pb/users/users.proto
	protoc --go-grpc_out=pkg/pb/devices --go_out=pkg/pb/devices pkg/pb/devices/devices.proto
	protoc --go-grpc_out=pkg/pb/interactions --go_out=pkg/pb/interactions --proto_path=pkg/pb/interactions pkg/pb/interactions/interactions.proto
	mockgen -source pkg/pb/devices/devices_grpc.pb.go -destination=pkg/mock/mockdevicesserviceclient.go -package=mock
	mockgen -source pkg/pb/users/users_grpc.pb.go -destination=pkg/mock/mockusersserviceclient.go -package=mock
	go generate ./...
	echo "done"

.PHONY: test-unit
test-unit:
	echo "running unit tests..."
	CGO_ENABLED=0 go test -v -tags unit ./...
	echo "done"

.PHONY: version
version:
	echo "setting versions"
	sed -i 's/version:.*/version: $(VERSION)/g' ./api/swagger.yaml
	echo "done"

# .PHONY: test-integration
# test-integration: build-test start-test wait-test run-test wait-test stop-test

# .PHONY: test
# test:
# test: test-unit test-integration test-benchmark

.PHONY: cover-unit
cover-unit:
	go test -tags unit -v ./... -coverprofile c.out; go tool cover -func c.out

.PHONY: cover-unit-html
cover-unit-html:
	go test -tags unit -v ./... -coverprofile c.out; go tool cover -html c.out

.PHONY: cover-unit-html-file
cover-unit-html-file:
	go test -tags unit -v ./... -coverprofile c.out; go tool cover -html c.out -o coverage.html

.PHONY: lint
lint:
	golint ./...

.PHONY: clean
clean: 
	go clean -cache

.PHONY: clean-docker
clean-docker: 
	echo "cleaning docker"
	dcp down
	docker rmi
	echo "done"

.PHONY: test-build
test-build:
	echo "building test db..."
	docker build -t "$(PROJECTNAME)"/test-db:"$(VERSION)" --label "version"="$(VERSION)" --label "build"="$(BUILD)" -f build/dockerfiles/db/devices/Dockerfile.test build/dockerfiles/db/devices/.
	echo "done"

.PHONY: test-start
test-start:
	echo "start test db..."
	docker run -it --name test-db -p 3306:3306 -d "$(PROJECTNAME)"/test-db:"$(VERSION)"
	echo "done"

.PHONY: test-run
test-run:
	DB_CONN="root:password@tcp(127.0.0.1:3306)/devices?charset=utf8&parseTime=True&loc=Local" go test -tags integration -v ./...
	echo "done"

.PHONY: test-clean
test-clean:
	echo "cleaning test db..."
	docker rm -f test-db 
	docker rmi "$(PROJECTNAME)"/test-db:"$(VERSION)"
	echo "done"

.PHONY: build-bin-apigeteway
build-bin-apigateway:
	echo "building apigateway binary..."
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo $(LDFLAGS) -o bin/apigateway cmd/apigateway/main.go
	echo "done"

.PHONY: build-bin-sensors
build-bin-sensors:
	echo "building sensors binary..."
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo $(LDFLAGS) -o bin/sensors cmd/sensors/main.go
	echo "done"

.PHONY: build-bin-devices
build-bin-devices:
	echo "building devices binary..."
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo $(LDFLAGS) -o bin/devices cmd/devices/main.go
	echo "done"

.PHONY: build-bin-users
build-bin-users:
	echo "building users binary..."
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo $(LDFLAGS) -o bin/users cmd/users/main.go
	echo "done"

.PHONY: build-bin-interactions
build-bin-interactions:
	echo "building interactions binary..."
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo $(LDFLAGS) -o bin/interactions cmd/interactions/main.go
	echo "done"

.PHONY: build-apigeteway
build-apigateway:
	echo "building apigateway..."
	docker build -t "$(PROJECTNAME)"/apigateway:"$(VERSION)" --label "version"="$(VERSION)" --label "build"="$(BUILD)" -f build/dockerfiles/apigateway/Dockerfile .
	echo "done"

.PHONY: build-sensors
build-sensors:
	echo "building sensors..."
	docker build -t "$(PROJECTNAME)"/sensors:"$(VERSION)" --label "version"="$(VERSION)" --label "build"="$(BUILD)" -f build/dockerfiles/apigateway/Dockerfile .
	echo "done"

.PHONY: build-devices
build-devices:
	echo "building devices..."
	docker build -t "$(PROJECTNAME)"/devices:"$(VERSION)" --label "version"="$(VERSION)" --label "build"="$(BUILD)" -f build/dockerfiles/devices/Dockerfile .
	echo "done"

.PHONY: build-users
build-users:
	echo "building users..."
	docker build -t "$(PROJECTNAME)"/users:"$(VERSION)" --label "version"="$(VERSION)" --label "build"="$(BUILD)" -f build/dockerfiles/users/Dockerfile .
	echo "done"

.PHONY: build-interactions
build-interactions:
	echo "building interactions..."
	docker build -t "$(PROJECTNAME)"/interactions:"$(VERSION)" --label "version"="$(VERSION)" --label "build"="$(BUILD)" -f build/dockerfiles/interactions/Dockerfile .
	echo "done"

.PHONY: build-all
build-all: build-apigateway build-sensors build-devices build-users build-interactions

.PHONY: up
up: 
	docker-compose up -d "$(PROJECTNAME)"/apigateway:"$(VERSION)"
	docker-compose up -d "$(PROJECTNAME)"/sensors:"$(VERSION)"
	docker-compose up -d "$(PROJECTNAME)"/devices:"$(VERSION)"
	docker-compose up -d "$(PROJECTNAME)"/users:"$(VERSION)"

.PHONY: help
help:
	echo "Choose a command run in $(PROJECTNAME):"
