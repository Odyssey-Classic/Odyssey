.PHONY: clean protoc

SERVER_PATH = ./server
CLIENT_PATH = ./client

clean:
	rm -rf ${SERVER_PATH}/pb/*
	rm -rf ${CLIENT_PATH}/build/pb/*

protoc:
	mkdir -p ${SERVER_PATH}/pb
	mkdir -p ${CLIENT_PATH}/build/pb
	protoc \
		--go_out=:./server/pb \
		--js_out=import_style=commonjs,binary:./client/build/pb \
		-I./proto game_message.proto
