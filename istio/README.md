# Deployment with Istio

```
istioctl install --set profile=default -y
```
```
kubectl create namespace microservices
```

```
kubectl label namespace microservices istio-injection=enabled
```

```
kubectl apply -f microservices-istio.yaml
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

```
kubectl create namespace telemetry
```

```
kubectl create namespace gui
```

## Prometheus
```
kubectl apply -f prometheus.yaml -n telemetry
```

## Kiali
```
kubectl apply -f kiali.yaml -n gui
```

Access the Kiali dashboard
```
istioctl dashboard kiali
```
## Jäger
```
kubectl apply -f jaeger.yaml -n telemetry
```
## Grafana
```
kubectl apply -f grafana.yaml -n telemetry
```

