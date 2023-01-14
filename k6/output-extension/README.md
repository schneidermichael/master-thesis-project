#  k6 Output Extensions for SMI

Created an output extension for [Service Mesh Performance](https://smp-spec.io/)

## Install xk6
```
install xk6
```

## Compile 
```
xk6 build --with xk6-output-logger=.
```

## Use extension
```
./k6 run script.js -o smp --quiet --no-summary --vus 5 --duration 10s
```

### Example of YAML Output
```
start_time: 2023-01-14T17:31:15.855756+01:00
end_time: 2023-01-14T17:31:27.627444+01:00
exp_uuid: d5ed0b72-5e85-41ec-8808-fa162a869b81
endpoint_url: http://localhost:8080
load_generator: k6
env:
    kubernetes: v1.25.2
client:
    connections: 1
    rps: 4.2474620872901765
    latencies_ms:
        min: 40.429
        average: 78.63252
        p50: 64.921
        p90: 112.574
        p99: 189.559
        max: 189.559
metrics: {}
```