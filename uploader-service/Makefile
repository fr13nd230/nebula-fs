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
	@go run ./cmd/main.go
build:
	@go build -o ./bin/uploader ./cmd
test:
	@go test ./...-v -cover
cleanup:
	@go mod tidy

# Docker Compose Commands
cmp-up:
	@docker compose up -d --build
cmp-down:
	@docker compose down -v

# Cache Layer Commands
startcache:
	@docker start $(CACHE_CONTAINER)
stopcache:
	@docker stop $(CACHE_CONTAINER)

.PHONY: cmp-up cmp-down startcache stopcache