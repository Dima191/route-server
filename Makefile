generate_code:
	protoc -I ./api --go_out=./pkg --go_opt=paths=source_relative --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative ./api/route_server.proto

migrate_up:
	migrate -verbose -path ./migrations -database "driver://user:your_password@localhost:port/db_name" up

migrate_down:
	migrate -verbose -path ./migrations -database "driver://user:your_password@localhost:port/db_name" down

run:
	go run ./cmd/app/main.go --debug=true

.PHONY: migrate_up migrate_down run

.DEFAULT_GOAL=run