.PHONY: clean protoc bundle build

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
