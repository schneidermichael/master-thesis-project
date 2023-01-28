# Deployment with Traefik Mesh

# 1.Step 

## Get Repo Info (required only once)
```
helm repo add traefik https://traefik.github.io/charts
```

# 2.Step 

## Install Chart
```
helm install traefik-mesh traefik-mesh/traefik-mesh -f traefik-values.yaml --create-namespace --namespace traefik --set kubedns=true
```

# 3.Step 

## Deploy microservices with loadbalancer (Option 1)
```
kubectl apply -f loadbalancer/microservices-traefik-mesh.yaml
```

## Deploy microservices with ingress (Option 2)
```
kubectl apply -f ingress/microservices-traefik-mesh.yaml

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

### Update Chart
```
helm upgrade --values traefik-values.yaml traefik-mesh traefik-mesh/traefik-mesh --namespace traefik
```