# Deployment with Open Service Mesh (OSM)

## Get Repo Info (required only once)
```
helm repo add osm https://openservicemesh.github.io/osm
```

## Create namespace
```
kubectl create namespace osm-system
```
## Resource limits (OSM specific)
```
kubectl apply -f loadbalancer/limits.yaml
```
## Install Chart
```
helm install osm osm --repo https://openservicemesh.github.io/osm --namespace osm-system --values osm-values.yaml --set osm.enableReconciler=true
```

## Label namespace
```
kubectl label namespace default openservicemesh.io/sidecar-injection=enabled openservicemesh.io/monitored-by=osm
kubectl annotate namespace default openservicemesh.io/sidecar-injection=enabled
```

## Deploy microservices
```
kubectl apply -f loadbalancer/microservices-osm.yaml
```

## Update Chart
```

```

## Forward frontend locally to port 8080 (only for local tests)
```
kubectl port-forward deployment/frontend 8080:8080
```
