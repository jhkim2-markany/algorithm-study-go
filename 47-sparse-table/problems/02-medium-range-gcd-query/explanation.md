# 구간 GCD 쿼리 - 해설

## 접근 방식

1. 배열을 입력받는다
2. GCD 연산을 사용하여 Sparse Table을 O(N log N)에 전처리한다
3. 각 쿼리에 대해 O(1)에 구간 GCD를 응답한다

## 핵심 아이디어

GCD(최대공약수)는 멱등 연산이다. 즉, gcd(a, a) = a이므로 같은 원소를 여러 번 GCD 연산에 포함시켜도 결과가 변하지 않는다. 따라서 Sparse Table의 겹치는 두 구간 기법을 그대로 적용할 수 있다.

구간 [L, R]의 GCD를 구할 때, k = floor(log₂(R-L+1))로 설정하면 [L, L+2^k-1]과 [R-2^k+1, R] 두 구간이 [L, R] 전체를 덮는다. 두 구간의 GCD를 다시 GCD하면 전체 구간의 GCD가 된다.

Sparse Table 구성 시 점화식은 다음과 같다:
- sparse[k][i] = gcd(sparse[k-1][i], sparse[k-1][i + 2^(k-1)])

## 복잡도 분석

| 구분 | 복잡도 |
| --- | --- |
| 시간 복잡도 | O(N log N · log(max) + M · log(max)) |
| 공간 복잡도 | O(N log N) |

GCD 연산 자체가 O(log(max))이므로 전처리에 O(N log N · log(max)), 각 쿼리에 O(log(max))가 소요된다. 실제로는 GCD 연산이 매우 빠르게 수렴하므로 상수가 작다.

## 대안적 접근

- 세그먼트 트리: 구간 GCD를 O(log N · log(max))에 구할 수 있으며, 값 갱신도 가능하다.
- 나이브 방법: 매 쿼리마다 구간을 순회하면 O(NM · log(max))으로 비효율적이다.
