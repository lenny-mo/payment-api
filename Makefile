
GOPATH:=$(shell go env GOPATH)
MODIFY= proto/

.PHONY: proto
proto:
    
	protoc  --micro_out=${MODIFY} --go_out=${MODIFY} proto/payment-api.proto
    

.PHONY: build
build: proto

	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o payment-api *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t payment-api:latest
