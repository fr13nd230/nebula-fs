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
	@go test ./...-v -cover
cleanup:
	@go mod tidy

.PHONY: genpb