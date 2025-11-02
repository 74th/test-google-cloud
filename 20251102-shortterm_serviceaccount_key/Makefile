.PHONY: build create-shrt-term-key/create-shrt-term-key
.DEFAULT_GOAL := build

create-shrt-term-key/create-shrt-term-key: create-short-term-key/main.go
	cd create-short-term-key; go build .

build: create-shrt-term-key/create-shrt-term-key
