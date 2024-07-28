build-all:
	cd cart && GOOS=linux GOARCH=amd64 make build && \
 	cd .. && \
	cd loms && GOOS=linux GOARCH=amd64 make build && \
	cd .. && \
	cd notifier && GOOS=linux GOARCH=amd64 make build

start-all: build-all
	docker-compose up -d && \
	sleep 5 && \
    docker-compose up -d

recreate-all: build-all
	docker-compose up -d --build --force-recreate --no-deps && \
	sleep 5 && \
	docker-compose up -d

stop-all:
	docker-compose down

restart-all:
	docker-compose restart

logs-notifier:
	echo 1 2 3 | xargs -n 1 -I {} docker logs homework-notifier-{}

logs-loms:
	docker logs loms



# SHARD-DB MIGRATION START

-include .env

CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin

MIGRATION_FOLDER=loms/migrations

.PHONY: install-goose
install-goose:
	GOBIN=${BINDIR} go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: goose-migrate-all
goose-migrate-all: goose-migrate-1 goose-migrate-2

.PHONY: goose-migrate-down-all
goose-migrate-down-all: goose-migrate-1-down goose-migrate-2-down

.PHONY: goose-migrate-1
goose-migrate-1:
	${BINDIR}/goose -dir ${MIGRATION_FOLDER} postgres "\
		host=localhost\
		user=${DOCKER_POSTGRES_USER_1}\
		dbname=${DOCKER_POSTGRES_DB_1}\
		password=${DOCKER_POSTGRES_PASSWORD_1}\
		port=${DOCKER_POSTGRES_HOST_PORT_1}\
		sslmode=disable" up

.PHONY: goose-migrate-2
goose-migrate-2:
	${BINDIR}/goose -dir ${MIGRATION_FOLDER} postgres "\
		host=localhost\
		user=${DOCKER_POSTGRES_USER_2}\
		dbname=${DOCKER_POSTGRES_DB_2}\
		password=${DOCKER_POSTGRES_PASSWORD_2}\
		port=${DOCKER_POSTGRES_HOST_PORT_2}\
		sslmode=disable" up

.PHONY: goose-migrate-1-down
goose-migrate-1-down:
	${BINDIR}/goose -dir ${MIGRATION_FOLDER} postgres "\
		host=localhost\
		user=${DOCKER_POSTGRES_USER_1}\
		dbname=${DOCKER_POSTGRES_DB_1}\
		password=${DOCKER_POSTGRES_PASSWORD_1}\
		port=${DOCKER_POSTGRES_HOST_PORT_1}\
		sslmode=disable" down

.PHONY: goose-migrate-2-down
goose-migrate-2-down:
	${BINDIR}/goose -dir ${MIGRATION_FOLDER} postgres "\
		host=localhost\
		user=${DOCKER_POSTGRES_USER_2}\
		dbname=${DOCKER_POSTGRES_DB_2}\
		password=${DOCKER_POSTGRES_PASSWORD_2}\
		port=${DOCKER_POSTGRES_HOST_PORT_2}\
		sslmode=disable" down

# SHARD-DB MIGRATION END
