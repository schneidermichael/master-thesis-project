# Deployment with Istio

## Install with CLI
```
istioctl install --set profile=minimal -y
```
## Label namespace
```
kubectl label namespace default istio-injection=enabled
```

## Deploy microservices with loadbalancer
```
kubectl apply -f loadbalancer/microservices-istio.yaml
```

## Deploy microservices with ingress
```
kubectl apply -f ingress/microservices-istio.yaml

kubectl apply -f ingress/ingress.yaml
```


Forward frontend locally to port 8080 (only for local tests)
```
kubectl port-forward deployment/frontend 8080:8080
```

Ensure that there are no issues with the configuration

```
istioctl analyze 
```

# Istio Telemetry Addons

## Prometheus
```
kubectl apply -f prometheus.yaml

kubectl -n istio-system port-forward deployment/prometheus 9090:9090
```

## Kiali
```
kubectl apply -f kiali.yaml
```

Access the Kiali dashboard
```
istioctl dashboard kiali
```
## Jäger
```
kubectl apply -f jaeger.yaml
```
## Grafana
```
kubectl apply -f grafana.yaml

kubectl -n istio-system port-forward deployment/grafana 3000:3000
```

