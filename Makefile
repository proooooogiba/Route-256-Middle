build-all:
	cd cart && GOOS=linux GOARCH=amd64 make build && \
 	cd .. && \
	cd loms && GOOS=linux GOARCH=amd64 make build


run-all: build-all
	docker-compose up -d postgres-loms && \
	cd loms && make migration-up && \
	cd .. && \
	docker-compose up --build -d cart loms
