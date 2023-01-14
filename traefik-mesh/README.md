# Deployment with Traefik Mesh

## Get Repo Info (required only once)
```
helm repo add traefik https://traefik.github.io/charts
```
## Install Chart
```
helm install traefik-mesh traefik-mesh/traefik-mesh -f traefik-values.yaml --create-namespace --namespace traefik --set kubedns=true
```

## Deploy microservices with loadbalancer
```
kubectl apply -f loadbalancer/microservices-traefik-mesh.yaml
```

## Deploy microservices with ingress
```
kubectl apply -f ingress/microservices-traefik-mesh.yaml

kubectl apply -f ingress/ingress.yaml
```

## Update Chart
```
helm upgrade --values traefik-values.yaml traefik-mesh traefik-mesh/traefik-mesh --namespace traefik
```