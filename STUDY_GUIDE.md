# HackerRank 코딩테스트 대비 학습 가이드

> **포지션**: 시니어 백엔드 엔지니어 (SDV 클라우드 플랫폼)
> **시험 형식**: HackerRank, 75분
> **예상 난이도**: 백준 실버3 ~ 골드4 (자료구조 중심)
> **마감**: 2026년 3월 30일 오전 10:59 KST
> **학습 기간**: 3/23(월) ~ 3/29(일), 7일

---

## 우선순위 선정 근거

solved.ac API로 백준 S3~G4 범위의 실제 문제 수를 조회하여 우선순위를 결정했다.
기존 테스트 응시자 피드백("자료구조 문제 중심, S3~G4 난이도")을 반영하여
HackerRank에서 자료구조를 직접 묻는 유형(스택, 힙, 해시)은 BOJ 문제 수 대비 가중치를 높였다.

### S3~G4 문제 수 TOP 15 (solved.ac 2026.03.23 기준)

| 순위 | 태그 | S3~G4 | 전체 | S3~G4 비율 |
|:---:|---|---:|---:|---:|
| 1 | implementation (구현) | 1,607 | 6,792 | 24% |
| 2 | bruteforcing (브루트포스) | 1,133 | 2,787 | 41% |
| 3 | greedy (그리디) | 957 | 3,441 | 28% |
| 4 | dp (동적 프로그래밍) | 811 | 5,104 | 16% |
| 5 | sorting (정렬) | 761 | 2,400 | 32% |
| 6 | graph_traversal (그래프 탐색) | 727 | 2,488 | 29% |
| 7 | string (문자열) | 693 | 2,946 | 24% |
| 8 | simulation (시뮬레이션) | 500 | 1,358 | 37% |
| 9 | bfs | 468 | 1,223 | 38% |
| 10 | set (집합/해시셋) | 445 | 1,602 | 28% |
| 11 | prefix_sum (누적합) | 317 | 1,361 | 23% |
| 12 | backtracking (백트래킹) | 308 | 652 | **47%** |
| 13 | binary_search (이진 탐색) | 298 | 1,650 | 18% |
| 14 | dfs | 287 | 968 | 30% |
| 15 | shortest_path (최단경로) | 152 | 980 | 16% |

참고: stack(139), two_pointer(133), trees(120), priority_queue(113)은 BOJ 기준 문제 수는 적지만,
HackerRank는 "자료구조" 카테고리에서 스택/힙/해시를 직접 출제하므로 별도 가중치 적용.

---

## 학습 우선순위

### 🔴 MUST — 반드시 풀 줄 알아야 함 (8개)

S3~G4 문제 수 상위 + HackerRank 자료구조 출제 경향 반영.

| 순서 | 폴더 | S3~G4 문제 수 | 핵심 포인트 | 연습 문제 |
|:---:|---|---:|---|---|
| 1 | [01-implementation-and-simulation](./01-implementation-and-simulation/) | 1,607+500 | 문제를 코드로 옮기는 구현력. S3~G4의 절대 다수 | easy + medium |
| 2 | [02-bruteforce](./02-bruteforce/) | 1,133 | 완전탐색, 순열/조합 열거. 비율 41%로 이 범위 집중 | easy + medium |
| 3 | [11-greedy](./11-greedy/) | 957 | 정렬 후 탐욕 선택, 활동 선택, 스케줄링 | easy + medium |
| 4 | [19-dynamic-programming](./19-dynamic-programming/) | 811 | 1차원/2차원 DP, 메모이제이션 | easy + medium |
| 5 | [03-sorting](./03-sorting/) | 761 | `sort.Slice`, 커스텀 정렬, 정렬 후 탐색 | easy + medium |
| 6 | [16-graph-bfs](./16-graph-bfs/) | 468 | BFS 최단거리, 레벨 탐색, 격자 탐색 | easy + medium |
| 7 | [15-graph-dfs](./15-graph-dfs/) | 287 | 연결 요소, 사이클 탐지, 재귀 DFS | easy + medium |
| 8 | [05-hash](./05-hash/) | 445 (set) | map 활용, 빈도수 세기, 중복 체크, 그룹핑 | easy + medium |

### 🟡 HIGH — 높은 확률로 출제 (7개)

S3~G4 300개 전후 + HackerRank 자료구조 직접 출제 유형.

