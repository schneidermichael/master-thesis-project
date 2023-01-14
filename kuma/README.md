# Deployment with Kuma

## Get Repo Info (required only once)
```
helm repo add kuma https://kumahq.github.io/charts
```

## Install Chart
```
helm install --create-namespace --namespace kuma-system kuma kuma/kuma
```

## Deploy microservices with loadbalancer
```
kubectl apply -f loadbalancer/microservices-kuma.yaml
```

## Deploy microservices with ingress
```
kubectl apply -f ingress/microservices-kuma.yaml

kubectl apply -f ingress/ingress.yaml
```

## Kuma GUI
```
kubectl port-forward svc/kuma-control-plane -n kuma-system 5681:5681
```