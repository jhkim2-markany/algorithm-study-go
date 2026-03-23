# Day 3: BFS/DFS + 해시

> S3~G4 문제 수: 그래프탐색 727개 + BFS 468개 + DFS 287개 + 해시셋 445개
> 예상 소요: 2~3시간

그래프 탐색은 S3~G4의 핵심. 해시는 HackerRank 자료구조 단골.

---

## 1. BFS (너비 우선 탐색)

> 📂 [16-graph-bfs](../../16-graph-bfs/) → theory.md → examples → problems

### 핵심 개념
- **큐** 기반, 가까운 정점부터 탐색
- 가중치 없는 그래프에서 **최단 거리 보장**
- 격자(grid) 탐색 문제에서 가장 많이 사용

### 시험에 나오는 패턴

**패턴 1: 격자 BFS (최단거리)**
```go
type Point struct{ x, y, dist int }

dx := []int{-1, 1, 0, 0}
dy := []int{0, 0, -1, 1}

queue := []Point{{sr, sc, 0}}
visited := make([][]bool, n)
for i := range visited {
    visited[i] = make([]bool, m)
}
visited[sr][sc] = true

for len(queue) > 0 {
    cur := queue[0]
    queue = queue[1:]

    if cur.x == er && cur.y == ec {
        return cur.dist // 최단 거리
    }

    for d := 0; d < 4; d++ {
        nx, ny := cur.x+dx[d], cur.y+dy[d]
        if nx >= 0 && nx < n && ny >= 0 && ny < m && !visited[nx][ny] && grid[nx][ny] != '#' {
            visited[nx][ny] = true
            queue = append(queue, Point{nx, ny, cur.dist + 1})
        }
    }
}
```

**패턴 2: 인접 리스트 BFS (연결 요소)**
```go
visited := make([]bool, n+1)
components := 0
for v := 1; v <= n; v++ {
    if !visited[v] {
        components++
        queue := []int{v}
        visited[v] = true
        for len(queue) > 0 {
            cur := queue[0]
            queue = queue[1:]
            for _, u := range graph[cur] {
                if !visited[u] {
                    visited[u] = true
                    queue = append(queue, u)
                }
            }
        }
    }
}
```

### 연습 문제
- [ ] [01-easy-maze-shortest-path](../../16-graph-bfs/problems/01-easy-maze-shortest-path/)
- [ ] [02-medium-tomato-ripening](../../16-graph-bfs/problems/02-medium-tomato-ripening/)

---

## 2. DFS (깊이 우선 탐색)

> 📂 [15-graph-dfs](../../15-graph-dfs/) → theory.md → examples → problems

### 핵심 개념
- **재귀/스택** 기반, 한 경로를 끝까지 탐색 후 되돌아감
- 연결 요소, 사이클 탐지, 경로 존재 여부에 사용
- BFS와 달리 최단 거리를 보장하지 않음

### BFS vs DFS 선택 기준
| 상황 | 선택 |
|---|---|
| 최단 거리/최소 이동 | **BFS** |
| 연결 요소 개수 | BFS 또는 DFS |
| 사이클 탐지 | **DFS** |
| 모든 경로 탐색 | **DFS** |
| 격자 영역 탐색 | BFS 또는 DFS |

### 시험에 나오는 패턴

**패턴 1: 재귀 DFS**
```go
visited := make([]bool, n+1)
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

**패턴 2: 사이클 탐지 (방향 그래프)**
```go
// 0: 미방문, 1: 탐색 중, 2: 탐색 완료
state := make([]int, n+1)
hasCycle := false

var dfs func(int)
dfs = func(v int) {
    state[v] = 1
    for _, u := range graph[v] {
        if state[u] == 1 {
            hasCycle = true
            return
        }
        if state[u] == 0 {
            dfs(u)
        }
    }
    state[v] = 2
}
```

### 연습 문제
- [ ] [01-easy-connected-components](../../15-graph-dfs/problems/01-easy-connected-components/)
- [ ] [02-medium-cycle-detection](../../15-graph-dfs/problems/02-medium-cycle-detection/)

---

## 3. 해시

> 📂 [05-hash](../../05-hash/) → theory.md → examples → problems

### 핵심 개념
- Go의 `map`이 해시맵. 평균 O(1) 삽입/조회/삭제
- **빈도수 세기, 중복 체크, 그룹핑**이 3대 활용법
- HackerRank에서 자료구조 문제로 직접 출제됨

### 시험에 나오는 패턴

**패턴 1: 빈도수 세기**
```go
freq := make(map[string]int)
for _, v := range arr {
    freq[v]++
}
// 최빈값 찾기
maxCount := 0
for _, count := range freq {
    if count > maxCount {
        maxCount = count
    }
}
```

**패턴 2: Two Sum (해시로 O(N))**
```go
seen := make(map[int]int) // 값 → 인덱스
for i, v := range arr {
    if j, ok := seen[target-v]; ok {
        return []int{j, i}
    }
    seen[v] = i
}
```

**패턴 3: 아나그램 그룹핑**
```go
groups := make(map[string][]string)
for _, word := range words {
    b := []byte(word)
    sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
    key := string(b)
    groups[key] = append(groups[key], word)
}
```

**패턴 4: 존재 확인 (Set)**
```go
set := make(map[int]bool)
for _, v := range arr {
    set[v] = true
}
if set[target] {
    // 존재
}
```

### 연습 문제
- [ ] [01-easy-count-frequency](../../05-hash/problems/01-easy-count-frequency/)
- [ ] [02-medium-longest-substring](../../05-hash/problems/02-medium-longest-substring/)

---

## Day 3 체크리스트

- [ ] BFS theory.md 읽기
- [ ] 격자 BFS 템플릿 직접 타이핑 (dx, dy, visited, queue)
- [ ] BFS easy + medium 풀기
- [ ] DFS theory.md 읽기
- [ ] 재귀 DFS 템플릿 직접 타이핑
- [ ] BFS vs DFS 선택 기준 암기
- [ ] DFS easy + medium 풀기
- [ ] 해시 theory.md 읽기
- [ ] `map` 패턴 4가지 직접 타이핑
- [ ] 해시 easy + medium 풀기
