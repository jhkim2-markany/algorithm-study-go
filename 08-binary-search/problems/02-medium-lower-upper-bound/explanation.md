# Lower/Upper Bound - 해설

## 접근 방식

1. N개의 정수를 배열에 저장한다 (이미 오름차순 정렬되어 있음)
2. 각 질의마다 lower_bound와 upper_bound를 이진 탐색으로 구한다
3. upper_bound - lower_bound가 해당 값의 등장 횟수이다
4. 결과를 출력한다

## 핵심 아이디어

정렬된 배열에서 특정 값의 개수를 세려면, 해당 값이 시작되는 위치(lower_bound)와 끝나는 다음 위치(upper_bound)를 알면 된다. 두 위치의 차이가 곧 등장 횟수이다.

- **Lower Bound**: target 이상인 첫 번째 위치. `arr[mid] < target`이면 lo = mid + 1, 아니면 hi = mid.
- **Upper Bound**: target 초과인 첫 번째 위치. `arr[mid] <= target`이면 lo = mid + 1, 아니면 hi = mid.

두 함수 모두 `lo < hi` 조건으로 반복하며, 초기 범위는 `[0, len(arr))`이다. hi를 len(arr)로 설정하는 이유는 모든 원소가 target 이하인 경우 결과가 len(arr)이 될 수 있기 때문이다.

## 복잡도 분석

| 구분 | 복잡도 |
| --- | --- |
| 시간 복잡도 | O(M × log N) |
| 공간 복잡도 | O(N) |

각 질의마다 lower_bound와 upper_bound를 각각 O(log N)에 수행하므로 전체 시간 복잡도는 O(M × log N)이다.

## 대안적 접근

- **해시맵 사용**: 각 값의 등장 횟수를 미리 세어두면 질의당 O(1)에 답할 수 있다. 다만 이 문제는 이진 탐색의 lower_bound/upper_bound 활용을 연습하기 위한 것이다.
- **선형 탐색**: O(N×M)으로 시간 초과 위험이 있다.
