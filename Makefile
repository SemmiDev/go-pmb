run-tests:
	go test ./... -v
run-app:
	go run cmd/api/main.go