.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./posts/posts
	
build:
	cd posts && GOOS=linux GOARCH=amd64 go build -o ./bin/posts ./