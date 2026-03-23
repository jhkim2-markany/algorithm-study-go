# 설계 문서: 코딩테스트 학습용 알고리즘 자료

## 개요 (Overview)

본 프로젝트는 코딩테스트 대비를 위한 알고리즘 학습 자료를 체계적으로 제공하는 정적 콘텐츠 저장소이다. 총 67개 알고리즘 유형을 학습 난이도 순으로 정리하며, 각 유형별로 이론 문서, Golang 예시 코드, 난이도별 문제/풀이/해설을 포함한다.

기존 26개 핵심 알고리즘 유형(01~26번)에 더해, 코딩테스트 빈출 및 실력 향상에 필요한 27개 확장 알고리즘 유형(27~53번), 코테 실전 필수 6개 Tier_1 유형(54~59번), 코테 상위권 차별화 8개 Tier_2 유형(60~67번)을 추가하여 총 67개 유형으로 구성한다. 확장 알고리즘은 BOJ 문제 수 기준 우선순위_1(1,000+), 우선순위_2(300~999), 우선순위_3(150~299)으로 분류되며, Tier_1은 한국 코딩테스트(삼성 SW, 카카오, 프로그래머스, BOJ 골드)에서 직접 출제되는 유형, Tier_2는 BOJ 골드 상위~플래티넘에서 등장하는 상위권 차별화 유형이다. 모든 확장 알고리즘은 기존 폴더와의 선수 학습 관계를 명시한다.

이 프로젝트는 런타임 애플리케이션이 아닌 파일 시스템 기반의 학습 콘텐츠 저장소이다. 핵심 설계 결정은 디렉토리 구조, 파일 명명 규칙, 문서 템플릿 형식에 집중된다.

## 아키텍처 (Architecture)

### 전체 구조

프로젝트는 플랫한 디렉토리 구조를 채택한다. 루트에 67개 알고리즘 폴더가 번호순으로 나열되며, 각 폴더 내부는 동일한 구조를 따른다. 01~26번은 기존 핵심 알고리즘, 27~53번은 확장 알고리즘, 54~59번은 Tier_1(코테 실전 필수), 60~67번은 Tier_2(코테 상위권 차별화)이다.

```
project-root/
├── README.md                          # 프로젝트 안내 문서
├── todo.md                            # 추후 추가할 알고리즘 목록
├── validate.sh                        # 검증 스크립트
├── 01-implementation-and-simulation/  # 기존 알고리즘 유형 폴더 (01~26)
│   ├── README.md                      # 폴더 안내
│   ├── theory.md                      # 이론 문서
│   ├── examples/                      # 예시 코드
│   │   └── simulation.go
│   └── problems/                      # 문제/풀이/해설
│       ├── 01-easy-problem-name/
│       │   ├── problem.md
│       │   ├── solution.go
│       │   └── explanation.md
│       ├── 02-medium-problem-name/
│       │   ├── problem.md
│       │   ├── solution.go
│       │   └── explanation.md
│       └── 03-hard-problem-name/
│           ├── problem.md
│           ├── solution.go
│           └── explanation.md
├── 02-bruteforce/
│   └── ... (동일 구조)
├── ... (03~26번 기존 폴더)
├── 27-geometry/                       # 확장 알고리즘 유형 폴더 (27~53)
│   └── ... (동일 구조)
├── ... (28~53번 확장 폴더)
├── 54-tree-dp/                        # Tier_1 알고리즘 유형 폴더 (54~59)
│   └── ... (동일 구조)
├── 55-lis/
│   └── ... (동일 구조)
├── 56-sqrt-decomposition/
│   └── ... (동일 구조)
├── 57-meet-in-the-middle/
│   └── ... (동일 구조)
├── 58-zero-one-bfs/
│   └── ... (동일 구조)
├── 59-flood-fill/
│   └── ... (동일 구조)
├── 60-fft/                            # Tier_2 알고리즘 유형 폴더 (60~67)
│   └── ... (동일 구조)
├── 61-ternary-search/
│   └── ... (동일 구조)
├── 62-euler-tour/
│   └── ... (동일 구조)
├── 63-mcmf/
│   └── ... (동일 구조)
├── 64-convex-hull-trick/
│   └── ... (동일 구조)
├── 65-gaussian-elimination/
│   └── ... (동일 구조)
├── 66-hld/
│   └── ... (동일 구조)
└── 67-centroid-decomposition/
    └── ... (동일 구조)
```

