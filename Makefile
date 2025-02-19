include ./configs/.env

run:
	go run cmd/supportchat/main.go

run-linter:
	echo "Starting linters"
	golangci-lint run ./...

# Docker commands
dev:
	echo "Starting docker dev environment"
	docker-compose --env-file ./configs/.env -f deployments/docker-compose.dev.yml up --build
logs:
	echo "Starting docker dev environment"
	docker-compose --env-file ./configs/.env -f deployments/docker-compose.dev.yml  logs -f
devd:
	echo "Starting docker dev environment"
	docker-compose --env-file ./configs/.env -f deployments/docker-compose.dev.yml up --build -d
dev-down:
	echo "Stopping docker dev environment"
	docker-compose --env-file ./configs/.env -f deployments/docker-compose.dev.yml down -v

migrate_create_docker:
	docker-compose --env-file ./configs/.env -f ./deployments/docker-compose.dev.yml exec supportchat_migrate migrate create -ext sql -dir /migrations $(MNAME)
migrate_up_docker:
	docker-compose --env-file ./configs/.env -f ./deployments/docker-compose.dev.yml exec supportchat_migrate migrate -path /migrations -database 'mysql://$(MYSQL_USERNAME):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)' up
migrate_down_docker:
	docker-compose --env-file ./configs/.env -f ./deployments/docker-compose.dev.yml exec supportchat_migrate migrate -path /migrations -database 'mysql://$(MYSQL_USERNAME):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DATABASE)' down