| 순서 | 폴더 | S3~G4 문제 수 | 핵심 포인트 | 연습 문제 |
|:---:|---|---:|---|---|
| 9 | [26-string-algorithm](./26-string-algorithm/) | 693 | 문자열 처리, 파싱, 패턴 매칭 기초 | easy + medium |
| 10 | [17-backtracking](./17-backtracking/) | 308 | 순열/조합 생성, 가지치기. **S3~G4 비율 47% 최고** | easy + medium |
| 11 | [06-prefix-sum](./06-prefix-sum/) | 317 | 구간 합, 부분 배열 합 = K | easy + medium |
| 12 | [08-binary-search](./08-binary-search/) | 298 | 이분 탐색, lower/upper bound, `sort.Search` | easy + medium |
| 13 | [04-stack-and-queue](./04-stack-and-queue/) | 139+30 | 괄호 매칭, 모노톤 스택. HackerRank 자료구조 단골 | easy + medium |
| 14 | [10-two-pointer-and-sliding-window](./10-two-pointer-and-sliding-window/) | 133+44 | 투 포인터, 슬라이딩 윈도우, 구간 최적화 | easy + medium |
| 15 | [12-heap-and-priority-queue](./12-heap-and-priority-queue/) | 113 | `container/heap`, K번째 원소. HackerRank 자료구조 단골 | easy + medium |

### 🟢 POSSIBLE — 나올 수 있음 (5개)

골드 경계 문제에서 등장할 수 있는 유형들.

| 순서 | 폴더 | S3~G4 문제 수 | 핵심 포인트 | 연습 문제 |
|:---:|---|---:|---|---|
| 16 | [21-shortest-path](./21-shortest-path/) | 152 | 다익스트라 기본 | easy |
| 17 | [13-tree](./13-tree/) | 120 | 트리 순회, 높이, 지름 | easy + medium |
| 18 | [09-parametric-search](./09-parametric-search/) | 111 | 최적화 → 결정 문제 변환 | easy |
| 19 | [20-union-find](./20-union-find/) | 60 | 연결 요소 관리, 그룹 병합 | easy |
| 20 | [59-flood-fill](./59-flood-fill/) | 57 | 영역 채우기 (BFS/DFS 응용). 비율 70% | easy |

---

## 7일 학습 플랜

> 하루 2~3시간 기준. theory.md → examples → problems 순서로 학습.
> **각 문제는 반드시 직접 풀어본 후 explanation.md 확인.**

### Day 1 (3/23 월) — 구현 + 브루트포스

S3~G4에서 가장 많이 나오는 두 유형. 여기서 기본기를 잡아야 함.

```
📂 01-implementation-and-simulation/
   theory.md → examples/ → problems/01-easy, 02-medium

📂 02-bruteforce/
   theory.md → examples/ → problems/01-easy, 02-medium
```

**체크리스트:**
- [ ] 행렬 회전, 나선형 순회 등 2D 배열 조작
- [ ] 시뮬레이션: 조건 분기를 정확히 코드로 옮기기
- [ ] 완전탐색: 중첩 루프, 재귀로 모든 경우 열거
- [ ] 시간복잡도 감 잡기: N ≤ 20이면 2^N, N ≤ 8이면 N! 가능

### Day 2 (3/24 화) — 그리디 + 정렬 + DP

문제 풀이의 3대 사고법.

```
📂 11-greedy/
   theory.md → examples/ → problems/01-easy, 02-medium

📂 03-sorting/
   theory.md → problems/01-easy, 02-medium

📂 19-dynamic-programming/
   theory.md → examples/ → problems/01-easy, 02-medium
```

**체크리스트:**
- [ ] 그리디 판단: 탐욕 선택 속성 + 최적 부분 구조
- [ ] `sort.Slice(arr, func(i, j int) bool { ... })` 패턴
- [ ] DP 접근법: 점화식 → 테이블 정의 → 초기값 → 순서

### Day 3 (3/25 수) — BFS/DFS + 해시

그래프 탐색과 자료구조 활용.

```
📂 16-graph-bfs/
   theory.md → examples/ → problems/01-easy, 02-medium

📂 15-graph-dfs/
   theory.md → examples/ → problems/01-easy, 02-medium

📂 05-hash/
   theory.md → examples/ → problems/01-easy, 02-medium
```

**체크리스트:**
- [ ] BFS: 큐 + visited + 레벨 카운팅 템플릿
- [ ] DFS: 재귀 vs 스택 기반, visited 처리
- [ ] 격자(grid) 탐색: `dx, dy` 4방향 패턴
- [ ] Go `map` 선언, 순회, 존재 확인 (`val, ok := m[key]`)

### Day 4 (3/26 목) — 문자열 + 백트래킹 + 누적합

S3~G4 비율이 높은 유형들.