### 설계 결정 사항

1. **플랫 루트 구조**: 알고리즘 폴더를 루트에 직접 배치한다. 중간 카테고리 폴더(예: `basic/`, `graph/`)를 두지 않는다. 번호 접두사가 학습 순서를 충분히 표현하며, 깊은 중첩은 탐색을 어렵게 한다.

2. **번호 접두사 + 케밥 케이스**: `01-implementation-and-simulation` 형식을 사용한다. 파일 시스템에서 자연스러운 정렬이 되고, URL/경로에서도 가독성이 좋다.

3. **문제별 하위 폴더**: 각 문제를 독립 폴더로 분리한다. 문제 설명(problem.md), 풀이(solution.go), 해설(explanation.md)을 한 곳에 모아 관리한다.

4. **예시 코드 분리**: 이론 설명용 예시 코드는 `examples/` 폴더에, 문제 풀이 코드는 `problems/` 폴더에 분리한다. 학습 목적이 다르기 때문이다.

5. **단일 언어(Golang)**: 모든 코드를 Golang으로 통일한다. 표준 라이브러리만 사용하며, `go run`으로 즉시 실행 가능하게 한다.

6. **확장 알고리즘 번호 체계**: 27~53번은 기존 26개 이후에 이어지며, 우선순위_1(27~29), 우선순위_2(30~39), 우선순위_3(40~53) 순서로 배치한다. BOJ 문제 수 기준으로 학습 가치가 높은 순서이다. 54~59번은 Tier_1(코테 실전 필수), 60~67번은 Tier_2(코테 상위권 차별화)로 배치한다. Tier_1은 한국 코딩테스트(삼성 SW, 카카오, 프로그래머스, BOJ 골드)에서 직접 출제되는 유형이며, Tier_2는 BOJ 골드 상위~플래티넘 문제와 카카오 후반부 문제에서 등장하는 유형이다.

7. **선수 학습 관계 명시**: 확장 알고리즘이 기존 폴더와 관련이 있는 경우, theory.md에 선수 학습 폴더를 명시하여 학습 경로를 안내한다. 이를 통해 학습자가 기초부터 심화까지 자연스럽게 진행할 수 있다.

## 컴포넌트 및 인터페이스 (Components and Interfaces)

이 프로젝트는 런타임 소프트웨어가 아니므로, 컴포넌트는 콘텐츠 단위로 정의한다.

### 컴포넌트 1: 루트 안내 문서 (README.md)

프로젝트 진입점 역할을 하는 문서이다.

포함 내용:
- 프로젝트 목적 및 대상 독자
- 67개 알고리즘 유형 목록 (폴더 링크 포함)
- 기존 알고리즘(01~26번)과 확장 알고리즘(27~53번), Tier_1(54~59번), Tier_2(60~67번) 구분 표시
- 확장 알고리즘의 우선순위 그룹별(우선순위_1, 우선순위_2, 우선순위_3) 분류
- Tier_1(코테 실전 필수) 그룹 분류
- Tier_2(코테 상위권 차별화) 그룹 분류
- 권장 학습 순서 안내
- Golang 개발 환경 설정 가이드
- 프로젝트 사용 방법

### 컴포넌트 2: 알고리즘 폴더 README (각 폴더의 README.md)

각 알고리즘 폴더의 진입점이다.

포함 내용:
- 알고리즘 유형 이름과 간단한 설명
- 폴더 내 파일 목록 및 설명
- 포함된 문제 목록 (난이도 표기)

### 컴포넌트 3: 이론 문서 (theory.md)

알고리즘의 이론적 배경을 설명하는 문서이다.

포함 내용:
- 알고리즘 개념 정의
- 동작 원리 (단계별 설명)
- 시간 복잡도 / 공간 복잡도
- 적합한 문제 유형
- 단계별 추적 (Trace) - 구체적인 입력 예시를 통한 알고리즘 실행 과정 시각화
- 실전 팁 - 코딩테스트 활용 노하우, 자주 하는 실수, 엣지 케이스 주의사항
- 관련 알고리즘 비교 - 유사한 알고리즘과의 차이점 및 선택 기준
- 변형 및 최적화 (해당하는 경우) - 변형 알고리즘이나 최적화 기법을 별도 하위 섹션으로 포함
- 관련 알고리즘 참조
- 선수 학습 안내 (확장 알고리즘의 경우, 관련 기존 폴더 번호를 명시)

