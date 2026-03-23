# 구현 계획: HackerRank 스타일 리팩토링

## 개요

67개 알고리즘 폴더(01~67)의 201개 문제 폴더에 있는 solution.go를 HackerRank 스타일로 분리한다. solution.go를 보일러플레이트 + 빈 핵심 함수(상세 한국어 설명 포함)로 재작성하고, answer.go를 동일 보일러플레이트 + 핵심 함수 구현 완성본으로 작성한다. 검증 스크립트(validate.sh, validate_test.go)도 answer.go를 포함하도록 업데이트한다.

## Tasks

- [x] 1. validate.sh 및 validate_test.go 업데이트
  - [x] 1.1 validate.sh에 answer.go 존재 여부 검사 추가
    - 각 문제 폴더 순회 시 `answer.go` 파일 존재 여부를 검사하는 로직 추가
    - 기존 `solution.go` 검사는 그대로 유지
    - _Requirements: 4.1, 4.2_

  - [x] 1.2 validate_test.go의 TestFolderStructureCompleteness에 answer.go 존재 확인 추가
    - 각 문제 폴더에서 `answer.go` 파일 존재 여부를 검증하는 코드 추가
    - _Requirements: 4.5_

  - [x] 1.3 validate_test.go의 TestGoCodeConventions에서 answer.go도 검증하도록 수정
    - `collectAllGoFiles`로 수집되는 파일에 answer.go가 자동 포함되므로, 기존 로직이 answer.go에도 적용되는지 확인
    - answer.go에 대해 package main, func main(), 한국어 주석, 표준 라이브러리만 import 규칙 검증
    - _Requirements: 4.3, 4.4_

  - [x] 1.4 Property 1 속성 테스트 작성: TestProblemFolderFileCompleteness
    - **Property 1: 문제 폴더 파일 완전성**
    - 모든 201개 문제 폴더에 problem.md, solution.go, answer.go, explanation.md 4개 파일이 존재하는지 검증
    - **Validates: Requirements 1.7, 5.1**

  - [x] 1.5 Property 2 속성 테스트 작성: TestAnswerGoCodeConventions
    - **Property 2: answer.go Go 코드 규칙 준수**
    - 모든 answer.go 파일이 package main, func main(), 한국어 주석, 표준 라이브러리만 import 규칙을 만족하는지 검증
    - **Validates: Requirements 1.1, 3.5, 3.6, 3.7, 3.8**

  - [x] 1.6 Property 3 속성 테스트 작성: TestSolutionGoCodeConventions
    - **Property 3: solution.go Go 코드 규칙 준수**
    - 모든 solution.go(problems/ 하위) 파일이 package main, func main(), 한국어 주석, 표준 라이브러리만 import 규칙을 만족하는지 검증
    - **Validates: Requirements 1.3, 3.1, 3.2, 3.3, 3.4**

  - [x] 1.7 Property 4 속성 테스트 작성: TestSolutionGoEmptyFunctionPattern
    - **Property 4: solution.go 빈 핵심 함수 패턴**
    - 모든 solution.go(problems/ 하위) 파일이 main 외 최소 1개 함수를 포함하고, `// 여기에 코드를 작성하세요` 주석을 포함하는지 검증
    - **Validates: Requirements 1.2, 1.4, 2.2**

- [x] 2. 폴더 01~10 리팩토링
  - [x] 2.1 01-implementation-and-simulation 리팩토링 (3개 문제)
    - 각 문제 폴더에서: 기존 코드 분석 → solution.go를 보일러플레이트 + 빈 핵심 함수(상세 설명)로 재작성, answer.go를 동일 보일러플레이트 + 함수 구현 완성본으로 작성
    - 핵심 함수 추출: 기존 코드 분석하여 알고리즘 로직을 함수로 분리
    - solution.go 빈 함수에 한국어 설명 주석(목적, 매개변수, 반환값) + `// 여기에 코드를 작성하세요` + 제로값 반환 포함
    - answer.go 함수에 한국어 설명 주석(목적, 매개변수, 반환값, 알고리즘 힌트) + 완전한 구현 코드 포함
    - _Requirements: 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 2.1, 2.2, 2.3, 2.4, 2.5, 2.6, 3.1, 3.2, 3.3, 3.4, 3.5, 3.6, 3.7, 3.8, 3.9, 5.1, 5.2_

  - [x] 2.2 02-bruteforce 리팩토링 (3개 문제)
    - 동일한 리팩토링 절차 적용
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 2.3 03-sorting 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 2.4 04-stack-and-queue 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 2.5 05-hash 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 2.6 06-prefix-sum 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 2.7 07-math-and-number-theory 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 2.8 08-binary-search 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 2.9 09-parametric-search 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 2.10 10-two-pointer-and-sliding-window 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

- [x] 3. 체크포인트 - 폴더 01~10 검증
  - `bash validate.sh` 실행하여 폴더 01~10의 구조 검증
  - `go test ./validate/... -v -count=1` 실행하여 테스트 통과 확인
  - 문제 발생 시 사용자에게 질문

