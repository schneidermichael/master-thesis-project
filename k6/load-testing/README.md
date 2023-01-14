# Load testing

## Use case local with result output as CSV
```
k6 --out csv=test_results.csv run 0_use-case-local.js
```

## Use case local
Running a 30-seconds, 10-VU load test
```
k6 run 0_use-case-local.js
```

## Use case 1
Running a 10-min, 25-VU load test
### Istio
```
k6 --out csv=test_results_use_case_one_istio.csv run 1_use-case-one.js
```
### Linkerd
```
k6 --out csv=test_results_use_case_one_linkerd.csv run 1_use-case-one.js
```
### Kuma
```
k6 --out csv=test_results_use_case_one_kuma.csv run 1_use-case-one.js
```
### Traefik Mesh
```
k6 --out csv=test_results_use_case_one_traefik.csv run 1_use-case-one.js
```

## Use case 2
Running a 10-min, 50-VU load test
### Istio
```
k6 --out csv=test_results_use_case_two_istio.csv run 2_use-case-two.js
```
### Linkerd
```
k6 --out csv=test_results_use_case_two_linkerd.csv run 2_use-case-two.js
```
### Kuma
```
k6 --out csv=test_results_use_case_two_kuma.csv run 2_use-case-two.js
```
### Traefik Mesh
```
k6 --out csv=test_results_use_case_one_traefik.csv run 2_use-case-one.js
```

## Use case 3
Running a 10-min, 100-VU load test

### Istio
```
k6 --out csv=test_results_use_case_three_istio.csv run 3_use-case-three.js
```
### Linkerd
```
k6 --out csv=test_results_use_case_three_linkerd.csv run 3_use-case-three.js
```
### Kuma
```
k6 --out csv=test_results_use_case_three_kuma.csv run 3_use-case-three.js
```
### Traefik Mesh
```
k6 --out csv=test_results_use_case_three_traefik.csv run 3_use-case-three.js
```