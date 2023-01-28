# Deployment with Istio

# 1.Step

## Istioctl (Option 1)

### Install with istioctl command line tool
```
istioctl install --set profile=minimal -y
```

## Helm (Option 2)

### Get Repo Info (required only once)
```
helm repo add istio https://istio-release.storage.googleapis.com/charts
```

### Create namespace
```
kubectl create namespace istio-system
```
### Install Charts
```
//Version 1.16.1
helm install istio-base istio/base -n istio-system

//Version 1.16.1
helm install istiod istio/istiod -n istio-system
```

# 2.Step

### Label namespace
```
kubectl label namespace default istio-injection=enabled
```

# 3.Step

### Deploy microservices with loadbalancer (Option 1)
```
kubectl apply -f loadbalancer/microservices-istio.yaml
```

### Deploy microservices with ingress (Option 2)
```
kubectl apply -f ingress/microservices-istio.yaml

kubectl apply -f ingress/ingress.yaml
```

### Deploy with helm (Option 3)
```
helm install microservices ../helm-chart -f microservices-values.yaml  
```
## Miscellaneous

### Forward frontend locally to port 8080 (only for local tests)
```
kubectl port-forward deployment/frontend 8080:8080
```

