DOCKER=docker compose
ENV=LOCAL

.PHONY: help
help:   ## This help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: up
up:   ## local実行
	$(DOCKER) up --build

.PHONY: down
down:  ## docker compose down
	$(DOCKER) down

.PHONY: restart
restart:  ## docker compose restart
	docker restart funcy_portfolio_backend-api-1

.PHONY: demo
demo:  ## demo
	$(DOCKER) up --build
	docker exec funcy_portfolio_backend-api-1 ./db/migrate -path db/migration/sql -database "mysql://root:admin@tcp(mysql:3306)/funcy?multiStatements=true" up
	docker restart funcy_portfolio_backend-api-1

.PHONY: create-%
create-%:  ## create sql file
	migrate create -ext sql -dir db/migration/sql ${@:create-%=%}
.PHONY: migrate
migrate:  ## migrate up
	migrate -path db/migration/sql -database "mysql://root:admin@tcp(127.0.0.1:3306)/funcy?multiStatements=true" up
	docker restart funcy_portfolio_backend-api-1

.PHONY: migrate-down
migrate-down:  ## migrate down
	migrate -path db/migration/sql -database "mysql://root:admin@tcp(127.0.0.1:3306)/funcy?multiStatements=true" down
	docker restart funcy_portfolio_backend-api-1

.PHONY: migrate-force
migrate-force:  ## migrate down
	migrate -path db/migration/sql -database "mysql://root:admin@tcp(127.0.0.1:3306)/funcy?multiStatements=true" force 20221026122655
	docker restart funcy_portfolio_backend-api-1

.PHONY: migrate-demo
migrate-demo:  ## migrate up dmeo
	docker exec funcy_portfolio_backend-api-1 ./db/migrate -path db/migration/sql -database "mysql://root:admin@tcp(mysql:3306)/funcy?multiStatements=true" up
	docker restart funcy_portfolio_backend-api-1