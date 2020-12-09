all:
	make build

build:
	go build *.go

dev:
	fiber dev -p "swag init" -D docs