### 컴포넌트 4: 예시 코드 (examples/*.go)

알고리즘의 기본 구현을 보여주는 Golang 코드이다.

규칙:
- `package main` 선언
- `main()` 함수에서 실행 가능한 예제 포함
- 주요 로직에 한국어 주석
- 표준 라이브러리만 사용
- 변형이 여러 개인 경우 별도 파일로 분리 (예: `bubble_sort.go`, `merge_sort.go`)

### 컴포넌트 5: 문제 파일 (problems/XX-difficulty-name/problem.md)

코딩테스트 유형의 문제를 정의하는 문서이다.

포함 내용:
- 문제 제목 및 난이도 표기
- 문제 설명
- 입력/출력 형식
- 예제 입력/출력
- 제약 조건 (입력 범위, 시간 제한)

### 컴포넌트 6: 풀이 코드 (problems/XX-difficulty-name/solution.go)

문제의 Golang 풀이이다.

규칙:
- `package main` 선언
- 표준 입출력 사용 (`fmt`, `bufio`)
- `go run solution.go`로 실행 가능
- 주요 로직에 한국어 주석

### 컴포넌트 7: 해설 문서 (problems/XX-difficulty-name/explanation.md)

풀이의 접근 방식을 설명하는 문서이다.

포함 내용:
- 문제 접근 방식 (단계별)
- 핵심 아이디어 및 사용된 알고리즘 기법
- 시간 복잡도 / 공간 복잡도 분석
- 대안적 접근 방식 (존재 시)


## 데이터 모델 (Data Models)

이 프로젝트는 데이터베이스를 사용하지 않는다. 데이터 모델은 파일 시스템의 디렉토리/파일 구조와 문서 템플릿으로 정의한다.

### 모델 1: 알고리즘 폴더 구조

```
{번호}-{영문-케밥-케이스}/
├── README.md
├── theory.md
├── examples/
│   ├── {algorithm_name}.go
│   └── {variant_name}.go        # 변형이 있는 경우
└── problems/
    ├── 01-easy-{problem_name}/
    │   ├── problem.md
    │   ├── solution.go
    │   └── explanation.md
    ├── 02-medium-{problem_name}/
    │   ├── problem.md
    │   ├── solution.go
    │   └── explanation.md
    └── 03-hard-{problem_name}/
        ├── problem.md
        ├── solution.go
        └── explanation.md
```

### 모델 2: 알고리즘 유형 매핑

