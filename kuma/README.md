# Deployment with Kuma

# 1.Step

## Get Repo Info (required only once)
```
helm repo add kuma https://kumahq.github.io/charts
```

# 2.Step

## Install Chart
```
helm install --create-namespace --namespace kuma-system kuma kuma/kuma
```

# 3 .Step
## Deploy microservices with loadbalancer (Option 1)
```
kubectl apply -f loadbalancer/microservices-kuma.yaml
```

## Deploy microservices with ingress (Option 2)
```
kubectl apply -f ingress/microservices-kuma.yaml

kubectl apply -f ingress/ingress.yaml
```

## Deploy microservices with helm (Option 3)
```
helm install microservices ../helm-chart -f microservices-values.yaml 
```

## Miscellaneous

### Forward frontend locally to port 8080 (only for local tests)
```
kubectl port-forward deployment/frontend 8080:8080
```

### Kuma GUI
```
kubectl port-forward svc/kuma-control-plane -n kuma-system 5681:5681
```

# Linkerd + HAProxy (Ingress)

## Get Repo Info (required only once)
```
helm repo add haproxytech https://haproxytech.github.io/helm-charts
```
## Install Kuma Chart
```
helm install --create-namespace --namespace kuma-system kuma kuma/kuma
```
## Install HAProxy Chart
```
helm install kubernetes-ingress haproxytech/kubernetes-ingress --set controller.service.type=LoadBalancer
```
## Deploy microservices
```
kubectl apply -f ingress-haproxy/microservices-kuma.yaml
```
## Deploy ingress (Option 3)
```
kubectl apply -f ingress-haproxy/ingress-haproxy.yaml
```