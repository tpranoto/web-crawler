gochallenge:
	@go build -o gochallenge main.go
	@./gochallenge
	
test:
	@go test -cover -race -short -count=1 ./... | grep -v *mock*