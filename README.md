## Product Service
The Product Service is written in golang and is in the `golangrpc/` directory. Please note that `proto/` directory is shared by both services.

### Dependency
```
$ go get google.golang.org/grpc
$ go get github.com/rs/xid
```

### Run
Under `golangrpc/` run `$ go run main.go`

## Category Service
This Category Service is written in node.js and is in the `nodeserver/` directory.

### Dependency
Use [npm](https://www.npmjs.com) or [yarn](https://yarnpkg.com)

```
$ npm install
```
or
```
$ yarn install
```

### Run
Under `nodeserver/` run `$ node index.js`

## Protocol Buffers
gRPC serivce code `proto/rpc.pb.go` is generated from `proto/rpc.proto`. If you need to change `rpc.proto`, you will need to regenerate the gRPC service code for `golang`.

### [Install Protocol Buffers v3](https://grpc.io/docs/quickstart/go.html)

Install the protoc compiler that is used to generate gRPC service code. The simplest way to do this is to download pre-compiled binaries for your platform(`protoc-<version>-<platform>.zip`) from here: https://github.com/google/protobuf/releases

+ Unzip this file.
+ Update the environment variable **PATH** to include the path to the protoc binary file.

Next, install the protoc plugin for Go

```
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

The compiler plugin, `protoc-gen-go`, will be installed in `$GOBIN`, defaulting to `$GOPATH/bin`. It must be in your `$PATH` for the protocol compiler, `protoc`, to find it.

```
$ export PATH=$PATH:$GOPATH/bin
```
