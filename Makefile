.DEFAULT_GOAL := build

.PHONY: fmt vet build
fmt:
	go fmt .

vet: fmt
	go vet .

build: vet clean
	go build -o ./bin/ . 
	@echo "Build complete."

clean:
	go clean ./bin

test:
	go test .