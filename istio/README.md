# Deployment with Istio

```
istioctl install --set profile=default -y
```

```
kubectl create namespace microservices | kubectl label namespace microservices istio-injection=enabled | kubectl apply -f microservices-istio.yaml
```

Forward frontend locally to port 8080
```
kubectl -n microservices port-forward deployment/frontend 8080:8080
```

Ensure that there are no issues with the configuration

```
istioctl analyze -n microservices
```

# Istio Telemetry Addons

## Prometheus
```
kubectl apply -f prometheus.yaml
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
```

