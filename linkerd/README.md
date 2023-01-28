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
### Install Linkerd CRDs

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

kubectl get deploy -o yaml \
  | linkerd inject - \
  | kubectl apply -f -
```

## Deploy microservices with ingress (Option 2)
```
kubectl apply -f ingress/microservices-linkerd.yaml

kubectl get deploy -o yaml \
  | linkerd inject - \
  | kubectl apply -f -

kubectl apply -f ingress/ingress.yaml
```

## Deploy microservices with ingress (Option 3)
```
helm install microservices ../helm-chart -f microservices-values.yaml 
```
## Forward frontend locally to port 8080 (only for local tests)
```
kubectl port-forward deployment/frontend 8080:8080
```

## Check data plane
```
linkerd -n microservices check --proxy
```