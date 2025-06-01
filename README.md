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

## WSL2: Using a Custom Domain for Discord OAuth

If you are running the registry inside WSL2 and want to use Discord OAuth (which does not allow `localhost` as a redirect URI), follow these steps:

1. **Find your WSL2 instance's IP address:**
   In your WSL2 terminal, run:
   ```bash
   hostname -I | awk '{print $1}'
   ```
   Copy the resulting IP address (e.g., `172.24.128.1`).

2. **Edit your Windows hosts file:**
   - Open Notepad as Administrator.
   - Open the file: `C:\Windows\System32\drivers\etc\hosts`
   - Add a line at the end:
     ```
     172.24.128.1 odyssey.local
     ```
     (Replace `172.24.128.1` with your actual WSL2 IP.)

3. **Update your `.env` and Discord app settings:**
   - In `.env`, set:
     ```
     ODY_REDIRECT_URL=http://odyssey.local:8080/identity/oauth/callback
     ```
   - In the Discord Developer Portal, add the same URL to your app's list of allowed redirect URIs.

4. **Restart your registry service** after making these changes.

Now you can use `http://odyssey.local:8080/identity/login` in your browser to start the OAuth flow.
