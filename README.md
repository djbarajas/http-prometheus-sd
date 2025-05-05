# Prometheus HTTP Service Discovery

This is a sample project to demonstrate building an HTTP service discovery service for prometheus.

## Go Service

The [go_service](go_service) directory is a sample service that exposes metrics on a given port. 

## Http Service Discovery

The [http_service_discovery](http_service_discovery) directory is an example service that gets a list of services to scrape in a hardcoded way. Its returned by calling `GET /service-discovery`.

In a production environment logic would need to be added to find services in a dynamic way. The standard output for service discovery is defined [here](https://prometheus.io/docs/prometheus/latest/http_sd/). And the output format is below:

```
[
  {
    "targets": [ "<host>", ... ],
    "labels": {
      "<labelname>": "<labelvalue>", ...
    }
  },
  ...
]
```

## Compose
The `compose.yaml` is a sample on how to run this in practice. It hard codes 2 metrics services and 1 http service discovery service. In practice, one would do the following:
1. GET /service-discovery
2. GET service-1/metrics, GET service-2/metrics, ...

Where step 1 returns a list of targets to scrape then step 2 actually scrapes for metrics.

## Running the Example
To run, do the following:
```
./build_binaries.sh
podman compose up # or docker compose up 
```

Then you can run the following to see the results:
```
# gets targets with service discovery
$ curl localhost:9000/service-discovery | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    78  100    78    0     0  26748      0 --:--:-- --:--:-- --:--:-- 39000
[
    {
      "targets": [
        "localhost:9090",
        "localhost:9091"
      ],
      "labels": {
        "host": "localhost"
      }
    }
]
# scrape a target retrieved from above
$ curl localhost:9090/metrics
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 1
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
...
```