| 번호 | 영문 폴더명 | 한국어 이름 |
| --- | --- | --- |
| 01 | implementation-and-simulation | 구현과 시뮬레이션 |
| 02 | bruteforce | 브루트포스 |
| 03 | sorting | 정렬 |
| 04 | stack-and-queue | 스택과 큐 |
| 05 | hash | 해시 |
| 06 | prefix-sum | 누적합 |
| 07 | math-and-number-theory | 수학과 정수론 |
| 08 | binary-search | 이진 탐색 |
| 09 | parametric-search | 파라메트릭 서치 |
| 10 | two-pointer-and-sliding-window | 투 포인터와 슬라이딩 윈도우 |
| 11 | greedy | 그리디 |
| 12 | heap-and-priority-queue | 힙과 우선순위 큐 |
| 13 | tree | 트리 |
| 14 | binary-tree | 이진 트리 |
| 15 | graph-dfs | 그래프 탐색 DFS |
| 16 | graph-bfs | 그래프 탐색 BFS |
| 17 | backtracking | 백트래킹 |
| 18 | divide-and-conquer | 분할 정복 |
| 19 | dynamic-programming | 동적 프로그래밍 |
| 20 | union-find | 유니온 파인드 |
| 21 | shortest-path | 최단 경로 |
| 22 | minimum-spanning-tree | 최소 신장 트리 |
| 23 | topological-sort | 위상 정렬 |
| 24 | graph-advanced | 그래프 알고리즘 기타 |
| 25 | segment-tree | 세그먼트 트리 |
| 26 | string-algorithm | 문자열 알고리즘 |
| 27 | geometry | 기하학 |
| 28 | combinatorics | 조합론 |
| 29 | bitmask | 비트마스킹 |
| 30 | game-theory | 게임 이론 |
| 31 | probability | 확률론 |
| 32 | bitmask-dp | 비트필드 DP |
| 33 | maximum-flow | 최대 유량 |
| 34 | primality-test | 소수 판정 |
| 35 | offline-queries | 오프라인 쿼리 |
| 36 | exponentiation-by-squaring | 분할 정복 거듭제곱 |
| 37 | knapsack | 배낭 문제 |
| 38 | dag | DAG |
| 39 | coordinate-compression | 좌표 압축 |
| 40 | recursion | 재귀 |
| 41 | euclidean-algorithm | 유클리드 호제법 |
| 42 | convex-hull | 볼록 껍질 |
| 43 | bipartite-matching | 이분 매칭 |
| 44 | sieve-of-eratosthenes | 에라토스테네스의 체 |
| 45 | inclusion-exclusion | 포함 배제의 원리 |
| 46 | lca | LCA |
| 47 | sparse-table | 희소 배열 |
| 48 | hashing | 해싱 |
| 49 | modular-inverse | 모듈로 곱셈 역원 |
| 50 | floyd-warshall | 플로이드-워셜 |
| 51 | trie | 트라이 |
| 52 | deque | 덱 |
| 53 | prime-factorization | 소인수분해 |
| 54 | tree-dp | 트리 DP |
| 55 | lis | LIS (최장 증가 부분 수열) |
| 56 | sqrt-decomposition | 제곱근 분할법 |
| 57 | meet-in-the-middle | 중간에서 만나기 |
| 58 | zero-one-bfs | 0-1 BFS |
| 59 | flood-fill | 플러드 필 |
| 60 | fft | FFT (고속 푸리에 변환) |
| 61 | ternary-search | 삼분 탐색 |
| 62 | euler-tour | 오일러 경로 테크닉 |
| 63 | mcmf | 최소 비용 최대 유량 |
| 64 | convex-hull-trick | 볼록 껍질 트릭 |
| 65 | gaussian-elimination | 가우스 소거법 |
| 66 | hld | HLD |
| 67 | centroid-decomposition | 센트로이드 분할 |

### 모델 3: 문서 템플릿

#### 이론 문서 템플릿 (theory.md)

```markdown
# {알고리즘 한국어 이름}

## 개념

{알고리즘의 정의와 핵심 개념}

## 동작 원리

{단계별 동작 설명}

1. {단계 1}
2. {단계 2}
...

## 단계별 추적 (Trace)

{구체적인 입력 예시를 통한 알고리즘 실행 과정}

### 예시 입력

{입력 데이터}

### 실행 과정

{각 단계별 상태 변화를 ASCII 다이어그램이나 텍스트로 시각화}

## 복잡도

| 구분 | 복잡도 |
| --- | --- |
| 시간 복잡도 (최선) | O(...) |
| 시간 복잡도 (평균) | O(...) |
| 시간 복잡도 (최악) | O(...) |
| 공간 복잡도 | O(...) |

## 적합한 문제 유형

- {문제 유형 1}
- {문제 유형 2}
...

## 실전 팁

### 활용 노하우

- {코딩테스트에서의 활용 팁}

### 자주 하는 실수

- {흔한 실수와 해결 방법}

### 엣지 케이스

- {주의해야 할 엣지 케이스}

## 관련 알고리즘 비교

| 알고리즘 | 특징 | 적합한 상황 |
| --- | --- | --- |
| {알고리즘 A} | {특징} | {상황} |
| {알고리즘 B} | {특징} | {상황} |

## 변형 및 최적화 (해당하는 경우)

### {변형 알고리즘 이름}

{변형 알고리즘 설명}

### {최적화 기법 이름}

{최적화 기법 설명}
```

#### 문제 파일 템플릿 (problem.md)

```markdown
# {문제 제목}

**난이도:** {하/중/상}

## 문제 설명

{문제 내용}

## 입력 형식

{입력 설명}

## 출력 형식

{출력 설명}

## 제약 조건

- {제약 조건 1}
- {제약 조건 2}

## 예제

### 예제 입력 1

```
{입력}
```

### 예제 출력 1

```
{출력}
```
```

#### 해설 문서 템플릿 (explanation.md)

```markdown
# {문제 제목} - 해설

## 접근 방식

{단계별 접근 방식 설명}

## 핵심 아이디어

{사용된 알고리즘 기법과 핵심 아이디어}

## 복잡도 분석

| 구분 | 복잡도 |
| --- | --- |
| 시간 복잡도 | O(...) |
| 공간 복잡도 | O(...) |

## 대안적 접근 (선택)

{다른 풀이 방법이 있는 경우}
```

