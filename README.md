# Edge Platform - T1

## Overview

![Edge Platform Architecture](https://docs.google.com/drawings/d/e/2PACX-1vROWOYEMLy9lta-RgUw9Ex0Gd3Um2wHl7wYJF_uLJ5_xpdm_YZnQOhU94z6pmJPS4nf9qClhej8ur46/pub?w=960&h=720)

## Building

1. This entire repository is built with Golang! Make sure you have installed [Golang](https://golang.org/doc/install)

1. We utilize Protobuf [Protocol Buffers](https://developers.google.com/protocol-buffers/) for platform communication over GRPC. There are a few dependencies that must be done to generate our protocol.
    1. Install the protocol generating tool: 
        ```bash
        go get -u github.com/golang/protobuf/protoc-gen-go
        ```
    1. Install the GRPC -> REST -> Swagger Plugin
        ```bash
        go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
        
        go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
        ```
    1. Ensure you can run `protoc-gen-go` and `protoc-gen-grpc-swagger` from your command line. If they are not found check your `$PATH`.
1. In order to ease the commands for building, we utilize `make` therefore ensure make is installed for [Linux](https://www.gnu.org/software/make/)/[Windows](http://gnuwin32.sourceforge.net/packages/make.htm)
1. Data is stored in a local redis key-value store. Ensure it is [installed](https://redis.io/topics/quickstart) and running locally.

1. **Finally!** Go ahead and run `make api` to generate the Protobuf. Then `make` to build `client.bin` and `server.bin`.


## Running

1. After building, run `server.bin` to run the server side.
1. Run the client side tests with `client.bin`. You can also use the REST API endpoints to test the API. To do this use [Swagger Editor](https://editor.swagger.io/) with the generated `api.swagger.json`