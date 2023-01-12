# Deployment with Kuma

```
helm repo add kuma https://kumahq.github.io/charts
```

```
helm install --create-namespace --namespace kuma-system kuma kuma/kuma

kubectl apply -f microservices-kuma.yaml

kubectl apply -f ingress.yaml
```

Kuma GUI
```
kubectl port-forward svc/kuma-control-plane -n kuma-system 5681:5681
```