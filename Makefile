build: fmt
	go build ./cmd/activate

fmt: 
	go fmt ./...

test:
	go test ./...

bin:

