include .env

# Protobuff Commands
genpb:
	@echo "Protocol Buffer Files Generation (BEGIN)"
	@protoc --go_out=./grpc \
		    --go-grpc_out=./grpc \
		    ./protobuff/*.proto
	@echo "Protocol Buffer Files Generation (END)"
	
# Application Commands
run:
	@go run ./cmd/.
build:
	@go build -o ./bin/uploader ./cmd
test:
	@go test -v -cover ./...
tidy:
	@go mod tidy

# Database migration
mig:
	@migrate create -ext sql -seq -dir repository/migration -seq files_schema
migup:
	@migrate -database $(PSQL_DB_PATH) -path repository/migration up 
migdwn:
	@migrate -database $(PSQL_DB_PATH) -path repository/migration down
	
# SQLc Code Generation
sqlcgen:
	@sqlc generate
	
# Docker Commands
startdb:
	@docker start $(CONTAINER_NAME)
stopdb:
	@docker stop $(CONTAINER_NAME)

.PHONY: genpb sqlcgen