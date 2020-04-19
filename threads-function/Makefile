.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./threads/threads
	
build:
	GOOS=linux GOARCH=amd64 go build -o threads/bin/threads ./threads