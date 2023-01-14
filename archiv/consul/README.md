# Deployment with Consul

```
helm repo add hashicorp https://helm.releases.hashicorp.com
```

```
helm install -f consul-values.yaml consul hashicorp/consul --create-namespace --namespace consul

kubectl apply -f microservices-consul.yaml

helm install example kong/kong --version 2.3.0 --values kong-values.yaml

kubectl apply -f kong-service-default.yaml

kubectl port-forward deployment/frontend 8080:8080

kubectl apply -f ingress.yaml

kubectl apply -f service-intentions.yaml
```


Only for updating helm start
```
helm upgrade --values consul-values.yaml consul hashicorp/consul --namespace consul
```


Forward frontend locally to port 8080
```
kubectl port-forward deployment/frontend 8080:8080
```

UI in Consul
```
kubectl -n consul port-forward service/consul-ui 8090:80 

or

minikube service consul-ui --namespace consul
```

## Prometheus and Grafana

### Prometheus
```
helm install -f prometheus-values.yaml prometheus prometheus-community/prometheus --version "15.5.3" 
```

### Grafana
```
helm install -f grafana-values.yaml grafana grafana/grafana --version "6.23.1" 

kubectl port-forward svc/grafana 3000:3000
```