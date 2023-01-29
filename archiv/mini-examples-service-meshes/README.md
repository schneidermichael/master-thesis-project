# Start Meshery with KinD

Creating a Kubernetes cluster with KinD

```zsh
kind create cluster
```

Set the KUBECONFIG environment

```zsh
export KUBECONFIG=${HOME}/.kube/config
```

check the connection of the cluster

```zsh
kubectl cluster-info --context kind-kind
```

Deploy Meshery to your Kubernetes
```zsh
mesheryctl system start
```

Only works if i change the loadbalancer to nodeport

Delete KinD cluster

```zsh
kind delete cluster --name kind
```


# Start Meshery with minikube

If a minikube cluster exist

```zsh
minikube delete
```

Start minikube

```zsh
minikube start --cpus 4 --memory 8192
```

Check minikube cluster
```zsh
minikube status
```

Deploy Meshery to your Kubernetes cluster(minikube)
```zsh
mesheryctl system start
```
Expose the LoadBalancer services
```zsh
minikube tunnel
```

Go to http://localhost:9081/

# Connect Meshery with Prometheus (without prometheus.yml) - External Docker

Start Prometheus

```zsh
docker run -d --name prometheus -p 9090:9090 prom/prometheus 
```

In Meshery under Settings connect to http://host.docker.internal:9090


# Connect Meshery with Grafana - External Docker

Start Grafana

```zsh
docker run -d --name grafana -p 3000:3000 grafana/grafana-enterprise
```


In Meshery under Settings connect to http://host.docker.internal:3000


# FAQ meshery

- It is possible that the endpoint of the meshery config is wrong in ~/.meshery/config.yaml 

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

Install Emojivoto application into the emojivoto namespace
```
curl --proto '=https' --tlsv1.2 -sSfL https://run.linkerd.io/emojivoto.yml \
  | kubectl apply -f -
```

Forward web-svc locally to port 8080 
```
kubectl -n emojivoto port-forward svc/web-svc 8080:80

```
Inject sidecars/proxies to Emojivoto application

```
kubectl get -n emojivoto deploy -o yaml \
  | linkerd inject - \
  | kubectl apply -f -
```

Check data plane
```
linkerd -n emojivoto check --proxy
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

# Deployment with Istio

```
istioctl install --set profile=default -y
```
```
kubectl create namespace emojivoto
```

```
kubectl label namespace emojivoto istio-injection=enabled
```

```
kubectl apply -f emojivoto.yml
```

Forward web-svc locally to port 8080
```
kubectl -n emojivoto port-forward svc/web-svc 8080:80
```

Ensure that there are no issues with the configuration

```
istioctl analyze -n emojivoto
```

# Deployment with Kuma

```
kumactl install control-plane | kubectl apply -f -
```

```
kubectl apply -f emojivoto.yml
```

```
kubectl -n emojivoto port-forward svc/web-svc 8080:80
```

UI in Kuma

```
kubectl port-forward svc/kuma-control-plane -n kuma-system 5681:5681
```

# Deployment with Traefik mesh

```
helm repo add traefik-mesh https://helm.traefik.io/mesh
```

```
helm repo update
```


```
helm install traefik-mesh traefik-mesh/traefik-mesh
```

```
kubectl apply -f emojivoto.yml 
```

```
kubectl -n emojivoto port-forward svc/web-svc 8080:80
```

# Deployment with Open Service Mesh

```
osm install --mesh-name "osm-system" --osm-namespace "osm" --set=osm.enablePermissiveTrafficPolicy=true --set=osm.deployPrometheus=true --set=osm.deployGrafana=true --set=osm.deployJaeger=true
```

```
kubectl create namespace emojivoto
````

```
osm namespace add emojivoto
```

```
osm --mesh-name osm-system namespace add emojivoto
```

```
kubectl apply -f emojivoto.yml 
```

```
kubectl -n emojivoto port-forward svc/web-svc 8080:80
```

# Deployment with Consul

```
helm repo add hashicorp https://helm.releases.hashicorp.com
```

```
helm install --values helm-consul-values.yaml consul hashicorp/consul --create-namespace --namespace consul --version "0.43.0"
```

```
kubectl apply -f emojivoto.yml 
```

```
kubectl -n emojivoto port-forward svc/web-svc 8090:80
```

UI in Consul
```
kubectl -n consul port-forward service/consul-ui 8080:80 

or

minikube service consul-ui --namespace consul
```





# Reference Documentation

## Service Mesh

- [Linkerd](https://linkerd.io/)
- [Istio](https://istio.io/latest/docs/setup/getting-started/)
- [Kuma](https://kuma.io/)
- [Traefik Mesh](https://doc.traefik.io/traefik-mesh/)
- [Open Service Mesh](https://openservicemesh.io/)
- [Consul](https://developer.hashicorp.com/consul/tutorials/kubernetes-deploy/kubernetes-minikube)
- [Meshery](https://docs.meshery.io/)


## Local kubernetes

- [KinD](https://kind.sigs.k8s.io/)
- [minikube](https://minikube.sigs.k8s.io/)

## Application Definition & Image Build  

- [Helm](https://helm.sh/)

## Monitoring

- [Grafana](https://grafana.com/docs/grafana/latest/setup-grafana/installation/)
- [Prometheus](https://prometheus.io/)

## Package manager for macOS

- [Homebrew](https://brew.sh/)




