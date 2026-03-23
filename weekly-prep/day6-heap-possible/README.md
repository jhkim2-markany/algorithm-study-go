# Day 6: 힙 + POSSIBLE 유형

> S3~G4 문제 수: 힙 113개 + 최단경로 152개 + 트리 120개 + 유니온파인드 60개 + 플러드필 57개
> 예상 소요: 2~3시간

힙은 HackerRank 자료구조 단골. 나머지는 골드 경계에서 나올 수 있는 유형.

---

## 1. 힙 / 우선순위 큐

> 📂 [12-heap-and-priority-queue](../../12-heap-and-priority-queue/) → theory.md → examples → problems

### 핵심 개념
- `container/heap` 인터페이스 구현이 핵심
- 최솟값/최댓값을 O(log N)에 삽입/삭제
- K번째 원소, 중앙값, 스케줄링 문제에 사용

### container/heap 템플릿 (암기)

```go
import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    x := old[len(old)-1]
    *h = old[:len(old)-1]
    return x
}

// 사용법
h := &MinHeap{}
heap.Init(h)
heap.Push(h, 3)
heap.Push(h, 1)
min := heap.Pop(h).(int) // 1
```

### 시험에 나오는 패턴

**패턴 1: K번째 큰 원소 (최소 힙, 크기 K 유지)**
```go
h := &MinHeap{}
heap.Init(h)
for _, v := range arr {
    heap.Push(h, v)
    if h.Len() > k {
        heap.Pop(h) // 가장 작은 것 제거
    }
}
kthLargest := (*h)[0] // 힙의 루트 = K번째 큰 값
```

**패턴 2: 구조체 힙 (다중 필드)**
```go
type Item struct {
    value    string
    priority int
}
type PQ []*Item

func (pq PQ) Len() int            { return len(pq) }
func (pq PQ) Less(i, j int) bool  { return pq[i].priority < pq[j].priority }
func (pq PQ) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x interface{}) { *pq = append(*pq, x.(*Item)) }
func (pq *PQ) Pop() interface{} {
    old := *pq
    x := old[len(old)-1]
    *pq = old[:len(old)-1]
    return x
}
```

### 연습 문제
- [ ] [01-easy-kth-largest](../../12-heap-and-priority-queue/problems/01-easy-kth-largest/)
- [ ] [02-medium-merge-k-sorted-lists](../../12-heap-and-priority-queue/problems/02-medium-merge-k-sorted-lists/)

---

## 2. 최단 경로 (다익스트라)

> 📂 [21-shortest-path](../../21-shortest-path/) → theory.md → problems/01-easy

### 핵심: 다익스트라 = BFS + 힙

```go
type Edge struct{ to, cost int }
type Node struct{ v, dist int }
type MinHeap []Node
// ... heap 인터페이스 구현 (Less: dist 기준)

dist := make([]int, n+1)
for i := range dist {
    dist[i] = math.MaxInt64
}
dist[start] = 0

h := &MinHeap{}
heap.Push(h, Node{start, 0})

for h.Len() > 0 {
    cur := heap.Pop(h).(Node)
    if cur.dist > dist[cur.v] {
        continue
    }
    for _, e := range graph[cur.v] {
        nd := cur.dist + e.cost
        if nd < dist[e.to] {
            dist[e.to] = nd
            heap.Push(h, Node{e.to, nd})
        }
    }
}
```

### 연습 문제
- [ ] [01-easy-single-source-shortest](../../21-shortest-path/problems/01-easy-single-source-shortest/)

---

## 3. 트리

> 📂 [13-tree](../../13-tree/) → theory.md → problems/01-easy, 02-medium

### 핵심 패턴: 트리 순회

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// 전위 (루트 → 왼 → 오)
func preorder(node *TreeNode) {
    if node == nil { return }
    visit(node.Val)
    preorder(node.Left)
    preorder(node.Right)
}

// 중위 (왼 → 루트 → 오) — BST에서 정렬 순서
func inorder(node *TreeNode) {
    if node == nil { return }
    inorder(node.Left)
    visit(node.Val)
    inorder(node.Right)
}

// 후위 (왼 → 오 → 루트)
func postorder(node *TreeNode) {
    if node == nil { return }
    postorder(node.Left)
    postorder(node.Right)
    visit(node.Val)
}
```

### 연습 문제
- [ ] [01-easy-tree-traversal](../../13-tree/problems/01-easy-tree-traversal/)
- [ ] [02-medium-tree-diameter](../../13-tree/problems/02-medium-tree-diameter/)

---

## 4. 유니온 파인드

> 📂 [20-union-find](../../20-union-find/) → theory.md → problems/01-easy

### 핵심 템플릿

```go
parent := make([]int, n)
for i := range parent {
    parent[i] = i
}
var find func(int) int
find = func(x int) int {
    if parent[x] != x {
        parent[x] = find(parent[x]) // 경로 압축
    }
    return parent[x]
}
union := func(a, b int) {
    parent[find(a)] = find(b)
}
// 같은 그룹인지 확인
sameGroup := func(a, b int) bool {
    return find(a) == find(b)
}
```

### 연습 문제
- [ ] [01-easy-connected-components](../../20-union-find/problems/01-easy-connected-components/)

---

## 5. 플러드 필

> 📂 [59-flood-fill](../../59-flood-fill/) → theory.md → problems/01-easy

### 핵심: BFS/DFS로 연결 영역 탐색

Day 3의 BFS 격자 탐색과 동일한 패턴. 시작점에서 같은 색/값인 인접 칸을 모두 방문.

```go
// BFS 플러드 필
func floodFill(grid [][]int, sr, sc, newColor int) {
    oldColor := grid[sr][sc]
    if oldColor == newColor { return }

    dx := []int{-1, 1, 0, 0}
    dy := []int{0, 0, -1, 1}
    queue := [][]int{{sr, sc}}
    grid[sr][sc] = newColor

    for len(queue) > 0 {
        cur := queue[0]; queue = queue[1:]
        for d := 0; d < 4; d++ {
            nx, ny := cur[0]+dx[d], cur[1]+dy[d]
            if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) && grid[nx][ny] == oldColor {
                grid[nx][ny] = newColor
                queue = append(queue, []int{nx, ny})
            }
        }
    }
}
```

### 연습 문제
- [ ] [01-easy-count-regions](../../59-flood-fill/problems/01-easy-count-regions/)

---

## Day 6 체크리스트

- [ ] `container/heap` 인터페이스 5개 메서드 암기 + 직접 타이핑
- [ ] K번째 큰 원소 패턴 이해
- [ ] 힙 easy + medium 풀기
- [ ] 다익스트라 = BFS + 힙 패턴 이해
- [ ] 최단경로 easy 풀기
- [ ] 트리 순회 3종 직접 타이핑
- [ ] 트리 easy + medium 풀기
- [ ] Union-Find 템플릿 직접 타이핑
- [ ] 유니온파인드 easy 풀기
- [ ] 플러드 필 easy 풀기
- [ ] **Day 1~5 중 약한 부분 medium 재풀이**
