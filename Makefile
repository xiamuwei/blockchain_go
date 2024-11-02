build:
	go build -o ./blockchain_go

run: build
	./blockchain_go
	
test:
	go test -v ./...