# Deployment with Consul+Kong

```
helm install -f consul-values.yaml consul hashicorp/consul 

kubectl apply -f proxy-default.yaml

helm repo add kong https://charts.konghq.com

helm install -n default example kong/kong -f kong/values.yaml

```

Token for UI
```
export CONSUL_HTTP_TOKEN=$(kubectl get secrets consul-bootstrap-acl-token -o jsonpath='{.data.token}' | base64 -d)
```