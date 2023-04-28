# Ogen integration/staging tree

[![Go Report](https://goreportcard.com/badge/github.com/kindynosmx/ogen)](https://goreportcard.com/report/github.com/kindynosmx/ogen)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/kindynosmx/ogen?tab=doc)](https://pkg.go.dev/github.com/kindynosmx/ogen?tab=doc)

## **DEPRECATED: use only as a reference.**

Ogen is the main implementation of the Olympus protocol.

## Building

```bash
./scripts/build.sh
```

## Run with Docker

### Full node with shared storage

> This configuration is the best to use as a validator

#### Pull the image

```
docker pull ghcr.io/kindynosmx/ogen-node:latest
```

#### Run the container

> Make sure you replace the LOCAL_HOST_FOLDER on the string to the host folder you want to store your files

```
docker run -p 80:8080 -p 81:8081 -d -v LOCAL_HOST_FOLDER:/root/.config/ogen ghcr.io/kindynosmx/ogen-node:latest
```

Now you will have a full-node instance running on the background with a dashboard exposed on port 80, and the REST API exposed on 81, and the full-node files are stored on your host on LOCAL_HOST_FOLDER

### Full node without shared storage

> This configuration is the best to run a simple full node without having a backup of the keystore.

#### Pull the image

```
docker pull ghcr.io/kindynosmx/ogen-node:latest
```

#### Run the container

```
docker run -p 80:8080 -p 81:8081 -d ghcr.io/kindynosmx/ogen-full-node:latest
```

Now you will have a full-node instance running on the background with a dashboard exposed on port 80, and the REST API exposed on 81
