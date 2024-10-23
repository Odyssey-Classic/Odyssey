.PHONY: clean protoc bundle build

SERVER_PATH = ./server
CLIENT_PATH = ./client

clean:
	rm -rf ${SERVER_PATH}/pb/*
	rm -rf ${CLIENT_PATH}/pb/*
	rm -rf ${CLIENT_PATH}/bundle/*

protoc:
	mkdir -p ${SERVER_PATH}/pb
	mkdir -p ${CLIENT_PATH}/pb
	protoc \
		--go_out=:./server/pb \
		--js_out=import_style=commonjs,binary:./client/pb \
		-I./proto game_message.proto

bundle:
	npx browserify ./client/src/index.js -o ./client/bundle/bundle.js

host:
	npx http-server ./client/

build: clean protoc bundle
