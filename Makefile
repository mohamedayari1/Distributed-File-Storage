build:
	go build -o bin/fs

run: build 
	go run ./bin/fs

test:
	go test ./... -v