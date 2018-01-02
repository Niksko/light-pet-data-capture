# Light Pet Data Capture Server

## Requirements

Docker and docker-compose

## Running

`docker-compose up`

## Testing

`go test`

## Compiling protobuf definitions

`docker-compose -f docker-compose-build.yml up`

This will build a container with the pre-requisites for compiling a
`.proto` file into Go based protobuf definitions.
