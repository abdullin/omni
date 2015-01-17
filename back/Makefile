# force rebuild on these targets
.PHONY: build run

# default build
default: build

build:
	go build -v -o ./bin/host ./host/main.go

run: build
	./bin/host


reset:
	go run ./host/reset.go
