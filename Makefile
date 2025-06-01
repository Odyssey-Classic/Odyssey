.PHONY: clean protoc bundle build registry mongodb-run

SERVER_PATH = ./server
CLIENT_PATH = ./client

clean:
	rm -rf ${SERVER_PATH}/pb/*
	rm -rf ${CLIENT_PATH}/pb/*
	rm -rf ${CLIENT_PATH}/dist/*
	cd ./client; npm run clean

protoc:
	mkdir -p ${SERVER_PATH}/pb
	mkdir -p ${CLIENT_PATH}/pb
	protoc \
		--go_out=:./server/pb \
		--ts_out=no_grpc:./client/pb \
		-I./proto game_message.proto

bundle:
	cd ./client; \
	npx webpack


host:
	npx http-server ./client/ -c-1

build: clean protoc bundle

registry: mongodb-run
	env $(shell grep -v '^#' ./registry/.env | xargs) go run ./registry/cmd/

mongodb-run:
	@if ! docker ps --format '{{.Names}}' | grep -q '^odyssey-mongo$$'; then \
		if docker ps -a --format '{{.Names}}' | grep -q '^odyssey-mongo$$'; then \
			docker start odyssey-mongo; \
		else \
			docker run --rm -d --name odyssey-mongo -p 27017:27017 mongo:7; \
		fi \
	else \
		echo "MongoDB container 'odyssey-mongo' is already running."; \
	fi

mongodb-stop:
	docker stop odyssey-mongo
