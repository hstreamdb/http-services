# http-services

HStreamDB's http-related services. Including.

- [http server](#http-server)
- [admin client](#admin-client)

## Usage

### HTTP SERVER

The http server of HStreamDB provides the relevant api for accessing the HStreamDB service through http requests.

#### Start

```go
git clone git@github.com:hstreamdb/http-services.git
cd http-services
make server
make runServer
```

#### Generate Swagger

- `make swag`
- Then you can start the server and go to http://localhost:8080/v1/swagger/index.html to see your Swagger UI.

### ADMIN CLIENT

`adminCtl` is the command line tool for managing `HStream-server` cluster.

Currently, user can use `adminCtl` to：

- create/list all streams
- create/list all subscriptions
- get server status
- get statistics

#### USAGE

- `make adminCtl` to build from source, the binary will be put in `{project_dir}/bin/adminCtl`
- `./bin/adminCtl -h` to get more info.

