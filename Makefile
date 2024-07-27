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

stop-all:
	docker-compose down

restart-all:
	docker-compose restart

logs-notifier:
	echo 1 2 3 | xargs -n 1 -I {} docker logs homework-notifier-{}

logs-loms:
	docker logs loms
