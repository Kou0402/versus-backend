.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./posts/posts

test:
	cd threads && go test -v --cover ./...
	
build:
	cd posts && GOOS=linux GOARCH=amd64 go build -o ./bin/posts ./cmd