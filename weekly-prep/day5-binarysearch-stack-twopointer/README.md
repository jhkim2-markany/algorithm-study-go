# Day 5: 이진탐색 + 스택/큐 + 투포인터

> S3~G4 문제 수: 이진탐색 298개 + 스택 139개 + 투포인터 133개 + 슬라이딩윈도우 44개
> 예상 소요: 2~3시간

탐색 최적화 기법들. HackerRank에서 스택은 자료구조 문제로 직접 출제된다.

---

## 1. 이진 탐색

> 📂 [08-binary-search](../../08-binary-search/) → theory.md → examples → problems

### 핵심 개념
- 정렬된 배열에서 O(log N) 탐색
- **전제 조건: 데이터가 정렬되어 있어야 함**
- Go의 `sort.Search`가 lower bound 역할

### 시험에 나오는 패턴

**패턴 1: sort.Search (lower bound)**
```go
// arr에서 target 이상인 첫 번째 인덱스
idx := sort.Search(len(arr), func(i int) bool {
    return arr[i] >= target
})
// target 존재 확인
if idx < len(arr) && arr[idx] == target {
    // 존재
}
```

**패턴 2: 직접 구현 이진 탐색**
```go
lo, hi := 0, len(arr)-1
for lo <= hi {
    mid := (lo + hi) / 2
    if arr[mid] == target {
        return mid
    } else if arr[mid] < target {
        lo = mid + 1
    } else {
        hi = mid - 1
    }
}
return -1 // not found
```

**패턴 3: 값의 개수 세기 (upper - lower)**
```go
lower := sort.Search(len(arr), func(i int) bool { return arr[i] >= target })
upper := sort.Search(len(arr), func(i int) bool { return arr[i] > target })
count := upper - lower
```

### 연습 문제
- [ ] [01-easy-find-target](../../08-binary-search/problems/01-easy-find-target/)
- [ ] [02-medium-lower-upper-bound](../../08-binary-search/problems/02-medium-lower-upper-bound/)

---

## 2. 스택과 큐

> 📂 [04-stack-and-queue](../../04-stack-and-queue/) → theory.md → examples → problems

### 핵심 개념
- 스택: LIFO. 괄호 매칭, 모노톤 스택이 핵심
- 큐: FIFO. BFS에서 이미 사용함
- HackerRank "Data Structures" 카테고리 단골

### 시험에 나오는 패턴

**패턴 1: 괄호 매칭**
```go
func isValid(s string) bool {
    stack := []byte{}
    match := map[byte]byte{')': '(', ']': '[', '}': '{'}
    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == '[' || s[i] == '{' {
            stack = append(stack, s[i])
        } else {
            if len(stack) == 0 || stack[len(stack)-1] != match[s[i]] {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}
```

**패턴 2: 모노톤 스택 (다음 큰 원소)**
```go
// 각 원소의 오른쪽에서 처음으로 자신보다 큰 원소의 인덱스
result := make([]int, n)
for i := range result {
    result[i] = -1
}
stack := []int{} // 인덱스 저장
for i := 0; i < n; i++ {
    for len(stack) > 0 && arr[stack[len(stack)-1]] < arr[i] {
        top := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result[top] = i
    }
    stack = append(stack, i)
}
```

**패턴 3: 스택으로 계산기**
```go
// 후위 표기법 계산
stack := []int{}
for _, token := range tokens {
    switch token {
    case "+", "-", "*", "/":
        b := stack[len(stack)-1]; stack = stack[:len(stack)-1]
        a := stack[len(stack)-1]; stack = stack[:len(stack)-1]
        switch token {
        case "+": stack = append(stack, a+b)
        case "-": stack = append(stack, a-b)
        case "*": stack = append(stack, a*b)
        case "/": stack = append(stack, a/b)
        }
    default:
        num, _ := strconv.Atoi(token)
        stack = append(stack, num)
    }
}
```

### 연습 문제
- [ ] [01-easy-valid-parentheses](../../04-stack-and-queue/problems/01-easy-valid-parentheses/)
- [ ] [02-medium-stock-price](../../04-stack-and-queue/problems/02-medium-stock-price/)

---

## 3. 투 포인터 + 슬라이딩 윈도우

> 📂 [10-two-pointer-and-sliding-window](../../10-two-pointer-and-sliding-window/) → theory.md → examples → problems

### 핵심 개념
- O(N²) → O(N) 최적화의 핵심 기법
- 투 포인터: 양끝 수렴 or 같은 방향 이동
- 슬라이딩 윈도우: 구간을 유지하며 확장/축소

### 시험에 나오는 패턴

**패턴 1: 양끝 수렴 (정렬된 배열에서 Two Sum)**
```go
sort.Ints(arr)
left, right := 0, len(arr)-1
for left < right {
    sum := arr[left] + arr[right]
    if sum == target {
        return true
    } else if sum < target {
        left++
    } else {
        right--
    }
}
```

**패턴 2: 슬라이딩 윈도우 (중복 없는 최장 부분 문자열)**
```go
seen := make(map[byte]int) // 문자 → 마지막 인덱스
maxLen := 0
left := 0
for right := 0; right < len(s); right++ {
    if idx, ok := seen[s[right]]; ok && idx >= left {
        left = idx + 1
    }
    seen[s[right]] = right
    if right-left+1 > maxLen {
        maxLen = right - left + 1
    }
}
```

**패턴 3: 합이 K 이상인 최소 길이 부분 배열**
```go
minLen := n + 1
sum := 0
left := 0
for right := 0; right < n; right++ {
    sum += arr[right]
    for sum >= k {
        if right-left+1 < minLen {
            minLen = right - left + 1
        }
        sum -= arr[left]
        left++
    }
}
```

### 연습 문제
- [ ] [01-easy-two-sum-sorted](../../10-two-pointer-and-sliding-window/problems/01-easy-two-sum-sorted/)
- [ ] [02-medium-longest-unique-substring](../../10-two-pointer-and-sliding-window/problems/02-medium-longest-unique-substring/)

---

## Day 5 체크리스트

- [ ] `sort.Search` 사용법 암기 (lower bound)
- [ ] 직접 구현 이진 탐색 패턴 타이핑
- [ ] 이진탐색 easy + medium 풀기
- [ ] 괄호 매칭 패턴 직접 타이핑
- [ ] 모노톤 스택 패턴 이해 + 타이핑
- [ ] 스택/큐 easy + medium 풀기
- [ ] 양끝 수렴 투포인터 패턴 타이핑
- [ ] 슬라이딩 윈도우 패턴 타이핑
- [ ] 투포인터 easy + medium 풀기
