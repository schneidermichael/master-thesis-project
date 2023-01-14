# Linkerd Monitoring 

## Viz & Prometheus

```
linkerd viz install | kubectl apply -f -

linkerd viz dashboard &
```
Uninstall viz extension
```
linkerd viz uninstall | kubectl delete -f -
```


# Grafana

Use Helm Charts

```
helm repo add grafana https://grafana.github.io/helm-charts

helm install grafana -n grafana --create-namespace grafana/grafana \
  -f https://raw.githubusercontent.com/linkerd/linkerd2/main/grafana/values.yaml

linkerd viz install --set grafana.url=grafana.grafana:3000 \
  | kubectl apply -f -

linkerd viz dashboard &
```

# Jäger

```
helm repo add linkerd https://helm.linkerd.io/stable

helm install linkerd-jaeger -n linkerd-jaeger --create-namespace linkerd/linkerd-jaeger

linkerd jaeger dashboard
```