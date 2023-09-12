# Pulsar

## Installing

go install github.com/cosmos/cosmos-proto/cmd/protoc-gen-go-pulsar

## Running

cd path/to/proto/files

protoc --go-pulsar\_out=. --go-pulsar\_opt=paths=source\_relative
\--go-pulsar\_opt=features=marshal+unmarshal+size -I . NAME\_OF\_FILE.proto

## Acknowledgements

Code for the generator structure/features and the functions marshal, unmarshal,
and size implemented by
[planetscale/vtprotobuf](https://github.com/planetscale/vtprotobuf) was used in
our `ProtoMethods` implementation.

Code used to produce default code stubs found in
[protobuf](https://pkg.go.dev/google.golang.org/protobuf) was copied into
[features/protoc](./features/protoc).
