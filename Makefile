
GO := GO111MODULE=on go
GOPATH ?= $(shell $(GO) env GOPATH)
GOBIN ?= $(GOPATH)/bin
GOLINT ?= $(GOBIN)/golint
GOSTATIC ?= $(GOBIN)/staticcheck

deps:
	go mod tidy
	go mod vendor

lint:
	@$(GOSTATIC) ./... | grep -v vendor/ && exit 1 || exit 0
	@$(GOLINT) ./... | grep -v vendor/ && exit 1 || exit 0

rpc:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  \
	proto/l2_v1/l2_v1.proto