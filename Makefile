run-api:
	go run cmd/http/main.go

build-api:
	go build -o ./bin/api cmd/http/main.go

run-manager:
	go run cmd/tcp/main.go

build-manager:
	go build -o ./bin/manager cmd/tcp/main.go
