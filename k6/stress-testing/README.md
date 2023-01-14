# Stress testing

## Use case local
Running a 
- 5 seconds, 100 VU load test
- 5 seconds, 200 VU load test
- 5 seconds, 300 VU load test
- 5 seconds, 400 VU load test
- 5 seconds, 500 VU load test

```
k6 run use-case-local.js ✅
```

## Use case production

Running a 
- 2 minutes, 100 VU load test
- 2 minutes, 200 VU load test
- 2 minutes, 300 VU load test
- 2 minutes, 400 VU load test
- 2 minutes, 500 VU load test

### Istio
```
k6 --out csv=test_results_use_case_production_istio.csv run use-case-production.js 
```
### Linkerd
```
k6 --out csv=test_results_use_case_production_linkerd.csv run use-case-production.js ✅
```
### Kuma
```
k6 --out csv=test_results_use_case_production_kuma.csv run use-case-production.js
```
### Traefik Mesh
```
k6 --out csv=test_results_use_case_production_traefik.csv run use-case-production.js
```