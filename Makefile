run-tests:
	go test ./... -v
run-mailer-server:
	go run ./mailer/asynq_email_server.go
run-app-server:
	go run main.go