# Day 2: 그리디 + 정렬 + DP

> S3~G4 문제 수: 그리디 957개 + 정렬 761개 + DP 811개 (3~5위)
> 예상 소요: 2~3시간

문제 풀이의 3대 사고법. "이 문제를 어떤 접근법으로 풀 것인가"를 결정하는 핵심.

---

## 1. 그리디

> 📂 [11-greedy](../../11-greedy/) → theory.md → examples → problems

### 핵심 개념
- 매 단계에서 **현재 최선의 선택**을 반복 → 전체 최적해
- 적용 조건: 탐욕 선택 속성 + 최적 부분 구조
- 대부분 **정렬이 선행**됨 → 정렬 후 순회하며 선택

### 그리디 vs DP 판단법
- 그리디: 한 번 선택하면 되돌아가지 않음. "지금 최선 = 전체 최선"
- DP: 모든 선택지를 고려해야 함. "지금 최선 ≠ 전체 최선"일 수 있음
- 확신이 없으면 DP로 풀어라 (그리디는 증명이 필요)

### 시험에 나오는 패턴

**패턴 1: 정렬 후 탐욕 선택 (활동 선택)**
```go
sort.Slice(activities, func(i, j int) bool {
    return activities[i].end < activities[j].end // 종료 시간 기준 정렬
})
count := 1
lastEnd := activities[0].end
for i := 1; i < len(activities); i++ {
    if activities[i].start >= lastEnd {
        count++
        lastEnd = activities[i].end
    }
}
```

**패턴 2: 거스름돈 (큰 단위부터)**
```go
coins := []int{500, 100, 50, 10}
count := 0
for _, coin := range coins {
    count += amount / coin
    amount %= coin
}
```

### 연습 문제
- [ ] [01-easy-coin-change-greedy](../../11-greedy/problems/01-easy-coin-change-greedy/)
- [ ] [02-medium-activity-selection](../../11-greedy/problems/02-medium-activity-selection/)

---

## 2. 정렬

> 📂 [03-sorting](../../03-sorting/) → theory.md → examples → problems

### 핵심 개념
- Go의 `sort.Slice`가 핵심. 직접 정렬 알고리즘 구현할 일 없음
- **커스텀 정렬** (다중 조건)이 시험에 자주 나옴
- 정렬은 그리디, 이진탐색, 투포인터의 **전처리 단계**

### 시험에 나오는 패턴

**패턴 1: 기본 정렬**
```go
sort.Ints(arr)                    // 오름차순
sort.Sort(sort.Reverse(sort.IntSlice(arr))) // 내림차순
```

**패턴 2: 커스텀 정렬 (다중 조건)**
```go
sort.Slice(items, func(i, j int) bool {
    if items[i].score == items[j].score {
        return items[i].name < items[j].name // 2차: 이름 오름차순
    }
    return items[i].score > items[j].score   // 1차: 점수 내림차순
})
```

**패턴 3: 안정 정렬**
```go
sort.SliceStable(arr, func(i, j int) bool {
    return arr[i].priority < arr[j].priority
})
```

### 연습 문제
- [ ] [01-easy-sort-numbers](../../03-sorting/problems/01-easy-sort-numbers/)
- [ ] [02-medium-custom-sort](../../03-sorting/problems/02-medium-custom-sort/)

---

## 3. 동적 프로그래밍 (DP)

> 📂 [19-dynamic-programming](../../19-dynamic-programming/) → theory.md → examples → problems

### 핵심 개념
- 큰 문제를 작은 하위 문제로 분해, 결과를 **저장하여 재활용**
- 적용 조건: 최적 부분 구조 + 중복 부분 문제
- 핵심 과정: **상태 정의 → 점화식 → 초기값 → 순서**

### DP 문제 풀이 5단계 (암기)
1. **상태 정의**: `dp[i]`가 무엇을 의미하는지 정한다
2. **점화식**: `dp[i]`를 이전 상태로 표현한다
3. **초기값**: `dp[0]`, `dp[1]` 등 기저 사례를 설정한다
4. **순서**: 작은 문제부터 큰 문제 순서로 채운다
5. **답 추출**: `dp[n]` 또는 `max(dp[...])` 등

### 시험에 나오는 패턴

**패턴 1: 1차원 DP (피보나치, 계단 오르기)**
```go
dp := make([]int, n+1)
dp[1] = 1
dp[2] = 2
for i := 3; i <= n; i++ {
    dp[i] = dp[i-1] + dp[i-2]
}
```

**패턴 2: 2차원 DP (배낭 문제)**
```go
dp := make([][]int, n+1)
for i := range dp {
    dp[i] = make([]int, W+1)
}
for i := 1; i <= n; i++ {
    for w := 0; w <= W; w++ {
        dp[i][w] = dp[i-1][w] // 안 넣는 경우
        if w >= weight[i] {
            dp[i][w] = max(dp[i][w], dp[i-1][w-weight[i]]+value[i])
        }
    }
}
```

**패턴 3: 문자열 DP (LCS)**
```go
dp := make([][]int, n+1)
for i := range dp {
    dp[i] = make([]int, m+1)
}
for i := 1; i <= n; i++ {
    for j := 1; j <= m; j++ {
        if a[i-1] == b[j-1] {
            dp[i][j] = dp[i-1][j-1] + 1
        } else {
            dp[i][j] = max(dp[i-1][j], dp[i][j-1])
        }
    }
}
```

### 연습 문제
- [ ] [01-easy-fibonacci](../../19-dynamic-programming/problems/01-easy-fibonacci/)
- [ ] [02-medium-knapsack](../../19-dynamic-programming/problems/02-medium-knapsack/)

---

## Day 2 체크리스트

- [ ] 그리디 theory.md 읽기
- [ ] 그리디 vs DP 판단법 이해
- [ ] 활동 선택 패턴 직접 타이핑
- [ ] 그리디 easy + medium 풀기
- [ ] `sort.Slice` 커스텀 정렬 패턴 암기
- [ ] 정렬 easy + medium 풀기
- [ ] DP 5단계 풀이법 암기
- [ ] 1차원 DP, 2차원 DP 패턴 직접 타이핑
- [ ] DP easy + medium 풀기
