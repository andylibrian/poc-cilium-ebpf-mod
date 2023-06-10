EXECUTABLE=bin

gen: 
	go generate ./...

build: gen
	go build -o ./$(EXECUTABLE)/ ./cmd/...