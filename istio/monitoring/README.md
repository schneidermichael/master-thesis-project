# Istio Monitoring 

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
