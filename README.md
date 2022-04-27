# http-services

HStreamDB's http-related services. Including.

- [http server](#http-server)
- [admin client](#admin-client)

## Installation

- You can download the binary release suitable for your system

- Or you can build from source code

  ```shell
  git clone git@github.com:hstreamdb/http-services.git
  cd http-services
  make all
  ```

  then you can find the binary in `{project_dir}/bin`

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

## ADMIN CLIENT

`adminCtl` is the command line tool for managing `HStream-server` cluster.

Currently, user can use `adminCtl` to：

- create/list all streams
- create/list all subscriptions
- get server status
- get cluster metrics statistics

### USAGE

- `./bin/adminCtl -h` to get more info.
