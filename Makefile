check:
	$(CURDIR)/scripts/check.sh

go: check
	# Standalone GOPATH
	$(CURDIR)/scripts/generate_go.sh
	GO111MODULE=on go mod tidy
	GO111MODULE=on go build ./pkg/...

build: main.go
	go build

.DEFAULT_GOAL := build
