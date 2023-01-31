# Deployment with Linkerd

# 1.Step

## Linkerd (Option 1)

### Validate Kubernetes cluster
```
linkerd check --pre
```
### Install with linkerd command line tool onto cluster

```
linkerd install --crds | kubectl apply -f -

// Flag enables root privileges
linkerd install --set proxyInit.runAsRoot=true | kubectl apply -f -
```

## Helm (Option 2)

###  Prerequisite for Helm option

1. Install step - `brew install step`
2. Trust anchor certificate - `step certificate create root.linkerd.cluster.local ca.crt ca.key \
--profile root-ca --no-password --insecure`
3. Issuer certificate and key - `step certificate create identity.linkerd.cluster.local issuer.crt issuer.key \
--profile intermediate-ca --not-after 8760h --no-password --insecure \
--ca ca.crt --ca-key ca.key`

### Get Repo Info (required only once)

```
helm repo add linkerd https://helm.linkerd.io/stable
```
### Install Linkerd CRDs helm chart

```
helm install linkerd-crds linkerd/linkerd-crds -n linkerd --create-namespace
```

### Install Linkerd control plane

```
helm install linkerd-control-plane -n linkerd \
  --set-file identityTrustAnchorsPEM=ca.crt \
  --set-file identity.issuer.tls.crtPEM=issuer.crt \
  --set-file identity.issuer.tls.keyPEM=issuer.key \
  linkerd/linkerd-control-plane
```

# 2.Step

## Deploy microservices with loadbalancer (Option 1)
```
kubectl apply -f loadbalancer/microservices-linkerd.yaml
```

## Deploy microservices with ingress (Option 2)
```
kubectl apply -f ingress/microservices-linkerd.yaml

kubectl apply -f ingress/ingress.yaml
```

## Deploy microservices with ingress (Option 3)
```
helm install microservices ../helm-chart -f microservices-values.yaml 
```

# Miscellaneous
## Forward frontend locally to port 8080 (only for local tests)
```
kubectl port-forward deployment/frontend 8080:8080
```

## Check data plane
```
linkerd check --proxy
```

# Linkerd + HAProxy (Ingress)

## Get Repo Info (required only once)
```
helm repo add haproxytech https://haproxytech.github.io/helm-charts
```
## Install Chart
```
helm install kubernetes-ingress haproxytech/kubernetes-ingress --set controller.service.type=LoadBalancer
```
## Deploy microservices
```
kubectl apply -f ingress-haproxy/microservices-linkerd.yaml
```
## Deploy ingress (Option 3)
```
kubectl apply -f ingress-haproxy/ingress-haproxy.yaml
```
