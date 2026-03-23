# Heavy-Light Decomposition (HLD)

## 개념

Heavy-Light Decomposition(HLD)이란 트리의 간선을 Heavy Edge와 Light Edge로 분류하고, Heavy Edge로 연결된 경로(Heavy Chain)를 연속 구간으로 매핑하여, 임의의 두 노드 사이의 경로 질의를 효율적으로 처리하는 기법이다.

핵심 아이디어는 **트리의 임의 경로가 최대 O(log N)개의 Heavy Chain으로 분해된다**는 것이다. 각 Heavy Chain을 세그먼트 트리의 연속 구간에 대응시키면, 경로 합·최댓값·갱신 등의 연산을 O(log²N)에 수행할 수 있다.

대표적인 활용:

- **경로 합/최솟값/최댓값 질의**: 두 노드 사이 경로의 합, 최솟값, 최댓값을 O(log²N)에 구한다
- **경로 일괄 갱신**: 경로 위 모든 노드(또는 간선)에 값을 더하거나 변경한다
- **LCA 계산**: HLD 구조를 이용하면 LCA를 O(log N)에 구할 수 있다
- **간선 가중치 경로 질의**: 간선 가중치를 자식 노드에 매핑하여 경로 질의를 처리한다

## 선수 학습

HLD를 학습하기 전에 다음 내용을 먼저 학습하는 것을 권장한다:

- [13-tree](../13-tree/) - 트리의 기본 개념, 순회, 부모-자식 관계
- [25-segment-tree](../25-segment-tree/) - 세그먼트 트리의 구조, 구간 질의/갱신, lazy propagation
- [15-graph-dfs](../15-graph-dfs/) - DFS의 동작 원리, 방문 순서, 스택 기반 탐색

## 동작 원리

### 1. Heavy/Light Edge 분류

루트에서 DFS를 수행하여 각 노드의 서브트리 크기 `sz[v]`를 구한다. 그 뒤 각 노드에서 자식 중 서브트리 크기가 가장 큰 자식으로 향하는 간선을 Heavy Edge, 나머지를 Light Edge로 분류한다.

```
정의:
- Heavy Child: 자식 중 sz가 가장 큰 자식 (동률이면 아무거나)
- Heavy Edge: 부모 → Heavy Child로 향하는 간선
- Light Edge: 부모 → Heavy Child가 아닌 자식으로 향하는 간선
- Heavy Chain: Heavy Edge로 연결된 극대 경로
```

핵심 성질: 임의의 노드에서 루트까지 올라가는 경로에서 Light Edge는 최대 O(log N)개이다. Light Edge를 타고 올라가면 서브트리 크기가 최소 2배 이상 증가하기 때문이다.

### 2. Heavy Chain 구성 및 DFS 번호 부여

두 번째 DFS에서 Heavy Child를 먼저 방문하여 같은 체인에 속하는 노드들이 연속된 DFS 번호를 갖도록 한다.

1. 각 노드 v에 대해 `top[v]` = v가 속한 Heavy Chain의 최상단 노드
2. 루트의 `top[root] = root`
3. Heavy Child u에 대해 `top[u] = top[v]` (같은 체인 연장)
4. Light Child u에 대해 `top[u] = u` (새로운 체인 시작)
5. DFS 순서대로 `pos[v]`를 부여하면, 같은 체인의 노드들은 연속 구간을 형성한다

```
전처리 결과:
- pos[v]: 노드 v의 세그먼트 트리 상 위치 (DFS 번호)
- top[v]: 노드 v가 속한 체인의 최상단 노드
- parent[v]: 노드 v의 부모
- depth[v]: 노드 v의 깊이
- sz[v]: 노드 v의 서브트리 크기
```

### 3. 세그먼트 트리 구축

DFS 번호 순서대로 노드 값을 배열에 배치하고, 그 위에 세그먼트 트리를 구축한다. 같은 체인의 노드들이 연속 구간이므로, 체인 내 구간 질의를 세그먼트 트리로 O(log N)에 처리할 수 있다.

### 4. 경로 질의 분해

두 노드 u, v 사이의 경로 질의를 처리하는 과정:

1. u와 v가 같은 체인에 속할 때까지 반복한다
2. `top[u]`와 `top[v]` 중 깊이가 더 깊은 쪽의 체인을 처리한다
3. 해당 체인에서 `[pos[top[u]], pos[u]]` 구간을 세그먼트 트리로 질의한다
4. u를 `parent[top[u]]`로 올린다 (체인을 넘어감)
5. u와 v가 같은 체인에 속하면 `[pos[u], pos[v]]` 구간을 질의한다 (깊이가 얕은 쪽이 왼쪽)

