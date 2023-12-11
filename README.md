# PaymentApi Service

This is the PaymentApi service

Generated with

```
micro new --namespace=go.micro --type=api payment-api
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.api.payment-api
- Type: api
- Alias: payment-api

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./payment-api-api
```

Build a docker image
```
make docker
```