- [x] 4. 폴더 11~20 리팩토링
  - [x] 4.1 11-greedy 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 4.2 12-heap-and-priority-queue 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 4.3 13-tree 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 4.4 14-binary-tree 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 4.5 15-graph-dfs 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 4.6 16-graph-bfs 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 4.7 17-backtracking 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 4.8 18-divide-and-conquer 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 4.9 19-dynamic-programming 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 4.10 20-union-find 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

- [x] 5. 체크포인트 - 폴더 11~20 검증
  - `bash validate.sh` 실행하여 폴더 11~20의 구조 검증
  - `go test ./validate/... -v -count=1` 실행하여 테스트 통과 확인
  - 문제 발생 시 사용자에게 질문

- [x] 6. 폴더 21~30 리팩토링
  - [x] 6.1 21-shortest-path 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 6.2 22-minimum-spanning-tree 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 6.3 23-topological-sort 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 6.4 24-graph-advanced 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 6.5 25-segment-tree 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 6.6 26-string-algorithm 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 6.7 27-geometry 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 6.8 28-combinatorics 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 6.9 29-bitmask 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 6.10 30-game-theory 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

- [x] 7. 체크포인트 - 폴더 21~30 검증
  - `bash validate.sh` 실행하여 폴더 21~30의 구조 검증
  - `go test ./validate/... -v -count=1` 실행하여 테스트 통과 확인
  - 문제 발생 시 사용자에게 질문

- [x] 8. 폴더 31~40 리팩토링
  - [x] 8.1 31-probability 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 8.2 32-bitmask-dp 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 8.3 33-maximum-flow 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 8.4 34-primality-test 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 8.5 35-offline-queries 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 8.6 36-exponentiation-by-squaring 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 8.7 37-knapsack 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 8.8 38-dag 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 8.9 39-coordinate-compression 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 8.10 40-recursion 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

- [x] 9. 체크포인트 - 폴더 31~40 검증
  - `bash validate.sh` 실행하여 폴더 31~40의 구조 검증
  - `go test ./validate/... -v -count=1` 실행하여 테스트 통과 확인
  - 문제 발생 시 사용자에게 질문

- [x] 10. 폴더 41~53 리팩토링
  - [x] 10.1 41-euclidean-algorithm 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 10.2 42-convex-hull 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 10.3 43-bipartite-matching 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 10.4 44-sieve-of-eratosthenes 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 10.5 45-inclusion-exclusion 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 10.6 46-lca 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 10.7 47-sparse-table 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 10.8 48-hashing 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 10.9 49-modular-inverse 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 10.10 50-floyd-warshall 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 10.11 51-trie 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 10.12 52-deque 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 10.13 53-prime-factorization 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

- [x] 11. 체크포인트 - 폴더 41~53 검증
  - `bash validate.sh` 실행하여 폴더 41~53의 구조 검증
  - `go test ./validate/... -v -count=1` 실행하여 테스트 통과 확인
  - 문제 발생 시 사용자에게 질문

- [x] 12. 폴더 54~67 리팩토링
  - [x] 12.1 54-tree-dp 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 12.2 55-lis 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 12.3 56-sqrt-decomposition 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 12.4 57-meet-in-the-middle 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 12.5 58-zero-one-bfs 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 12.6 59-flood-fill 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 12.7 60-fft 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 12.8 61-ternary-search 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 12.9 62-euler-tour 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 12.10 63-mcmf 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 12.11 64-convex-hull-trick 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 12.12 65-gaussian-elimination 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 12.13 66-hld 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

  - [x] 12.14 67-centroid-decomposition 리팩토링 (3개 문제)
    - _Requirements: 1.1~1.7, 2.1~2.6, 3.1~3.9, 5.1~5.3_

- [x] 13. 최종 체크포인트 - 전체 검증
  - `bash validate.sh` 실행하여 전체 67개 폴더 구조 검증
  - `go test ./validate/... -v -count=1` 실행하여 모든 테스트 통과 확인
  - 201개 문제 폴더 모두에 answer.go가 존재하고, solution.go가 빈 핵심 함수 패턴을 따르는지 최종 확인
  - 문제 발생 시 사용자에게 질문

## 참고

- `*` 표시된 태스크는 선택 사항이며 빠른 MVP를 위해 건너뛸 수 있다
- 각 리팩토링 태스크의 공통 절차: (1) 기존 solution.go 분석하여 핵심 함수 추출, (2) solution.go를 보일러플레이트 + 빈 핵심 함수(상세 한국어 설명 포함)로 재작성, (3) answer.go를 동일 보일러플레이트 + 핵심 함수 구현 완성본으로 작성
- 핵심 함수 추출 시 설계 문서의 패턴 A~D 및 함수명 명명 규칙을 따른다
- solution.go의 빈 핵심 함수에는 반드시 한국어 설명 주석(목적, 매개변수, 반환값), `// 여기에 코드를 작성하세요` 안내 주석, 제로값 반환문을 포함한다 (알고리즘 힌트는 포함하지 않는다)
- answer.go의 핵심 함수에는 한국어 설명 주석(목적, 매개변수, 반환값, 알고리즘 힌트)과 완전한 구현 코드를 포함한다
- 속성 테스트는 설계 문서의 정확성 속성(Property 1~4)을 검증한다
