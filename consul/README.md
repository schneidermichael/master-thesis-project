# Deployment with Consul

```
helm repo add hashicorp https://helm.releases.hashicorp.com
```

```
helm install -f consul-values.yaml consul hashicorp/consul --create-namespace --namespace consul --version "0.43.0" 
```
```
kubectl create namespace microservices | kubectl apply -f microservices-consul.yaml
```

Forward frontend locally to port 8080
```
kubectl -n microservices port-forward deployment/frontend 8080:8080
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