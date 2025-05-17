# Odyssey
Modern remake of Odyssey Classic

# Development

## Local Dev
Use build and bundle source:  
`make bundle`

Use http-server to host client locally:  
`make host`

## Protobufs
`npm install -g protoc-gen-ts`  
`go install github.com/golang/protobuf/protoc-gen-go@latest`

https://github.com/protocolbuffers/protobuf  
https://github.com/protocolbuffers/protobuf-javascript  
https://github.com/protocolbuffers/protobuf-go

## Running the Registry Service Locally
To start the registry service and a local MongoDB instance for development:

```
make registry
```

This will automatically start a MongoDB container (using Docker) and then run the registry service. To stop the MongoDB container when you're done:

```
make mongodb-stop
```

## Requirements

- **Docker**: Required for running MongoDB locally via the Makefile. If you do not have Docker installed, follow the instructions for your platform:
  - [Docker installation for Linux](https://docs.docker.com/engine/install/)
  - [Docker installation for Windows](https://docs.docker.com/desktop/install/windows-install/)
  - [Docker installation for Mac](https://docs.docker.com/desktop/install/mac-install/)
