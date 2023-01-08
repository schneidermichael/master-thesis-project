# Load testing

## Use case local with result output as CSV
```
k6 --out csv=test_results.csv run ./load-testing/use-case-local.js
```

## Use case local
Running a 30-seconds, 10-VU load test
```
k6 run ./load-testing/use-case-local.js
```

## Use case 1
Running a 10-min, 10-VU load test
```
k6 run ./load-testing/use-case-one.js

k6 --out csv=test_results_use_case_one_istio.csv run ./load-testing/use-case-one.js
```

## Use case 2
Running a 10-min, 50-VU load test
```
k6 run ./load-testing/use-case-two.js

k6 --out csv=test_results_use_case_two_istio.csv run ./load-testing/use-case-two.js
```

## Use case 3
Running a 10-min, 100-VU load test
```
k6 run ./load-testing/use-case-three.js

k6 --out csv=test_results_use_case_three_istio.csv run ./load-testing/use-case-three.js
```

## Use case 4
Running a 10-min, 150-VU load test
```
k6 run ./load-testing/use-case-four.js

k6 --out csv=test_results_use_case_four_istio.csv run ./load-testing/use-case-four.js
```