```
📂 26-string-algorithm/
   theory.md → problems/01-easy, 02-medium

📂 17-backtracking/
   theory.md → examples/ → problems/01-easy, 02-medium

📂 06-prefix-sum/
   theory.md → problems/01-easy, 02-medium
```

**체크리스트:**
- [ ] 문자열: `strings` 패키지, 룬 처리, 부분 문자열
- [ ] 백트래킹: 선택 → 탐색 → 복원 패턴. 가지치기 조건
- [ ] 누적합: `prefix[i] = prefix[i-1] + arr[i]`, 구간합 = `prefix[r] - prefix[l-1]`

### Day 5 (3/27 금) — 이진탐색 + 스택/큐 + 투포인터

탐색 기법과 HackerRank 자료구조 대비.

```
📂 08-binary-search/
   theory.md → examples/ → problems/01-easy, 02-medium

📂 04-stack-and-queue/
   theory.md → examples/ → problems/01-easy, 02-medium

📂 10-two-pointer-and-sliding-window/
   theory.md → examples/ → problems/01-easy, 02-medium
```

**체크리스트:**
- [ ] `sort.Search(n, func(i int) bool { ... })` 활용법
- [ ] 스택: slice로 push/pop, 모노톤 스택 패턴
- [ ] 투 포인터: 양끝 수렴 vs 같은 방향 이동
- [ ] 슬라이딩 윈도우: 윈도우 확장/축소 조건

### Day 6 (3/28 토) — 힙 + POSSIBLE 유형 + 약점 보강

```
📂 12-heap-and-priority-queue/
   theory.md → examples/ → problems/01-easy, 02-medium

📂 21-shortest-path/
   theory.md → problems/01-easy

📂 13-tree/
   theory.md → problems/01-easy

📂 약점 폴더 medium 문제 재풀이
```

**체크리스트:**
- [ ] `container/heap` 인터페이스 구현 패턴 (Len, Less, Swap, Push, Pop)
- [ ] 다익스트라 = BFS + 힙 조합
- [ ] 트리 순회 3종 (전위/중위/후위)

### Day 7 (3/29 일) — 종합 복습 + 실전 시뮬레이션

```
오전: 핵심 패턴 코드 템플릿 총정리 (아래 치트시트 참고)
오후: 75분 타이머 맞추고 각 폴더 medium 문제 3~4개 연속 풀기
저녁: HackerRank 샘플 테스트 풀기 (초대 메일의 연습 링크)
```

---

## Go 코드 템플릿 치트시트

시험 중 빠르게 쓸 수 있도록 손에 익혀둘 패턴들.

### 1. 입출력 빠른 처리 (BOJ 스타일, HackerRank에선 불필요할 수 있음)

```go
reader := bufio.NewReader(os.Stdin)
writer := bufio.NewWriter(os.Stdout)
defer writer.Flush()
fmt.Fscan(reader, &n)
```

### 2. 해시맵 빈도수 세기

```go
freq := make(map[string]int)
for _, v := range arr {
    freq[v]++
}
```

### 3. 스택 (slice 기반)

```go
stack := []int{}
stack = append(stack, val)           // push
top := stack[len(stack)-1]           // peek
stack = stack[:len(stack)-1]         // pop
```

### 4. 힙 (container/heap)

```go
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
```

### 5. BFS 격자 탐색

```go
dx := []int{-1, 1, 0, 0}
dy := []int{0, 0, -1, 1}

queue := [][]int{{sr, sc}}
visited := make([][]bool, n)
for i := range visited {
    visited[i] = make([]bool, m)
}
visited[sr][sc] = true

for len(queue) > 0 {
    cur := queue[0]
    queue = queue[1:]
    for d := 0; d < 4; d++ {
        nx, ny := cur[0]+dx[d], cur[1]+dy[d]
        if nx >= 0 && nx < n && ny >= 0 && ny < m && !visited[nx][ny] {
            visited[nx][ny] = true
            queue = append(queue, []int{nx, ny})
        }
    }
}
```

### 6. DFS 재귀

```go
visited := make([]bool, n)
var dfs func(int)
dfs = func(v int) {
    visited[v] = true
    for _, u := range graph[v] {
        if !visited[u] {
            dfs(u)
        }
    }
}
```

### 7. 이진 탐색

```go
idx := sort.Search(len(arr), func(i int) bool {
    return arr[i] >= target
})
```

### 8. Union-Find

```go
parent := make([]int, n)
for i := range parent {
    parent[i] = i
}
var find func(int) int
find = func(x int) int {
    if parent[x] != x {
        parent[x] = find(parent[x])
    }
    return parent[x]
}
union := func(a, b int) {
    parent[find(a)] = find(b)
}
```

### 9. 커스텀 정렬

