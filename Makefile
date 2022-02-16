test:
	@go test ./...

upgrade:
	@echo "Upgrading dependencies..."
	@go get -u
	@go mod tidy
	
run:
	@go run main.go

build:
	@go build -o vwap_calculator main.go

clean:
	@rm -rf vwap_calculator