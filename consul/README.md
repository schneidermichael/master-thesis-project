# Deployment with Consul

```
helm repo add hashicorp https://helm.releases.hashicorp.com
```

```
helm install --values helm-consul-values.yaml consul hashicorp/consul --create-namespace --namespace consul --version "0.43.0"
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
