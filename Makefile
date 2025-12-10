.PHONY: install_completion test

files := $(shell find . -name "*.go")
version:= $(shell git describe --tags HEAD)

svlogj: ${files} go.mod Makefile
	go build -o $@

test:
	go test -v ./pkg/utils

install_completion: svlogj
	./svlogj completion bash > ~/.bash_completion.d/svlogj

# vim:ft=Make:noexpandtab
