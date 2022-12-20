# Deployment with Linkerd

Validate Kubernetes cluster

```
linkerd check --pre
```

Install Linkerd onto cluster

```
linkerd install --crds | kubectl apply -f -

// Flag enables root privileges
linkerd install --set proxyInit.runAsRoot=true | kubectl apply -f -
```

```
kubectl create namespace microservices | kubectl apply -f microservices-linkerd.yaml
```

Forward frontend locally to port 8080
```
kubectl -n microservices port-forward deployment/frontend 8080:8080
```

Inject sidecars/proxies to Microservices application

```
kubectl get -n microservices deploy -o yaml \
  | linkerd inject - \
  | kubectl apply -f -
```

Check data plane
```
linkerd -n microservices check --proxy
```

Install viz extension

```
linkerd viz install | kubectl apply -f -

linkerd viz dashboard &
```
Uninstall viz extension
```
linkerd viz uninstall | kubectl delete -f -
```

# Install Grafana in Kubernetes for Linkerd

Use Helm Charts

```
helm repo add grafana https://grafana.github.io/helm-charts

helm install grafana -n grafana --create-namespace grafana/grafana \
  -f https://raw.githubusercontent.com/linkerd/linkerd2/main/grafana/values.yaml

linkerd viz install --set grafana.url=grafana.grafana:3000 \
  | kubectl apply -f -

linkerd viz dashboard &
```

# Install Jäger in Kubernetes for Linkerd

```
helm repo add linkerd https://helm.linkerd.io/stable

helm install linkerd-jaeger -n linkerd-jaeger --create-namespace linkerd/linkerd-jaeger

linkerd jaeger dashboard
```