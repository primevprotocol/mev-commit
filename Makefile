LDFLAGS ?= -s -w \
-X github.com/primevprotocol/mev-commit.version=$(shell git describe --tags)

.PHONY: build
build: export CGO_ENABLED=0
build: bin
	go build -ldflags '$(LDFLAGS)' -o bin/mev-commit ./cmd

bin:
	mkdir $@

bufgen:
	cd rpc && buf generate -o ../
	cd messages && buf generate -o ../
