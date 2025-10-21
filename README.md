# textproto-gen

`textproto-gen` provide `protoc-gen-textproto` - a protoc plugin for generate textproto files.

## Install

```shell
go install github.com/bookweb/textproto-gen/cmd/protoc-gen-textproto@latest
```

## Usage

```shell
protoc --textproto_out=. service.proto
```
