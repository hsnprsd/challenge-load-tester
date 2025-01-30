.PHONY: all test

all: ccload

test: ccload
	./ccload -u https://divar.ir -n 100 -c 10

ccload: main.go
	go build -o ccload