### 모델 4: 예시 코드 구조

```go
package main

import "fmt"

// {알고리즘 이름} - {간단한 설명}
// 시간 복잡도: O(...)
// 공간 복잡도: O(...)

func algorithmName(params) returnType {
    // {한국어 주석: 핵심 로직 설명}
    ...
}

func main() {
    // 실행 예제
    ...
    fmt.Println(result)
}
```

### 모델 5: 풀이 코드 구조

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    // {한국어 주석: 입력 처리}
    ...

    // {한국어 주석: 핵심 로직}
    ...

    // {한국어 주석: 출력}
    fmt.Fprintln(writer, result)
}
```

### 모델 6: 문제 난이도 분류 체계

| 난이도 | 표기 | 폴더 접두사 | 설명 |
| --- | --- | --- | --- |
| 하 | 쉬움 | easy | 해당 알고리즘의 기본 적용 |
| 중 | 보통 | medium | 알고리즘 응용 또는 조합 |
| 상 | 어려움 | hard | 복잡한 응용, 최적화 필요 |

### 모델 7: 폴더 README 구조

```markdown
# {번호}. {알고리즘 한국어 이름}

{알고리즘에 대한 한 줄 설명}

## 구성

- `theory.md` - 이론 및 개념 설명
- `examples/` - 기본 구현 예시 코드
- `problems/` - 연습 문제 및 풀이

## 문제 목록

| 번호 | 문제 | 난이도 |
| --- | --- | --- |
| 01 | {문제명} | 하 |
| 02 | {문제명} | 중 |
| 03 | {문제명} | 상 |
```


## 정확성 속성 (Correctness Properties)

*속성(property)이란 시스템의 모든 유효한 실행에서 참이어야 하는 특성 또는 동작이다. 속성은 사람이 읽을 수 있는 명세와 기계가 검증할 수 있는 정확성 보장 사이의 다리 역할을 한다.*

이 프로젝트는 정적 콘텐츠 저장소이므로, 정확성 속성은 파일 시스템 구조와 문서/코드의 형식 규칙에 초점을 맞춘다. 검증 스크립트를 통해 자동화된 테스트가 가능하다.

### Property 1: 폴더 구조 완전성

*For any* 알고리즘 폴더(01~67번), 해당 폴더는 반드시 다음 파일/디렉토리를 모두 포함해야 한다: README.md, theory.md, examples/ 디렉토리(최소 1개의 .go 파일 포함), problems/ 디렉토리(최소 3개의 문제 폴더 포함, 각 문제 폴더에 problem.md, solution.go, explanation.md 포함).

**Validates: Requirements 1.5, 2.1, 3.1, 4.1, 4.6, 5.1, 10.1, 10.2, 11.1, 11.2**

### Property 2: 폴더 명명 규칙

*For any* 알고리즘 폴더명(01~67번), 해당 이름은 `^[0-9]{2}-[a-z]+(-[a-z0-9]+)*$` 정규식 패턴(두 자리 번호 + 영문 케밥 케이스)을 만족해야 한다.

**Validates: Requirements 1.3, 10.3, 11.3**

### Property 3: 이론 문서 구조 완전성

*For any* 이론 문서(theory.md), 해당 문서는 반드시 다음 섹션을 모두 포함해야 한다: "개념" 섹션, "동작 원리" 섹션, "복잡도" 섹션(시간 복잡도와 공간 복잡도 모두 포함), "적합한 문제 유형" 섹션.

**Validates: Requirements 2.2, 2.3, 2.4, 2.5**

### Property 4: 문제 파일 구조 완전성

*For any* 문제 파일(problem.md), 해당 문서는 반드시 다음을 모두 포함해야 한다: 문제 설명, "입력 형식" 섹션, "출력 형식" 섹션, "예제 입력"과 "예제 출력", "제약 조건" 섹션, 그리고 난이도 표기("하", "중", "상" 중 하나).

**Validates: Requirements 4.2, 4.3, 4.4, 4.5, 4.8**

### Property 5: 해설 문서 구조 완전성

*For any* 해설 문서(explanation.md), 해당 문서는 반드시 다음 섹션을 모두 포함해야 한다: "접근 방식" 섹션, "핵심 아이디어" 섹션, "복잡도 분석" 섹션(시간 복잡도와 공간 복잡도 모두 포함).

**Validates: Requirements 5.2, 5.3, 5.4**

### Property 6: 난이도 분포 균형

*For any* 알고리즘 폴더의 problems/ 디렉토리, 해당 디렉토리는 "easy" 접두사 폴더 최소 1개, "medium" 접두사 폴더 최소 1개, "hard" 접두사 폴더 최소 1개를 포함하여 세 난이도가 모두 존재해야 한다.

**Validates: Requirements 4.9**

### Property 7: Go 코드 규칙 준수

*For any* .go 파일, 해당 파일은 다음 규칙을 모두 만족해야 한다: .go 확장자 사용, `package main` 선언 포함, `func main()` 함수 포함, 한국어 주석 최소 1개 포함, 표준 라이브러리만 import.

**Validates: Requirements 3.2, 3.4, 3.5, 7.2, 7.3**

### Property 8: 이론 문서 상세 보충 완전성

*For any* 이론 문서(theory.md, 01~67번 전체), 해당 문서는 반드시 다음 섹션을 모두 포함해야 한다: "단계별 추적" 또는 "Trace" 섹션, "실전 팁" 섹션, "관련 알고리즘 비교" 섹션.

**Validates: Requirements 12.2, 12.4, 12.5, 12.9**$` 정규식 패턴(두 자리 번호 + 영문 케밥 케이스)을 만족해야 한다.

