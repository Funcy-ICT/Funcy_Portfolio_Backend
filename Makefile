DOCKER=docker-compose
ENV=LOCAL

.PHONY: help
help:   ## This help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run
run:   ## local実行
	$(DOCKER) up --build

.PHONY: down
down:  ## docker compose down
	$(DOCKER) down

.PHONY: create-%
create-%:  ## create sql file
	migrate create -ext sql -dir db/migration/sql ${@:create-%=%}
.PHONY: migrate
migrate:  ## migrate up
	migrate -path db/migration/sql -database "mysql://root:admin@tcp(127.0.0.1:3306)/funcy?multiStatements=true" up

.PHONY: migrate-down
migrate-down:  ## migrate down
	migrate -path db/migration/sql -database "mysql://root:admin@tcp(127.0.0.1:3306)/funcy?multiStatements=true" down