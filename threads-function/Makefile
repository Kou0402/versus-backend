.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./threads/threads
	
test:
	cd threads && go test -v --cover ./...

build:
	cd threads && GOOS=linux GOARCH=amd64 go build -o ./bin/threads ./cmd