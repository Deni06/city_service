# CityService Service

This is the CityService service

Generated with

```
micro new gitlab.visionet.co.id\ikhsan\CityService --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.CityService
- Type: srv
- Alias: CityService

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./CityService-srv
```

Build a docker image
```
make docker
```