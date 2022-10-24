.PHONY:
.SILENT:
#.DEFAULT_GOAL:= up

# ==============================================================================
# Api commands

api-init: api-install api-wait-db migrate-up
api-install: api-mod-download api-mod-vendor
api-check: api-lint api-test

api-mod-tidy:
	docker-compose exec api-go go mod tidy

api-mod-vendor:
	docker-compose exec api-go go mod vendor

api-mod-update:
	docker-compose exec api-go go get -u ./...
	docker-compose exec api-go go mod tidy
	docker-compose exec api-go go mod vendor

api-mod-download:
	docker-compose exec api-go go mod download

api-wait-db:
	docker-compose exec api-go wait-for-it api-postgres:5432 -t 30

api-lint:
	docker-compose exec api-go golangci-lint run

api-test:
	docker-compose exec api-go go test -count=1 -p=8 -parallel=8 -race ./...

api-test-coverage:
	docker-compose exec api-go go test --short -coverprofile=./tmp/test/cover.out -v ./...
	docker-compose exec api-go go tool cover -func=./tmp/test/cover.out

api-gen-full: api-gen-oapi api-gen-wire api-gen-gojet

api-gen-wire:
	docker-compose exec api-go wire ./internal/infrastructure/wire/

api-gen-oapi:
	docker-compose exec api-go go generate ./specs/openapi/

api-gen-gojet:
	docker-compose exec api-go gojet -path=./internal/models -ignore-tables=goose_db_version

# ==============================================================================
# Console commands

cli-build:
	docker-compose exec api-go go build -o bin/cli cmd/cli/main.go

cli-run:
	docker-compose exec api-go go run cmd/cli/main.go $(p)

# ==============================================================================
# Migrate postgresql

migrate-create:
	docker-compose exec api-go migrate create $(p) $(or $(t), sql)

migrate-up:
	docker-compose exec api-go migrate up

migrate-down:
	docker-compose exec api-go migrate down

migrate-redo:
	docker-compose exec api-go migrate redo

migrate-reset:
	docker-compose exec api-go migrate reset

migrate-version:
	docker-compose exec api-go migrate version

migrate-fix:
	docker-compose exec api-go migrate fix