```
경로 질의 의사코드:
function pathQuery(u, v):
    result = 0 (또는 항등원)
    while top[u] ≠ top[v]:
        if depth[top[u]] < depth[top[v]]:
            swap(u, v)
        result = merge(result, segQuery(pos[top[u]], pos[u]))
        u = parent[top[u]]
    if depth[u] > depth[v]:
        swap(u, v)
    result = merge(result, segQuery(pos[u], pos[v]))
    return result
```

## 복잡도

| 구분 | 복잡도 |
| --- | --- |
| 전처리 (DFS 2회 + 세그먼트 트리 구축) | O(N) |
| 경로 질의 (1회) | O(log²N) |
| 경로 갱신 (1회) | O(log²N) |
| LCA 계산 (1회) | O(log N) |
| 공간 복잡도 | O(N) |

경로 질의/갱신에서 O(log N)개의 체인을 순회하고, 각 체인에서 세그먼트 트리 질의 O(log N)을 수행하므로 총 O(log²N)이다.

## 적합한 문제 유형

- 트리에서 두 노드 사이 경로의 합, 최솟값, 최댓값을 반복적으로 질의하는 문제
- 경로 위 모든 노드(또는 간선)의 값을 일괄 갱신하는 문제
- 경로 질의와 갱신이 혼합된 문제 (갱신 후 질의)
- 간선 가중치에 대한 경로 질의 문제
- LCA를 빠르게 구해야 하는 문제
- 서브트리 질의와 경로 질의가 동시에 필요한 문제 (오일러 투어 + HLD 결합)

## 단계별 추적 (Trace)

### 예시: HLD 구성 및 경로 합 질의

다음 트리에서 HLD를 구성하고, 경로 합을 구한다 (루트: 1).

```
         1 (val=2)
        / \
       2   3 (val=5)
  (val=3) / \
     / \  4   5
    6   7 (val=1)(val=4)
(val=6)(val=8)
```

노드 값: val = [_, 2, 3, 5, 1, 4, 6, 8] (1-indexed)
간선: (1,2), (1,3), (2,6), (2,7), (3,4), (3,5)

**1단계: 서브트리 크기 계산 (1차 DFS)**

```
dfs1(1):
  dfs1(2): sz[6]=1, sz[7]=1 → sz[2]=1+1+1=3
  dfs1(3): sz[4]=1, sz[5]=1 → sz[3]=1+1+1=3
  sz[1] = 1+3+3 = 7
```

```
노드:  1  2  3  4  5  6  7
sz:    7  3  3  1  1  1  1
```

**2단계: Heavy Child 결정**

```
노드 1: 자식 {2(sz=3), 3(sz=3)} → Heavy Child = 2 (동률, 먼저 나온 쪽)
노드 2: 자식 {6(sz=1), 7(sz=1)} → Heavy Child = 6 (동률, 먼저 나온 쪽)
노드 3: 자식 {4(sz=1), 5(sz=1)} → Heavy Child = 4 (동률, 먼저 나온 쪽)
```

**Heavy/Light Edge 표시:**

```
         1
        /H\L
       2    3
      /H\L /H\L
     6   7 4   5

H = Heavy Edge, L = Light Edge
```

**3단계: 2차 DFS (Heavy Child 우선 방문) → DFS 번호 및 체인 구성**

```
dfs2(1, top=1):
  pos[1]=0, top[1]=1
  Heavy Child 2 방문:
    dfs2(2, top=1):
      pos[2]=1, top[2]=1
      Heavy Child 6 방문:
        dfs2(6, top=1):
          pos[6]=2, top[6]=1  (리프)
      Light Child 7 방문:
        dfs2(7, top=7):
          pos[7]=3, top[7]=7  (새 체인)
  Light Child 3 방문:
    dfs2(3, top=3):
      pos[3]=4, top[3]=3  (새 체인)
      Heavy Child 4 방문:
        dfs2(4, top=3):
          pos[4]=5, top[4]=3
      Light Child 5 방문:
        dfs2(5, top=5):
          pos[5]=6, top[5]=5  (새 체인)
```

**결과:**

```
노드:   1  2  3  4  5  6  7
pos:    0  1  4  5  6  2  3
top:    1  1  3  3  5  1  7
depth:  0  1  1  2  2  2  2
```

**Heavy Chain 목록:**

```
체인 A: 1 → 2 → 6  (pos: 0, 1, 2)  top = 1
체인 B: 7           (pos: 3)         top = 7
체인 C: 3 → 4       (pos: 4, 5)      top = 3
체인 D: 5           (pos: 6)         top = 5
```

**세그먼트 트리 배열 (pos 순서):**

```
pos:   0  1  2  3  4  5  6
노드:  1  2  6  7  3  4  5
값:    2  3  6  8  5  1  4
```

**4단계: 경로 합 질의 — query(6, 5)**

노드 6에서 노드 5까지의 경로: 6 → 2 → 1 → 3 → 5, 합 = 6+3+2+5+4 = 20

