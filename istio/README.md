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

### Install Charts
```
//Version 1.16.1
helm install istio-base istio/base --create-namespace -n istio-system

//Version 1.16.1
helm install istiod istio/istiod -n istio-system
```

# 2.Step

### Label namespace (Option 1) 
```
kubectl label namespace default istio-injection=enabled
```

### Label on Pod (Option 2)  - Deployment completed with this option

```
kubectl apply -f loadbalancer/microservices-istio-pod.yaml
```

# 3.Step

### Deploy microservices with loadbalancer (Option 1)
```
kubectl apply -f loadbalancer/microservices-istio.yaml
```

### Deploy microservices with ingress (Option 2)
```
kubectl apply -f ingress/microservices-istio-namespace.yaml

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

# Istio + HAProxy (Ingress)

## Get Repo Info (required only once)
```
helm repo add haproxytech https://haproxytech.github.io/helm-charts
```
## Install HAProxy Chart
```
helm install kubernetes-ingress haproxytech/kubernetes-ingress --set controller.service.type=LoadBalancer
```
## Install Istio
```
istioctl install --set profile=minimal -y
```
## Deploy microservices
```
kubectl label namespace default istio-injection=enabled

kubectl apply -f ingress/microservices-istio.yaml
```
## Deploy ingress
```
kubectl apply -f ingress-haproxy/ingress-haproxy.yaml
```
