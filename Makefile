.PHONY:
.SILENT:
.DEFAULT_GOAL:= up

include ./api/Makefile
include ./front/Makefile

# ==============================================================================
# Main

init: full-clear env-init docker-build up \
 full-init full-done

up: docker-up
down: docker-down
restart: down up

full-clear: docker-down-clear api-clear front-clear
full-init: api-init front-init
full-done: api-done front-done

full-check: full-lint full-test
full-lint: api-lint front-lint
full-test: api-test front-test
full-upgrade: api-mod-update front-yarn-upgrade

# ==============================================================================
# Docker support

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down --remove-orphans

docker-build:
	docker-compose build $(p)

docker-rebuild:
	docker-compose up -d --no-deps --build $(p)

docker-down-clear:
	docker-compose down -v --remove-orphans

# ==============================================================================
# Common

env-init: env-main-create

env-main-create:
	[ -f .env ] || { cp .env.example .env && echo "Created .env"; }
