build-all:
	cd cart && GOOS=linux GOARCH=amd64 make build && \
 	cd .. && \
	cd loms && GOOS=linux GOARCH=amd64 make build && \
	cd .. && \
	cd notifier && GOOS=linux GOARCH=amd64 make build

run-all: build-all
	docker-compose up --force-recreate -d postgres-loms kafka-ui kafka0 kafka-init-topics && \
	cd loms && make container-migration-up && \
	cd .. && \
	docker-compose up --build -d cart notifier loms

logs-notifier:
	echo 1 2 3 | xargs -n 1 -I {} docker logs homework-notifier-{}

logs-loms:
	docker logs loms