**Validates: Requirements 1.3**

### Property 3: 이론 문서 구조 완전성

*For any* 이론 문서(theory.md), 해당 문서는 반드시 다음 섹션을 모두 포함해야 한다: "개념" 섹션, "동작 원리" 섹션, "복잡도" 섹션(시간 복잡도와 공간 복잡도 모두 포함), "적합한 문제 유형" 섹션.

**Validates: Requirements 2.2, 2.3, 2.4, 2.5**

### Property 4: 문제 파일 구조 완전성

*For any* 문제 파일(problem.md), 해당 문서는 반드시 다음을 모두 포함해야 한다: 문제 설명, "입력 형식" 섹션, "출력 형식" 섹션, "예제 입력"과 "예제 출력", "제약 조건" 섹션, 그리고 난이도 표기("하", "중", "상" 중 하나).

**Validates: Requirements 4.2, 4.3, 4.4, 4.5, 4.8**

### Property 5: 해설 문서 구조 완전성

*For any* 해설 문서(explanation.md), 해당 문서는 반드시 다음 섹션을 모두 포함해야 한다: "접근 방식" 섹션, "핵심 아이디어" 섹션, "복잡도 분석" 섹션(시간 복잡도와 공간 복잡도 모두 포함).

**Validates: Requirements 5.2, 5.3, 5.4**

### Property 6: 난이도 분포 균형

*For any* 알고리즘 폴더의 problems/ 디렉토리, 해당 디렉토리는 "easy" 접두사 폴더 최소 1개, "medium" 접두사 폴더 최소 1개, "hard" 접두사 폴더 최소 1개를 포함하여 세 난이도가 모두 존재해야 한다.

**Validates: Requirements 4.9**

### Property 7: Go 코드 규칙 준수

*For any* .go 파일, 해당 파일은 다음 규칙을 모두 만족해야 한다: .go 확장자 사용, `package main` 선언 포함, `func main()` 함수 포함, 한국어 주석 최소 1개 포함, 표준 라이브러리만 import.

**Validates: Requirements 3.2, 3.4, 3.5, 7.2, 7.3**

## 오류 처리 (Error Handling)

이 프로젝트는 런타임 애플리케이션이 아닌 정적 콘텐츠 저장소이므로, 전통적인 오류 처리 대신 콘텐츠 품질 보장 관점에서 오류를 정의하고 대응한다.

### 콘텐츠 오류 유형

| 오류 유형 | 설명 | 대응 방식 |
| --- | --- | --- |
| 구조 누락 | 필수 파일/디렉토리 누락 | 검증 스크립트로 사전 탐지 후 보완 |
| 템플릿 불일치 | 문서가 정의된 템플릿 섹션을 누락 | 검증 스크립트로 사전 탐지 후 보완 |
| 명명 규칙 위반 | 폴더/파일명이 규칙에 맞지 않음 | 검증 스크립트로 사전 탐지 후 수정 |
| 난이도 불균형 | 특정 난이도의 문제가 누락됨 | 검증 스크립트로 사전 탐지 후 문제 추가 |

