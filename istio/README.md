# Deployment with Istio


## Istioctl (Option 1)

### Install istioctl command line tool
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

## Deploy microservies (Prerequisite option 1 or option 2)

### Label namespace
```
kubectl label namespace default istio-injection=enabled
```

### Deploy microservices with loadbalancer
```
kubectl apply -f loadbalancer/microservices-istio.yaml
```

### Deploy microservices with ingress
```
kubectl apply -f ingress/microservices-istio.yaml

kubectl apply -f ingress/ingress.yaml
```

## Miscellaneous

### Forward frontend locally to port 8080 (only for local tests)
```
kubectl port-forward deployment/frontend 8080:8080
```

