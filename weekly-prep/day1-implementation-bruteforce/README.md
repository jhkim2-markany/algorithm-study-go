# Day 1: 구현 + 브루트포스

> S3~G4 문제 수: 구현 1,607개 + 브루트포스 1,133개 (1~2위)
> 예상 소요: 2~3시간

이 범위에서 가장 많이 나오는 유형. 특별한 알고리즘 없이 **문제를 정확히 코드로 옮기는 능력**을 테스트한다.

---

## 1. 구현과 시뮬레이션

> 📂 [01-implementation-and-simulation](../../01-implementation-and-simulation/) → theory.md → examples → problems

### 핵심 개념
- 문제의 조건과 절차를 **그대로** 코드로 표현
- 배열 조작, 좌표 이동, 상태 관리가 핵심
- 실수 없이 꼼꼼하게 구현하는 것이 관건

### 시험에 나오는 패턴

**패턴 1: 방향 배열 (격자 이동)**
```go
// 상하좌우
dx := []int{-1, 1, 0, 0}
dy := []int{0, 0, -1, 1}

for d := 0; d < 4; d++ {
    nx, ny := x+dx[d], y+dy[d]
    if nx >= 0 && nx < n && ny >= 0 && ny < m {
        // 유효한 좌표
    }
}
```

**패턴 2: 행렬 90도 시계 회전**
```go
func rotate90(matrix [][]int) [][]int {
    n := len(matrix)
    result := make([][]int, n)
    for i := range result {
        result[i] = make([]int, n)
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            result[j][n-1-i] = matrix[i][j]
        }
    }
    return result
}
```

**패턴 3: 나선형 순회**
```go
// 방향: 오른쪽 → 아래 → 왼쪽 → 위
dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
d := 0
r, c := 0, 0
for i := 0; i < n*m; i++ {
    result = append(result, matrix[r][c])
    visited[r][c] = true
    nr, nc := r+dirs[d][0], c+dirs[d][1]
    if nr < 0 || nr >= n || nc < 0 || nc >= m || visited[nr][nc] {
        d = (d + 1) % 4 // 방향 전환
        nr, nc = r+dirs[d][0], c+dirs[d][1]
    }
    r, c = nr, nc
}
```

### 연습 문제
- [ ] [01-easy-matrix-rotation](../../01-implementation-and-simulation/problems/01-easy-matrix-rotation/) — 행렬 회전
- [ ] [02-medium-spiral-matrix](../../01-implementation-and-simulation/problems/02-medium-spiral-matrix/) — 나선형 순회

---

## 2. 브루트포스

> 📂 [02-bruteforce](../../02-bruteforce/) → theory.md → examples → problems

### 핵심 개념
- 가능한 **모든 경우의 수**를 탐색하여 정답을 찾는 방법
- 입력 크기가 작으면 무조건 브루트포스부터 고려
- 시간 내에 가능한지 **경우의 수를 먼저 계산**

### 적용 판단 기준 (외워라)

| N 범위 | 가능한 복잡도 | 방법 |
|---|---|---|
| N ≤ 10 | O(N!) | 순열 전탐색 |
| N ≤ 20 | O(2^N) | 부분집합, 비트마스크 |
| N ≤ 100 | O(N³) | 3중 루프 |
| N ≤ 10,000 | O(N²) | 2중 루프 |

### 시험에 나오는 패턴

**패턴 1: 2중 루프 (모든 쌍)**
```go
for i := 0; i < n; i++ {
    for j := i + 1; j < n; j++ {
        if arr[i]+arr[j] == target {
            // 정답
        }
    }
}
```

**패턴 2: 비트마스크 부분집합**
```go
for mask := 0; mask < (1 << n); mask++ {
    sum := 0
    for i := 0; i < n; i++ {
        if mask&(1<<i) != 0 {
            sum += arr[i]
        }
    }
    if sum == target {
        // 정답
    }
}
```

**패턴 3: 재귀 선택/비선택**
```go
var solve func(idx, current int)
solve = func(idx, current int) {
    if idx == n {
        if current == target {
            count++
        }
        return
    }
    solve(idx+1, current+arr[idx]) // 선택
    solve(idx+1, current)           // 비선택
}
solve(0, 0)
```

### 연습 문제
- [ ] [01-easy-sum-of-pairs](../../02-bruteforce/problems/01-easy-sum-of-pairs/) — 쌍의 합
- [ ] [02-medium-permutation-check](../../02-bruteforce/problems/02-medium-permutation-check/) — 순열 확인

---

## Day 1 체크리스트

- [ ] 구현 theory.md 읽기
- [ ] 방향 배열 (dx, dy) 패턴 직접 타이핑
- [ ] 행렬 회전 공식 암기: `result[j][n-1-i] = matrix[i][j]`
- [ ] 구현 easy + medium 문제 풀기
- [ ] 브루트포스 theory.md 읽기
- [ ] 비트마스크 부분집합 패턴 직접 타이핑
- [ ] N 범위별 복잡도 판단 기준 암기
- [ ] 브루트포스 easy + medium 문제 풀기
