# 최장 공통 부분 수열

**난이도:** 중
**출처:** [The Longest Common Subsequence](https://www.hackerrank.com/challenges/dynamic-programming-classics-the-longest-common-subsequence)

## 문제 설명

두 개의 수열이 주어질 때, 최장 공통 부분 수열(LCS)을 구하시오. 부분 수열은 원래 수열에서 일부 원소를 삭제하여 얻을 수 있는 수열이다 (원소의 상대적 순서는 유지).

## 입력 형식

- 첫째 줄에 두 수열의 길이 N과 M이 공백으로 구분되어 주어진다
- 둘째 줄에 첫 번째 수열의 N개 원소가 공백으로 구분되어 주어진다
- 셋째 줄에 두 번째 수열의 M개 원소가 공백으로 구분되어 주어진다

## 출력 형식

- LCS의 원소를 공백으로 구분하여 출력한다

## 제약 조건

- 1 ≤ N, M ≤ 100
- 수열의 원소는 정수이다

## 예제

### 예제 입력 1

```text
5 6
1 2 3 4 1
3 4 1 2 1 3
```

### 예제 출력 1

```text
1 2 3
```
