# http-services

HStreamDB's http-related services. Including:

- [http server](#http-server)

## Installation

- You can download the binary release suitable for your system

- Or you can build from source code

  ```shell
  git clone git@github.com:hstreamdb/http-services.git
  cd http-services
  make all
  ```

  then you can find the binary in `{project_dir}/bin`

## Quickstart with Docker

- You can use the [HStream image](https://hub.docker.com/r/hstreamdb/hstream)
```sh
docker run -td --network host \
--name hstream-http-server hstreamdb/hstream:v0.9.0 \
hstream-http-server -address "localhost:8080" -log-level "info" -services-url "localhost:6580"
```

- Or use the standalone image `docker pull ghcr.io/hstreamdb/http-services:latest`

## HTTP SERVER

The http server of HStreamDB provides the relevant api for accessing the HStreamDB service through http requests.

### Usage

- First, you need to confirm `HStreamDB` cluster is up. 
  - To set up a `HStreamDB` cluster, you can see [Manual Deployment with Docker](https://hstream.io/docs/en/latest/deployment/deploy-docker.html) and deployment related part.

- Start `http-server`：`./bin/http-server -services-url <hstreamdb-server address>` 
  - Use `http-server -h` to see more details.

### Generate Swagger

- `make swag`
- Then you can start the server and go to http://localhost:8080/v1/swagger/index.html to see your Swagger UI.
