all:
	make build

build:
	go build *.go

build-compose:
	export CONTAINER_NAME=go_rest_api \
	&& export IMAGE_TAG=go_rest_api:latest \
	&& export DB_PASS=root \
	&& export DB_NAME=go_rest_api \
	&& envsubst < docker-compose.tmpl.yml > docker-compose.yml \
	&& docker build -t $$IMAGE_TAG .

dev:
	fiber dev -p "swag init" -D docs
