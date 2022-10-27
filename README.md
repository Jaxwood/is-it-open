# Is it open

A small service to check if a port is reachable from within a Kubernetes namespace.

The service will return `200 OK` if the endpoint is reachable.

Otherwise it will return `500 Internal Service Error`.

## Build

Build using docker

```sh
docker build -t is-it-open:0.0.1 --platform=linux/amd64 .
```

## Usage

Port forward the `is-it-open` service and use curl to test an endpoint

```sh
curl -X POST -v -d '{"url": "http://internal-service-name.namespace.svc.cluster.local"}' -H "Content-Type: application/json" localhost:1323/v1/open
```

