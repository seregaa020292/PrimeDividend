.PHONY:
.SILENT:
#.DEFAULT_GOAL:= up

# ==============================================================================
# Front commands

front-init: front-yarn-install
front-check: front-lint front-test

front-yarn-install:
	docker-compose exec front-node yarn install

front-yarn-add:
	docker-compose exec front-node yarn add $(p)

front-yarn-upgrade:
	docker-compose exec front-node yarn upgrade --latest

front-lint:
	docker-compose exec front-node yarn lint

front-test:
	docker-compose exec front-node yarn test

front-test-coverage:
	docker-compose exec front-node yarn test:cov