### 검증 스크립트 전략

프로젝트 루트에 `validate.sh` 스크립트를 제공하여 위 오류들을 자동으로 탐지한다. 이 스크립트는 다음을 수행한다:

1. 67개 알고리즘 폴더 존재 여부 확인
2. 각 폴더 내 필수 파일 존재 여부 확인
3. 문서 템플릿 필수 섹션 존재 여부 확인
4. 난이도 분포 균형 확인
5. 폴더/파일 명명 규칙 준수 확인

## 테스트 전략 (Testing Strategy)

### 테스트 접근 방식

이 프로젝트는 정적 콘텐츠 저장소이므로, 테스트는 두 가지 축으로 진행한다:

1. **단위 테스트 (Unit Tests)**: 특정 예시와 엣지 케이스를 검증
2. **속성 기반 테스트 (Property-Based Tests)**: 모든 콘텐츠에 대해 보편적 규칙을 검증

두 접근 방식은 상호 보완적이다. 단위 테스트는 구체적인 버그를 잡고, 속성 기반 테스트는 전체적인 정확성을 보장한다.

### 속성 기반 테스트 (Property-Based Tests)

검증 스크립트를 Go 테스트로 구현하며, `testing/quick` 패키지 또는 외부 PBT 라이브러리(`github.com/leanovate/gopter`)를 사용한다.

각 속성 테스트는 최소 100회 반복 실행하며, 설계 문서의 속성을 참조하는 태그를 포함한다.

| 테스트 | 대상 속성 | 태그 |
| --- | --- | --- |
| TestFolderStructureCompleteness | Property 1 | Feature: algorithm-study-guide, Property 1: 폴더 구조 완전성 |
| TestFolderNamingConvention | Property 2 | Feature: algorithm-study-guide, Property 2: 폴더 명명 규칙 |
| TestTheoryDocStructure | Property 3 | Feature: algorithm-study-guide, Property 3: 이론 문서 구조 완전성 |
| TestProblemDocStructure | Property 4 | Feature: algorithm-study-guide, Property 4: 문제 파일 구조 완전성 |
| TestExplanationDocStructure | Property 5 | Feature: algorithm-study-guide, Property 5: 해설 문서 구조 완전성 |
| TestDifficultyDistribution | Property 6 | Feature: algorithm-study-guide, Property 6: 난이도 분포 균형 |
| TestGoCodeConventions | Property 7 | Feature: algorithm-study-guide, Property 7: Go 코드 규칙 준수 |
| TestTheoryDocDetailedSections | Property 8 | Feature: algorithm-study-guide, Property 8: 이론 문서 상세 보충 완전성 |

### 단위 테스트 (Unit Tests)

단위 테스트는 속성 테스트로 커버하기 어려운 구체적 사례와 엣지 케이스에 집중한다:

- 루트 README.md에 67개 알고리즘 유형 목록과 링크가 모두 존재하는지 확인 (요구사항 6.3, 9.1)
- 루트 README.md에 확장 알고리즘이 우선순위 그룹별(우선순위_1, 우선순위_2, 우선순위_3)과 Tier_1, Tier_2로 구분되어 표시되는지 확인 (요구사항 9.2)
- 루트 README.md에 목적, 대상 독자, 학습 순서, 환경 설정 섹션이 존재하는지 확인 (요구사항 6.2, 6.4, 6.5)
- 알고리즘 유형 매핑 테이블의 67개 항목이 정확히 일치하는지 확인 (요구사항 1.2, 1.4, 8.1, 10.1, 11.1)
- 정렬 폴더처럼 여러 변형이 있는 알고리즘의 examples/ 폴더에 복수의 .go 파일이 존재하는지 확인 (요구사항 3.6)
- Tier_1(54~59번) 알고리즘의 theory.md에 지정된 선수 학습 폴더가 참조되어 있는지 확인 (요구사항 10.5, 10.6)
- Tier_2(60~67번) 알고리즘의 theory.md에 지정된 선수 학습 폴더가 참조되어 있는지 확인 (요구사항 11.5, 11.6)

### 테스트 실행 방법

```bash
# 전체 검증 스크립트 실행
bash validate.sh

# Go 테스트 실행 (속성 기반 + 단위 테스트)
go test ./validate/... -v -count=1
```
