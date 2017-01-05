build-client: fmt
	go build main/tick-ext -o out/tick-ext

build-server: fmt
	go build main/tick-extd -o out/tick-extd

fmt: 
	go fmt ./...

test:
	go test ./...

bin:

