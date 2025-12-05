.PHONY: install_completion

files := $(shell find . -name "*.go")

svlogj: ${files} go.mod Makefile
	go build -o $@

install_completion: svlogj
	./svlogj --generate-completion=bash > ~/.bash_completion.d/svlogj

# vim:ft=Make:noexpandtab
