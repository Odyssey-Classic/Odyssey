# Odyssey
Modern remake of Odyssey Classic

# Development

## Local Dev
Use browserify to bundle client source:  
`npx browserify ./client/src/index.js -o ./client/bundle.js`

Use http-server to host client locally:  
`npx http-server ./client`

## Protobufs
`npm install -g protoc-gen-js`  
`go install github.com/golang/protobuf/protoc-gen-go@latest`

https://github.com/protocolbuffers/protobuf  
https://github.com/protocolbuffers/protobuf-javascript  
https://github.com/protocolbuffers/protobuf-go  
