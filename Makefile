.PHONY: all

all: ccload

ccload: main.go
	go build -o ccload
