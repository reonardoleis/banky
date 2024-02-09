run-api:
	go run adapter/http/main.go

build-api:
	go build -o ./bin/api adapter/http/main.go

run-manager:
	go run adapter/tcp/main.go

build-manager:
	go build -o ./bin/manager adapter/tcp/main.go
