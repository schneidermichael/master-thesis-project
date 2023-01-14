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

## Forward frontend locally to port 8080 (only for local tests)
```
kubectl port-forward deployment/frontend 8080:8080
```

