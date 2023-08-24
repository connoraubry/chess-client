build:
	go build -o bin/chess main.go 
run:
	go run main.go
dev:
	go run main.go -level debug
test:
	go test ./...
