build:
	@echo "Compiling..."
	@go build -o bin/bin_file main.go

run: build 
	@./bin/bin_file