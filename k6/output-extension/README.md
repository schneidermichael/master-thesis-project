```
xk6 build --with xk6-output-logger=.
```

```
./k6 run script.js -o smp --quiet --no-summary --vus 5 --duration 10s
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
