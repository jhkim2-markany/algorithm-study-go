# Day 4: 문자열 + 백트래킹 + 누적합

> S3~G4 문제 수: 문자열 693개 + 백트래킹 308개(비율 47%) + 누적합 317개
> 예상 소요: 2~3시간

문자열은 의외로 S3~G4에서 7위. 백트래킹은 이 범위 집중도가 가장 높다.

---

## 1. 문자열

> 📂 [26-string-algorithm](../../26-string-algorithm/) → theory.md → examples → problems

### 핵심 개념
- S3~G4에서는 KMP보다 **기본 문자열 처리**가 중요
- Go 문자열은 **불변(immutable)**, 수정하려면 `[]byte`로 변환
- `strings` 패키지 함수들을 잘 알아두면 구현 시간 단축

### Go 문자열 필수 지식

```go
// 기본 조작
len(s)                          // 바이트 길이
s[i]                            // i번째 바이트 (byte)
s[l:r]                          // 부분 문자열

// 룬(유니코드) 순회
for i, r := range s {           // r은 rune, i는 바이트 인덱스
}
runes := []rune(s)              // 룬 슬라이스로 변환

// 수정이 필요할 때
b := []byte(s)
b[i] = 'x'
s = string(b)

// strings 패키지
strings.Contains(s, "abc")      // 포함 여부
strings.Index(s, "abc")         // 위치 (-1 if not found)
strings.Split(s, ",")           // 분리
strings.Join(arr, ",")          // 합치기
strings.ToLower(s)              // 소문자
strings.Replace(s, "a", "b", -1) // 치환
```

### 시험에 나오는 패턴

**패턴 1: 문자 빈도수 (배열 활용)**
```go
var count [26]int
for _, c := range s {
    count[c-'a']++
}
```

**패턴 2: 회문(팰린드롬) 판별**
```go
func isPalindrome(s string) bool {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        if s[i] != s[j] {
            return false
        }
    }
    return true
}
```

**패턴 3: 브루트포스 패턴 매칭**
```go
for i := 0; i <= len(text)-len(pattern); i++ {
    if text[i:i+len(pattern)] == pattern {
        // 매칭 위치: i
    }
}
```

### 연습 문제
- [ ] [01-easy-pattern-matching](../../26-string-algorithm/problems/01-easy-pattern-matching/)
- [ ] [02-medium-kmp-search](../../26-string-algorithm/problems/02-medium-kmp-search/)

---

## 2. 백트래킹

> 📂 [17-backtracking](../../17-backtracking/) → theory.md → examples → problems

### 핵심 개념
- 브루트포스 + **가지치기(pruning)** = 백트래킹
- 순열, 조합, 부분집합 생성의 표준 방법
- S3~G4 비율 **47%**로 이 범위에 가장 집중된 유형

### 시험에 나오는 패턴

**패턴 1: 조합 (nCk)**
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

**패턴 2: 순열 (nPk)**
```go
used := make([]bool, n)
var bt func(path []int)
bt = func(path []int) {
    if len(path) == k {
        tmp := make([]int, k)
        copy(tmp, path)
        result = append(result, tmp)
        return
    }
    for i := 0; i < n; i++ {
        if !used[i] {
            used[i] = true
            bt(append(path, arr[i]))
            used[i] = false // 복원!
        }
    }
}
bt(nil)
```

**패턴 3: 가지치기 (N-Queen 스타일)**
```go
var bt func(row int)
bt = func(row int) {
    if row == n {
        count++
        return
    }
    for col := 0; col < n; col++ {
        if isValid(row, col) { // 가지치기
            place(row, col)
            bt(row + 1)
            remove(row, col) // 복원
        }
    }
}
```

### 연습 문제
- [ ] [01-easy-n-and-m](../../17-backtracking/problems/01-easy-n-and-m/)
- [ ] [02-medium-n-queen](../../17-backtracking/problems/02-medium-n-queen/)

---

## 3. 누적합

> 📂 [06-prefix-sum](../../06-prefix-sum/) → theory.md → examples → problems

### 핵심 개념
- 구간 합을 O(1)에 구하기 위한 전처리 기법
- `prefix[i] = prefix[i-1] + arr[i-1]`
- 구간 `[l, r]` 합 = `prefix[r+1] - prefix[l]`

### 시험에 나오는 패턴

**패턴 1: 1차원 누적합**
```go
prefix := make([]int, n+1)
for i := 1; i <= n; i++ {
    prefix[i] = prefix[i-1] + arr[i-1]
}
// 구간 [l, r] 합 (0-indexed)
sum := prefix[r+1] - prefix[l]
```

**패턴 2: 부분 배열 합 = K (누적합 + 해시)**
```go
count := 0
prefixSum := 0
seen := map[int]int{0: 1}
for _, v := range arr {
    prefixSum += v
    if c, ok := seen[prefixSum-k]; ok {
        count += c
    }
    seen[prefixSum]++
}
```

**패턴 3: 2차원 누적합**
```go
// 전처리
p := make([][]int, n+1)
for i := range p {
    p[i] = make([]int, m+1)
}
for i := 1; i <= n; i++ {
    for j := 1; j <= m; j++ {
        p[i][j] = grid[i-1][j-1] + p[i-1][j] + p[i][j-1] - p[i-1][j-1]
    }
}
// (r1,c1) ~ (r2,c2) 합 (0-indexed)
sum := p[r2+1][c2+1] - p[r1][c2+1] - p[r2+1][c1] + p[r1][c1]
```

### 연습 문제
- [ ] [01-easy-range-sum](../../06-prefix-sum/problems/01-easy-range-sum/)
- [ ] [02-medium-subarray-sum-equals-k](../../06-prefix-sum/problems/02-medium-subarray-sum-equals-k/)

---

## Day 4 체크리스트

- [ ] Go 문자열 기본 조작 + `strings` 패키지 함수 정리
- [ ] 문자 빈도수 배열 패턴 직접 타이핑
- [ ] 문자열 easy + medium 풀기
- [ ] 백트래킹 theory.md 읽기
- [ ] 조합/순열 템플릿 직접 타이핑 (copy 잊지 마라!)
- [ ] 백트래킹 easy + medium 풀기
- [ ] 누적합 공식 암기: `prefix[r+1] - prefix[l]`
- [ ] 누적합 + 해시 패턴 직접 타이핑
- [ ] 누적합 easy + medium 풀기
