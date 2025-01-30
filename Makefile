.PHONY: all test

all: ccload

test: ccload
	./ccload -u https://divar.ir -n 100 -c 10 --method POST --expect 200

ccload: main.go
	go build -o ccload
