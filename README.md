# Disney Infinity 1.0  (Wii U) replacement NEX server
This server only includes the NEX servers that were hosted by Nintendo (authentication and secure). **This does not include the Disney APIs.**

## Compiling

### Setup
Install [Go](https://go.dev/doc/install) and [git](https://git-scm.com/downloads), then clone and enter the repository

```bash
$ git clone https://github.com/BDMCGaming/disney-infinity-3-nex
$ cd disney-infinity-1-nex
```

### Compiling using `go`
To compile using Go, `go get` the required modules and then `go build` to your desired location. You may also want to tidy the go modules, though this is optional

```bash
$ go get 
$ go mod tidy
$ go build -o build/disney-infinity-1-nex
```

The server is now built to `build/disney-infinity-1-nex`

When compiling with only Go, the authentication servers build string is not automatically set. This should not cause any issues with gameplay, but it means that the server build will not be visible in any packet dumps or logs a title may produce

To compile the servers with the authentication server build string, add `-ldflags "-X 'main.serverBuildString=BUILD_STRING_HERE'"` to the build command, or use `make` to compile the server

### Compiling using `make`
Compiling using `make` will read the local `.git` directory to create a dynamic authentication server build string, based on your repositories remote origin and current commit. It will also use the current folders name as the executables name

Install `make` onto your system (this varies by OS), and run `make` while inside the repository

```bash
$ make
```

The server is now built to `build/disney-infinity-1-nex` with the authentication server build string already set

### Developing core libraries
This server depends on [nex-protocols-common-go](https://github.com/PretendoNetwork/nex-protocols-common-go),
[nex-protocols-go](https://github.com/PretendoNetwork/nex-protocols-go), and
[nex-go](https://github.com/PretendoNetwork/nex-go) to implement the bulk of its functionality. Sometimes features or
bugfixes depend on unreleased development versions of these libraries; or changes have to be made at the library level
to work on this server.

To use these unreleased libraries, add a `go.work` file with something like the following content:
```go
go 1.25

use (
	.
)

replace (
	github.com/PretendoNetwork/nex-go/v2 => ../nex-go
	github.com/PretendoNetwork/nex-protocols-go/v2 => ../nex-protocols-go
	github.com/PretendoNetwork/nex-protocols-common-go/v2 => ../nex-protocols-common-go
)
```

The build will then build these libraries out of the folders specified instead of pulling the released versions.

## Configuration
All configuration options are handled via environment variables

`.env` files are supported

| Name                                      | Description                                                                                     | Required                                      |
|-------------------------------------------|-------------------------------------------------------------------------------------------------|-----------------------------------------------|
| `PN_INFINITY1_POSTGRES_URI`               | Fully qualified URI to your Postgres server (Example `postgres://username:password@localhost/disneyinfinity1?sslmode=disable`)  | Yes                                           |
| `PN_INFINITY1_AUTHENTICATION_SERVER_PORT` | Port for the authentication server                                                              | Yes                                           |
| `PN_INFINITY1_SECURE_SERVER_HOST`         | Host name for the secure server (should point to the same address as the authentication server) | Yes                                           |
| `PN_INFINITY1_SECURE_SERVER_PORT`         | Port for the secure server                                                                      | Yes                                           |
| `PN_INFINITY1_ACCOUNT_GRPC_HOST`          | Host name for your account server gRPC service                                                  | Yes                                           |
| `PN_INFINITY1_ACCOUNT_GRPC_PORT`          | Port for your account server gRPC service                                                       | Yes                                           |
| `PN_INFINITY1_ACCOUNT_GRPC_API_KEY`       | API key for your account server gRPC service                                                    | No (Assumed to be an open gRPC API)           |
| `PN_INFINITY1_FRIENDS_GRPC_HOST`          | Host name for Friends gRPC service (used for friends matchmaking)                               | Yes                                           |
| `PN_INFINITY1_FRIENDS_GRPC_PORT`          | Port for Friends gRPC service                                                                   | Yes                                           |
| `PN_INFINITY1_FRIENDS_GRPC_API_KEY`       | API key for Friends gRPC service                                                                | No (Assumed to be an open gRPC API)|
| `PN_INFINITY1_S3_ENDPOINT`				| URL for DataStore S3 storage (Example `10.0.0.241:9000`)										  | Yes											  |
| `PN_INFINITY1_S3_ACCESS_KEY`				| Access key for S3																				  | Yes											  |
| `PN_INFINITY1_S3_ACCESS_SECRET`			| Secret key for S3																				  | Yes											  |
| `PN_INFINITY1_S3_BUCKET`					| Bucket for DataStore storage																	  | Yes											  |
| `PN_INFINITY1_S3_KEY_BASE`				| Base for Datastore keys																		  | Yes											  |
| `PN_INFINITY1_S3_INSECURE`				| Use HTTP for S3 endpoints	(Boolean)															  | No (Off by default)							  |
| `PN_INFINITY1_HEALTH_CHECK_PORT`          | Port for the UDP health check                                                                   | No