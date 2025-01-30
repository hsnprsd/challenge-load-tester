.PHONY: all test

all: ccload

test: ccload
	./ccload -u https://google.com -n 100 -c 10 --method GET --expect 200

ccload: main.go
	go build -o ccload
