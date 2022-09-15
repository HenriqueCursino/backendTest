coverage:
	@echo "Running project tests..."
	@go test -v -cover ./...
	@go test ./... -v -coverprofile=cover.out
	@go tool cover -html=cover.out -o cover.html

cover:
	@go test -v ./...
	@echo "Running project tests..."
	@go test ./... -v -covermode=count -coverprofile=count.out
	@go tool cover -html=count.out -o count.html

    