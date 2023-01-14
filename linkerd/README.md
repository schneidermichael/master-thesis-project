# Deployment with Linkerd

## Validate Kubernetes cluster
```
linkerd check --pre
```
## Install Linkerd onto cluster

```
linkerd install --crds | kubectl apply -f -

// Flag enables root privileges
linkerd install --set proxyInit.runAsRoot=true | kubectl apply -f -
```

## Deploy microservices

```
kubectl apply -f ingress/microservices-linkerd.yaml
```

## Inject sidecars/proxies to Microservices application

```
kubectl get deploy -o yaml \
  | linkerd inject - \
  | kubectl apply -f -
```
### Loadbalancer or ingress

## Loadbalancer
```
kubectl apply -f loadbalancer/microservices-linkerd.yaml
```

## Ingress
```
kubectl apply -f ingress/ingress.yaml
```

## Forward frontend locally to port 8080 (only for local tests)
```
kubectl port-forward deployment/frontend 8080:8080
```

## Check data plane
```
linkerd -n microservices check --proxy
```