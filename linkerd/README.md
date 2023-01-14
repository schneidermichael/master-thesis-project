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

## Deploy microservices with loadbalancer
```
kubectl apply -f loadbalancer/microservices-linkerd.yaml
```

## Deploy microservices with ingress
```
kubectl apply -f ingress/microservices-linkerd.yaml

kubectl apply -f ingress/ingress.yaml
```

## Forward frontend locally to port 8080 (only for local tests)
```
kubectl port-forward deployment/frontend 8080:8080
```

## Inject sidecars/proxies to Microservices application

```
kubectl get deploy -o yaml \
  | linkerd inject - \
  | kubectl apply -f -
```

## Check data plane
```
linkerd -n microservices check --proxy
```