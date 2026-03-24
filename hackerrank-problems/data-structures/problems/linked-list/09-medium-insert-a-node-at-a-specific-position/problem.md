# 특정 위치에 노드 삽입

**난이도:** 중
**출처:** [Insert a node at a specific position in a linked list](https://www.hackerrank.com/challenges/insert-a-node-at-a-specific-position-in-a-linked-list)

## 문제 설명

단일 연결 리스트의 헤드 포인터, 삽입할 데이터 값, 삽입할 위치(0-based)가 주어진다. 해당 위치에 새 노드를 삽입하고, 갱신된 리스트의 헤드 포인터를 반환하시오.

## 입력 형식

- 첫째 줄에 기존 노드의 개수 N이 주어진다
- 다음 N개의 줄에 각 노드의 데이터 값이 주어진다
- 다음 줄에 삽입할 데이터 값이 주어진다
- 마지막 줄에 삽입할 위치 position이 주어진다 (0-based)

## 출력 형식

- 삽입 후 연결 리스트의 각 노드 데이터를 한 줄에 하나씩 출력한다

## 제약 조건

- 1 ≤ N ≤ 1000
- 0 ≤ position ≤ N
- 1 ≤ 각 노드의 데이터 ≤ 1000

## 예제

### 예제 입력 1

```text
3
16
13
7
1
2
```

### 예제 출력 1

```text
16
13
1
7
```
