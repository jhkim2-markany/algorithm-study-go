# 이진 탐색 트리 최소 공통 조상

**난이도:** 중
**출처:** [Binary Search Tree: Lowest Common Ancestor](https://www.hackerrank.com/challenges/binary-search-tree-lowest-common-ancestor)

## 문제 설명

이진 탐색 트리(BST)의 루트 노드와 두 개의 값 v1, v2가 주어진다. 두 값의 최소 공통 조상(Lowest Common Ancestor, LCA)을 찾아 그 노드의 데이터를 출력하시오.

최소 공통 조상은 두 노드의 공통 조상 중 가장 깊은(루트에서 가장 먼) 노드이다. 노드 자기 자신도 자신의 조상으로 간주한다.

## 입력 형식

- 첫째 줄에 노드의 개수 N이 주어진다
- 다음 N개의 줄에 각 노드의 데이터가 주어진다 (BST 삽입 순서대로)
- 마지막 줄에 두 값 v1, v2가 공백으로 구분되어 주어진다

## 출력 형식

- 최소 공통 조상 노드의 데이터를 출력한다

## 제약 조건

- 1 ≤ N ≤ 25
- 1 ≤ 각 노드의 데이터 ≤ 25
- v1 ≠ v2
- v1, v2는 트리에 존재하는 값이다

## 예제

### 예제 입력 1

```text
6
4
2
7
1
3
6
1 7
```

### 예제 출력 1

```text
4
```
