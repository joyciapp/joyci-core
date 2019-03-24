# joyci-core
Core of JoyCI

# Tests

To run the test suite runs:
```
$ go test ./... -v
```

# GRPC

## Protocol Buffers
To compile protobuffers, in the project's root folder run the command bellow:
```
$ protoc -I grpc grpc/proto/core.proto --go_out=plugins=grpc:grpc
```

## Server
*TODO*