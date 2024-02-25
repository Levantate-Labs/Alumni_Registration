build:
	@go build -o bin/fworld

run: build
	@./bin/fworld