```go
sort.Slice(arr, func(i, j int) bool {
    if arr[i].score == arr[j].score {
        return arr[i].name < arr[j].name
    }
    return arr[i].score > arr[j].score
})
```

### 10. DP 기본 패턴

```go
dp := make([]int, n+1)
dp[0] = baseCase
for i := 1; i <= n; i++ {
    dp[i] = max(dp[i-1]+val[i], dp[i-2]+other[i])
}
```

### 11. 백트래킹 템플릿

```go
var result [][]int
var bt func(start int, path []int)
bt = func(start int, path []int) {
    if len(path) == k {
        tmp := make([]int, k)
        copy(tmp, path)
        result = append(result, tmp)
        return
    }
    for i := start; i < n; i++ {
        bt(i+1, append(path, arr[i]))
    }
}
bt(0, nil)
```

### 12. 문자열 처리

```go
// 룬 단위 순회
for i, r := range s {
    // r은 rune, i는 바이트 인덱스
}
// 문자열 → []byte 변환 후 조작
b := []byte(s)
b[i] = 'x'
s = string(b)
```

---

## 시험 당일 전략 (75분)

### 시간 배분
| 구간 | 시간 | 할 일 |
|---|---|---|
| 0~5분 | 5분 | 전체 문제 훑기, 난이도 파악, 풀이 순서 결정 |
| 5~25분 | 20분 | Easy 문제 확실하게 풀기 |
| 25~55분 | 30분 | Medium 문제 2개 도전 |
| 55~70분 | 15분 | Hard 문제 시도 또는 Medium 마무리 |
| 70~75분 | 5분 | 엣지 케이스 점검, 제출 확인 |

### 제약 조건으로 풀이법 추론하기
| 제약 조건 | 가능한 복잡도 | 풀이 방향 |
|---|---|---|
| N ≤ 20 | O(2^N) | 브루트포스, 비트마스크 |
| N ≤ 10 | O(N!) | 순열 전탐색 |
| N ≤ 100 | O(N³) | 3중 루프, 플로이드 |
| N ≤ 10,000 | O(N²) | 2중 루프, DP |
| N ≤ 1,000,000 | O(N log N) | 정렬, 이분탐색, 힙 |
| N ≤ 10,000,000 | O(N) | 투 포인터, 해시, 누적합 |

### 핵심 원칙
1. **Easy를 절대 놓치지 마라** — 확실한 점수 먼저 확보
2. **문제를 끝까지 읽어라** — 제약 조건에 풀이 힌트가 있음
3. **막히면 넘겨라** — 75분은 짧다. 20분 이상 한 문제에 쓰지 마라
4. **부분 점수를 노려라** — 최적해를 못 찾으면 브루트포스라도 제출

### HackerRank 특이사항
- 함수 시그니처가 주어짐 → 입출력 파싱 불필요
- 테스트 케이스가 숨겨져 있음 → 엣지 케이스 직접 생각해야 함
- 코드 실행 버튼으로 중간 테스트 가능 → 적극 활용
- 브라우저 새로고침/뒤로가기 금지

### 자주 놓치는 엣지 케이스
- 빈 배열/문자열 입력
- 원소가 1개인 경우
- 모든 원소가 같은 경우
- 음수 포함 여부
- int 오버플로우 (Go는 int가 64비트이므로 대부분 안전)
- 0-indexed vs 1-indexed

---

## 학습 진행 체크리스트

### MUST (Day 1~3)
- [ ] 01-implementation-and-simulation: theory + easy + medium
- [ ] 02-bruteforce: theory + easy + medium
- [ ] 11-greedy: theory + easy + medium
- [ ] 19-dynamic-programming: theory + easy + medium
- [ ] 03-sorting: theory + easy + medium
- [ ] 16-graph-bfs: theory + easy + medium
- [ ] 15-graph-dfs: theory + easy + medium
- [ ] 05-hash: theory + easy + medium

### HIGH (Day 4~5)
- [ ] 26-string-algorithm: theory + easy + medium
- [ ] 17-backtracking: theory + easy + medium
- [ ] 06-prefix-sum: theory + easy + medium
- [ ] 08-binary-search: theory + easy + medium
- [ ] 04-stack-and-queue: theory + easy + medium
- [ ] 10-two-pointer-and-sliding-window: theory + easy + medium
- [ ] 12-heap-and-priority-queue: theory + easy + medium

### POSSIBLE (Day 6)
- [ ] 21-shortest-path: theory + easy
- [ ] 13-tree: theory + easy + medium
- [ ] 09-parametric-search: theory + easy
- [ ] 20-union-find: theory + easy
- [ ] 59-flood-fill: theory + easy