```
u=6, v=5

반복 1: top[6]=1, top[5]=5 → top이 다름
  depth[top[6]]=0, depth[top[5]]=2 → depth[top[5]]가 더 깊음
  swap → u=5, v=6
  segQuery(pos[top[5]], pos[5]) = segQuery(6, 6) = 4
  u = parent[top[5]] = parent[5] = 3

반복 2: u=3, v=6, top[3]=3, top[6]=1 → top이 다름
  depth[top[3]]=1, depth[top[6]]=0 → depth[top[3]]가 더 깊음
  segQuery(pos[top[3]], pos[3]) = segQuery(4, 4) = 5
  u = parent[top[3]] = parent[3] = 1

반복 3: u=1, v=6, top[1]=1, top[6]=1 → top이 같음! 루프 종료

같은 체인 내 질의:
  depth[1]=0, depth[6]=2 → u=1이 더 얕음
  segQuery(pos[1], pos[6]) = segQuery(0, 2) = 2+3+6 = 11

최종 결과: 4 + 5 + 11 = 20 ✓
```

## 실전 팁

### 활용 노하우

- HLD의 핵심은 "Heavy Child를 먼저 방문하여 같은 체인이 연속 구간이 되게 한다"는 것이다. 이 순서만 지키면 세그먼트 트리와 자연스럽게 결합된다
- 간선 가중치 문제는 각 간선의 가중치를 자식 노드에 매핑한다. 경로 질의 시 LCA 노드는 제외해야 한다 (`segQuery(pos[u]+1, pos[v])`)
- 서브트리 질의도 HLD의 DFS 번호를 이용하면 `[pos[v], pos[v]+sz[v]-1]` 구간으로 처리할 수 있다. 오일러 투어와 동일한 원리이다
- `top` 배열을 이용한 LCA 계산: 두 노드의 `top`이 같아질 때까지 깊은 쪽을 올리고, 같은 체인에서 깊이가 얕은 노드가 LCA이다
- 경로 갱신도 질의와 동일한 체인 분해 과정을 거친다. `segUpdate` 대신 `segRangeUpdate`를 사용하면 된다

### 자주 하는 실수

- 2차 DFS에서 Heavy Child를 먼저 방문하지 않는 실수. Heavy Child를 인접 리스트의 맨 앞으로 swap하거나, 별도로 먼저 처리해야 한다
- 간선 가중치 문제에서 LCA 노드를 포함하여 질의하는 실수. 간선 가중치는 자식 노드에 매핑되므로, 같은 체인 내 질의 시 `pos[u]+1`부터 시작해야 한다
- `top`이 같은 체인에서 깊이 비교를 빠뜨리는 실수. 항상 `depth[u] ≤ depth[v]`가 되도록 swap한 뒤 질의한다
- 세그먼트 트리의 인덱스와 pos 배열의 인덱스를 혼동하는 실수. 0-indexed와 1-indexed를 통일하자

### 엣지 케이스

- N = 1: 노드가 하나뿐인 트리. 경로 질의 시 자기 자신만 포함, 체인 1개
- 일직선 트리 (경로 그래프): 모든 간선이 Heavy Edge, 체인 1개. 경로 질의가 O(log N)으로 최적
- 스타 그래프: 루트에서 모든 리프로 Light Edge. 체인이 N개. 경로 질의 시 체인 2개만 순회
- 경로의 두 끝점이 같은 노드인 경우: 질의 결과 = 해당 노드의 값
- 간선 가중치가 0인 경우: 최댓값 질의에서 0이 답이 될 수 있으므로 초기값을 -∞로 설정

## 관련 알고리즘 비교

| 알고리즘 | 특징 | 적합한 상황 |
| --- | --- | --- |
| HLD + 세그먼트 트리 | 경로를 O(log N)개 구간으로 분할, O(log²N) 질의 | 경로 질의/갱신 반복 |
| 오일러 투어 + 세그먼트 트리 | 서브트리를 연속 구간으로 변환, O(log N) 질의 | 서브트리 질의/갱신 반복 |
| LCA (Binary Lifting) | 두 노드의 최소 공통 조상, O(log N) | LCA만 필요한 경우 |
| 센트로이드 분할 | 트리 분할 정복, O(N log N) | 경로 조건 쌍 세기, 오프라인 질의 |
| Link-Cut Tree | 동적 트리, amortized O(log N) | 간선 추가/삭제가 있는 동적 트리 |
| 트리 DP | 서브트리 단위 DP, O(N) | 한 번의 전처리로 결과 계산 |

- 경로 질의/갱신 반복 → HLD
- 서브트리 질의/갱신 반복 → 오일러 투어
- LCA만 필요 → Binary Lifting
- 경로 조건 쌍 세기 → 센트로이드 분할
- 동적 트리 (간선 추가/삭제) → Link-Cut Tree
