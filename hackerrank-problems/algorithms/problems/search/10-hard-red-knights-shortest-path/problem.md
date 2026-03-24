# 레드 나이트의 최단 경로

**난이도:** 상
**출처:** [Red Knight's Shortest Path](https://www.hackerrank.com/challenges/red-knights-shortest-path)

## 문제 설명

N×N 체스판에서 레드 나이트는 다음 6가지 방향으로 이동할 수 있다:
- UL (Upper Left): (r-2, c-1)
- UR (Upper Right): (r-2, c+1)
- R (Right): (r, c+2)
- LR (Lower Right): (r+2, c+1)
- LL (Lower Left): (r+2, c-1)
- L (Left): (r, c-2)

시작 위치에서 목표 위치까지의 최소 이동 횟수와 경로를 출력하시오. 최소 이동 경로가 여러 개이면 위의 우선순위(UL > UR > R > LR > LL > L)에 따라 사전순으로 가장 빠른 경로를 출력한다. 도달 불가능하면 "Impossible"을 출력한다.

## 입력 형식

- 첫째 줄에 체스판 크기 N이 주어진다
- 둘째 줄에 시작 위치 (r_start, c_start)가 주어진다
- 셋째 줄에 목표 위치 (r_end, c_end)가 주어진다

## 출력 형식

- 도달 가능하면 첫째 줄에 최소 이동 횟수, 둘째 줄에 이동 방향을 공백으로 구분하여 출력한다
- 도달 불가능하면 "Impossible"을 출력한다

## 제약 조건

- 5 ≤ N ≤ 200
- 0 ≤ r_start, c_start, r_end, c_end < N

## 예제

### 예제 입력 1

```text
7
6 6
0 1
```

### 예제 출력 1

```text
4
UL UL UL L
```

### 예제 입력 2

```text
6
5 1
0 5
```

### 예제 출력 2

```text
Impossible
```
