# 구현 계획: HackerRank 문제 학습 가이드

## 개요

HackerRank 문제 80개(8개 서브카테고리 × 10문제)를 자료구조와 알고리즘 두 대분류로 나눈 학습 가이드를 생성한다. 기존 워크스페이스의 파일 패턴(`problem.md`, `explanation.md`, `problem.go`, `answer.go`)과 한국어 문서 작성 방식을 그대로 유지한다. 모든 Go 코드는 컴파일 가능해야 한다.

## Tasks

- [x] 1. 최상위 폴더 구조 및 README 생성
  - `hackerrank-problems/` 폴더 생성
  - `hackerrank-problems/README.md` 작성 (프로젝트 개요, 카테고리 구성, 문제 수, 난이도 분포를 한국어로 포함)
  - `hackerrank-problems/data-structures/` 폴더 생성
  - `hackerrank-problems/algorithms/` 폴더 생성
  - _요구사항: 1.1, 1.2, 1.3, 1.4_

- [x] 2. 자료구조 카테고리 구조 생성
  - [x] 2.1 자료구조 카테고리 README 및 디렉토리 생성
    - `data-structures/README.md` 작성 (카테고리 개요, 세부 주제 목록, 문제 수)
    - `data-structures/examples/` 디렉토리 생성
    - `data-structures/problems/` 디렉토리 생성
    - `data-structures/problems/` 아래에 `linked-list/`, `tree/`, `heap/`, `stack-queue/` 서브카테고리 폴더 생성
    - _요구사항: 2.1, 2.4_

  - [x] 2.2 자료구조 theory.md 작성
    - `data-structures/theory.md` 작성
    - 연결 리스트, 트리, 힙, 스택/큐 각각의 개념, 동작 원리, 시간/공간 복잡도 비교, 실전 팁, 자료구조 선택 기준을 한국어로 포함
    - 기존 `01-implementation-and-simulation/theory.md` 패턴(개념, 동작 원리, 복잡도, 적합한 문제 유형, 단계별 추적, 실전 팁, 관련 알고리즘 비교)을 따름
    - _요구사항: 2.2, 14.1_

  - [x] 2.3 자료구조 예시 코드 작성 — linked_list.go
    - `data-structures/examples/linked_list.go` 작성
    - 연결 리스트의 기본 연산(삽입, 삭제, 탐색, 출력)을 보여주는 Go 예시 코드
    - 파일 상단에 시간/공간 복잡도 주석, 각 함수에 한국어 주석, main 함수에서 예시 실행 결과 출력
    - 기존 `01-implementation-and-simulation/examples/simulation.go` 패턴을 따름
    - _요구사항: 2.3, 15.1, 15.8, 15.9, 15.10_

  - [x] 2.4 자료구조 예시 코드 작성 — tree.go
    - `data-structures/examples/tree.go` 작성
    - 이진 트리의 기본 연산(삽입, 순회)을 보여주는 Go 예시 코드
    - 파일 상단에 시간/공간 복잡도 주석, 각 함수에 한국어 주석, main 함수에서 예시 실행 결과 출력
    - _요구사항: 2.3, 15.2, 15.8, 15.9, 15.10_

  - [x] 2.5 자료구조 예시 코드 작성 — heap.go
    - `data-structures/examples/heap.go` 작성
    - 힙(우선순위 큐)의 기본 연산(삽입, 추출)을 보여주는 Go 예시 코드
    - 파일 상단에 시간/공간 복잡도 주석, 각 함수에 한국어 주석, main 함수에서 예시 실행 결과 출력
    - _요구사항: 2.3, 15.3, 15.8, 15.9, 15.10_

- [x] 3. 알고리즘 카테고리 구조 생성
  - [x] 3.1 알고리즘 카테고리 README 및 디렉토리 생성
    - `algorithms/README.md` 작성 (카테고리 개요, 세부 주제 목록, 문제 수)
    - `algorithms/examples/` 디렉토리 생성
    - `algorithms/problems/` 디렉토리 생성
    - `algorithms/problems/` 아래에 `greedy/`, `dynamic-programming/`, `graph/`, `search/` 서브카테고리 폴더 생성
    - _요구사항: 3.1, 3.4_

  - [x] 3.2 알고리즘 theory.md 작성
    - `algorithms/theory.md` 작성
    - 그리디, 동적 프로그래밍, 그래프, 탐색 각각의 개념, 동작 원리, 시간/공간 복잡도 비교, 실전 팁, 알고리즘 선택 기준을 한국어로 포함
    - 기존 theory.md 패턴을 따름
    - _요구사항: 3.2, 14.1_

  - [x] 3.3 알고리즘 예시 코드 작성 — greedy.go
    - `algorithms/examples/greedy.go` 작성
    - 그리디 알고리즘의 기본 패턴을 보여주는 Go 예시 코드
    - 파일 상단에 시간/공간 복잡도 주석, 각 함수에 한국어 주석, main 함수에서 예시 실행 결과 출력
    - _요구사항: 3.3, 15.4, 15.8, 15.9, 15.10_

  - [x] 3.4 알고리즘 예시 코드 작성 — dp.go
    - `algorithms/examples/dp.go` 작성
    - 동적 프로그래밍의 기본 패턴(메모이제이션, 타뷸레이션)을 보여주는 Go 예시 코드
    - 파일 상단에 시간/공간 복잡도 주석, 각 함수에 한국어 주석, main 함수에서 예시 실행 결과 출력
    - _요구사항: 3.3, 15.5, 15.8, 15.9, 15.10_

  - [x] 3.5 알고리즘 예시 코드 작성 — graph.go
    - `algorithms/examples/graph.go` 작성
    - 그래프 탐색(BFS, DFS)의 기본 패턴을 보여주는 Go 예시 코드
    - 파일 상단에 시간/공간 복잡도 주석, 각 함수에 한국어 주석, main 함수에서 예시 실행 결과 출력
    - _요구사항: 3.3, 15.6, 15.8, 15.9, 15.10_

  - [x] 3.6 알고리즘 예시 코드 작성 — search.go
    - `algorithms/examples/search.go` 작성
    - 이진 탐색의 기본 패턴을 보여주는 Go 예시 코드
    - 파일 상단에 시간/공간 복잡도 주석, 각 함수에 한국어 주석, main 함수에서 예시 실행 결과 출력
    - _요구사항: 3.3, 15.7, 15.8, 15.9, 15.10_

- [x] 4. 체크포인트 — 카테고리 구조 확인
  - 모든 카테고리 폴더, README, theory.md, examples/ 파일이 올바르게 생성되었는지 확인
  - 모든 Go 예시 코드가 컴파일 가능한지 확인
  - 문제가 있으면 사용자에게 질문

- [x] 5. 연결 리스트 문제 10개 생성 (linked-list/)
  - `data-structures/problems/linked-list/` 아래에 10개 문제 폴더 생성
  - 각 문제 폴더에 `problem.md`, `explanation.md`, `problem.go`, `answer.go` 4개 파일 생성
  - 문제 목록:
    - `01-easy-print-the-elements-of-a-linked-list` (https://www.hackerrank.com/challenges/print-the-elements-of-a-linked-list)
    - `02-easy-insert-a-node-at-the-tail` (https://www.hackerrank.com/challenges/insert-a-node-at-the-tail-of-a-linked-list)
    - `03-easy-insert-a-node-at-the-head` (https://www.hackerrank.com/challenges/insert-a-node-at-the-head-of-a-linked-list)
    - `04-easy-delete-a-node` (https://www.hackerrank.com/challenges/delete-a-node-from-a-linked-list)
    - `05-easy-reverse-a-linked-list` (https://www.hackerrank.com/challenges/reverse-a-linked-list)
    - `06-medium-compare-two-linked-lists` (https://www.hackerrank.com/challenges/compare-two-linked-lists)
    - `07-medium-merge-two-sorted-linked-lists` (https://www.hackerrank.com/challenges/merge-two-sorted-linked-lists)
    - `08-medium-get-node-value` (https://www.hackerrank.com/challenges/get-the-value-of-the-node-at-a-specific-position-from-the-tail)
    - `09-medium-insert-a-node-at-a-specific-position` (https://www.hackerrank.com/challenges/insert-a-node-at-a-specific-position-in-a-linked-list)
    - `10-hard-reverse-a-doubly-linked-list` (https://www.hackerrank.com/challenges/reverse-a-doubly-linked-list)
  - `problem.md`: 난이도, HackerRank 출처 URL, 문제 설명, 입력/출력 형식, 제약 조건, 예제를 한국어로 포함
  - `explanation.md`: 접근 방식, 핵심 아이디어, 복잡도 분석, 대안적 접근을 한국어로 포함
  - `problem.go`: 핵심 함수 시그니처 + `// 여기에 코드를 작성하세요` 주석 + bufio 기반 main 함수
  - `answer.go`: 완전한 구현 + 단계별 한국어 주석 + bufio I/O + `[알고리즘 힌트]` 함수 주석
  - _요구사항: 4.1, 4.2, 4.3, 12.1, 12.2, 12.3, 12.4, 12.5, 13.1, 13.2, 13.3, 13.4, 13.5, 13.6, 14.1, 14.2, 14.3_

- [x] 6. 트리 문제 10개 생성 (tree/)
  - `data-structures/problems/tree/` 아래에 10개 문제 폴더 생성
  - 각 문제 폴더에 `problem.md`, `explanation.md`, `problem.go`, `answer.go` 4개 파일 생성
  - 문제 목록:
    - `01-easy-tree-preorder-traversal` (https://www.hackerrank.com/challenges/tree-preorder-traversal)
    - `02-easy-tree-postorder-traversal` (https://www.hackerrank.com/challenges/tree-postorder-traversal)
    - `03-easy-tree-inorder-traversal` (https://www.hackerrank.com/challenges/tree-inorder-traversal)
    - `04-easy-tree-height-of-a-binary-tree` (https://www.hackerrank.com/challenges/tree-height-of-a-binary-tree)
    - `05-easy-tree-level-order-traversal` (https://www.hackerrank.com/challenges/tree-level-order-traversal)
    - `06-medium-tree-top-view` (https://www.hackerrank.com/challenges/tree-top-view)
    - `07-medium-binary-search-tree-insertion` (https://www.hackerrank.com/challenges/binary-search-tree-insertion)
    - `08-medium-binary-search-tree-lowest-common-ancestor` (https://www.hackerrank.com/challenges/binary-search-tree-lowest-common-ancestor)
    - `09-medium-is-this-a-binary-search-tree` (https://www.hackerrank.com/challenges/is-binary-search-tree)
    - `10-hard-swap-nodes-algo` (https://www.hackerrank.com/challenges/swap-nodes-algo)
  - `problem.md`: 난이도, HackerRank 출처 URL, 문제 설명, 입력/출력 형식, 제약 조건, 예제를 한국어로 포함
  - `explanation.md`: 접근 방식, 핵심 아이디어, 복잡도 분석, 대안적 접근을 한국어로 포함
  - `problem.go`: 핵심 함수 시그니처 + `// 여기에 코드를 작성하세요` 주석 + bufio 기반 main 함수
  - `answer.go`: 완전한 구현 + 단계별 한국어 주석 + bufio I/O + `[알고리즘 힌트]` 함수 주석
  - _요구사항: 5.1, 5.2, 5.3, 12.1, 12.2, 12.3, 12.4, 12.5, 13.1, 13.2, 13.3, 13.4, 13.5, 13.6, 14.1, 14.2, 14.3_

- [x] 7. 힙 문제 10개 생성 (heap/)
  - `data-structures/problems/heap/` 아래에 10개 문제 폴더 생성
  - 각 문제 폴더에 `problem.md`, `explanation.md`, `problem.go`, `answer.go` 4개 파일 생성
  - 문제 목록:
    - `01-easy-jesse-and-cookies` (https://www.hackerrank.com/challenges/jesse-and-cookies)
    - `02-easy-qheap1` (https://www.hackerrank.com/challenges/qheap1)
    - `03-easy-minimum-average-waiting-time` (https://www.hackerrank.com/challenges/minimum-average-waiting-time)
    - `04-medium-find-the-running-median` (https://www.hackerrank.com/challenges/find-the-running-median)
    - `05-medium-components-in-a-graph` (https://www.hackerrank.com/challenges/components-in-graph)
    - `06-medium-kundu-and-tree` (https://www.hackerrank.com/challenges/kundu-and-tree)
    - `07-medium-heap-full-sort` (https://www.hackerrank.com/challenges/heapsort)
    - `08-hard-median-updates` (https://www.hackerrank.com/challenges/median)
    - `09-hard-kth-minimum` (https://www.hackerrank.com/challenges/kth-minimum-in-range)
    - `10-hard-largest-rectangle` (https://www.hackerrank.com/challenges/largest-rectangle)
  - `problem.md`: 난이도, HackerRank 출처 URL, 문제 설명, 입력/출력 형식, 제약 조건, 예제를 한국어로 포함
  - `explanation.md`: 접근 방식, 핵심 아이디어, 복잡도 분석, 대안적 접근을 한국어로 포함
  - `problem.go`: 핵심 함수 시그니처 + `// 여기에 코드를 작성하세요` 주석 + bufio 기반 main 함수
  - `answer.go`: 완전한 구현 + 단계별 한국어 주석 + bufio I/O + `[알고리즘 힌트]` 함수 주석
  - _요구사항: 6.1, 6.2, 6.3, 12.1, 12.2, 12.3, 12.4, 12.5, 13.1, 13.2, 13.3, 13.4, 13.5, 13.6, 14.1, 14.2, 14.3_

- [x] 8. 스택/큐 문제 10개 생성 (stack-queue/)
  - `data-structures/problems/stack-queue/` 아래에 10개 문제 폴더 생성
  - 각 문제 폴더에 `problem.md`, `explanation.md`, `problem.go`, `answer.go` 4개 파일 생성
  - 문제 목록:
    - `01-easy-maximum-element` (https://www.hackerrank.com/challenges/maximum-element)
    - `02-easy-balanced-brackets` (https://www.hackerrank.com/challenges/balanced-brackets)
    - `03-easy-equal-stacks` (https://www.hackerrank.com/challenges/equal-stacks)
    - `04-easy-queue-using-two-stacks` (https://www.hackerrank.com/challenges/queue-using-two-stacks)
    - `05-medium-simple-text-editor` (https://www.hackerrank.com/challenges/simple-text-editor)
    - `06-medium-waiter` (https://www.hackerrank.com/challenges/waiter)
    - `07-medium-largest-rectangle` (https://www.hackerrank.com/challenges/largest-rectangle)
    - `08-medium-poisonous-plants` (https://www.hackerrank.com/challenges/poisonous-plants)
    - `09-hard-game-of-two-stacks` (https://www.hackerrank.com/challenges/game-of-two-stacks)
    - `10-hard-castle-on-the-grid` (https://www.hackerrank.com/challenges/castle-on-the-grid)
  - `problem.md`: 난이도, HackerRank 출처 URL, 문제 설명, 입력/출력 형식, 제약 조건, 예제를 한국어로 포함
  - `explanation.md`: 접근 방식, 핵심 아이디어, 복잡도 분석, 대안적 접근을 한국어로 포함
  - `problem.go`: 핵심 함수 시그니처 + `// 여기에 코드를 작성하세요` 주석 + bufio 기반 main 함수
  - `answer.go`: 완전한 구현 + 단계별 한국어 주석 + bufio I/O + `[알고리즘 힌트]` 함수 주석
  - _요구사항: 7.1, 7.2, 7.3, 12.1, 12.2, 12.3, 12.4, 12.5, 13.1, 13.2, 13.3, 13.4, 13.5, 13.6, 14.1, 14.2, 14.3_

- [x] 9. 체크포인트 — 자료구조 문제 확인
  - 자료구조 4개 서브카테고리(linked-list, tree, heap, stack-queue)의 40개 문제 폴더가 올바르게 생성되었는지 확인
  - 각 문제 폴더에 4개 파일(problem.md, explanation.md, problem.go, answer.go)이 존재하는지 확인
  - 모든 Go 코드가 컴파일 가능한지 확인
  - 문제가 있으면 사용자에게 질문

- [x] 10. 그리디 문제 10개 생성 (greedy/)
  - `algorithms/problems/greedy/` 아래에 10개 문제 폴더 생성
  - 각 문제 폴더에 `problem.md`, `explanation.md`, `problem.go`, `answer.go` 4개 파일 생성
  - 문제 목록:
    - `01-easy-minimum-absolute-difference-in-an-array` (https://www.hackerrank.com/challenges/minimum-absolute-difference-in-an-array)
    - `02-easy-marc-cakewalk` (https://www.hackerrank.com/challenges/marcs-cakewalk)
    - `03-easy-grid-challenge` (https://www.hackerrank.com/challenges/grid-challenge)
    - `04-easy-luck-balance` (https://www.hackerrank.com/challenges/luck-balance)
    - `05-medium-greedy-florist` (https://www.hackerrank.com/challenges/greedy-florist)
    - `06-medium-max-min` (https://www.hackerrank.com/challenges/angry-children)
    - `07-medium-jim-and-the-orders` (https://www.hackerrank.com/challenges/jim-and-the-orders)
    - `08-medium-permuting-two-arrays` (https://www.hackerrank.com/challenges/two-arrays)
    - `09-hard-chief-hopper` (https://www.hackerrank.com/challenges/chief-hopper)
    - `10-hard-sherlock-and-minimax` (https://www.hackerrank.com/challenges/sherlock-and-minimax)
  - `problem.md`: 난이도, HackerRank 출처 URL, 문제 설명, 입력/출력 형식, 제약 조건, 예제를 한국어로 포함
  - `explanation.md`: 접근 방식, 핵심 아이디어, 복잡도 분석, 대안적 접근을 한국어로 포함
  - `problem.go`: 핵심 함수 시그니처 + `// 여기에 코드를 작성하세요` 주석 + bufio 기반 main 함수
  - `answer.go`: 완전한 구현 + 단계별 한국어 주석 + bufio I/O + `[알고리즘 힌트]` 함수 주석
  - _요구사항: 8.1, 8.2, 8.3, 12.1, 12.2, 12.3, 12.4, 12.5, 13.1, 13.2, 13.3, 13.4, 13.5, 13.6, 14.1, 14.2, 14.3_

- [x] 11. 동적 프로그래밍 문제 10개 생성 (dynamic-programming/)
  - `algorithms/problems/dynamic-programming/` 아래에 10개 문제 폴더 생성
  - 각 문제 폴더에 `problem.md`, `explanation.md`, `problem.go`, `answer.go` 4개 파일 생성
  - 문제 목록:
    - `01-easy-fibonacci-modified` (https://www.hackerrank.com/challenges/fibonacci-modified)
    - `02-easy-the-coin-change-problem` (https://www.hackerrank.com/challenges/coin-change)
    - `03-easy-equal` (https://www.hackerrank.com/challenges/equal)
    - `04-medium-sherlock-and-cost` (https://www.hackerrank.com/challenges/sherlock-and-cost)
    - `05-medium-sam-and-substrings` (https://www.hackerrank.com/challenges/sam-and-substrings)
    - `06-medium-abbreviation` (https://www.hackerrank.com/challenges/abbr)
    - `07-medium-candies` (https://www.hackerrank.com/challenges/candies)
    - `08-medium-the-longest-common-subsequence` (https://www.hackerrank.com/challenges/dynamic-programming-classics-the-longest-common-subsequence)
    - `09-hard-knapsack` (https://www.hackerrank.com/challenges/unbounded-knapsack)
    - `10-hard-construct-the-array` (https://www.hackerrank.com/challenges/construct-the-array)
  - `problem.md`: 난이도, HackerRank 출처 URL, 문제 설명, 입력/출력 형식, 제약 조건, 예제를 한국어로 포함
  - `explanation.md`: 접근 방식, 핵심 아이디어, 복잡도 분석, 대안적 접근을 한국어로 포함
  - `problem.go`: 핵심 함수 시그니처 + `// 여기에 코드를 작성하세요` 주석 + bufio 기반 main 함수
  - `answer.go`: 완전한 구현 + 단계별 한국어 주석 + bufio I/O + `[알고리즘 힌트]` 함수 주석
  - _요구사항: 9.1, 9.2, 9.3, 12.1, 12.2, 12.3, 12.4, 12.5, 13.1, 13.2, 13.3, 13.4, 13.5, 13.6, 14.1, 14.2, 14.3_

- [x] 12. 그래프 문제 10개 생성 (graph/)
  - `algorithms/problems/graph/` 아래에 10개 문제 폴더 생성
  - 각 문제 폴더에 `problem.md`, `explanation.md`, `problem.go`, `answer.go` 4개 파일 생성
  - 문제 목록:
    - `01-easy-bfs-shortest-reach` (https://www.hackerrank.com/challenges/bfsshortreach)
    - `02-easy-dfs-connected-cell-in-a-grid` (https://www.hackerrank.com/challenges/ctci-connected-cell-in-a-grid)
    - `03-easy-roads-and-libraries` (https://www.hackerrank.com/challenges/torque-and-development)
    - `04-medium-journey-to-the-moon` (https://www.hackerrank.com/challenges/journey-to-the-moon)
    - `05-medium-even-tree` (https://www.hackerrank.com/challenges/even-tree)
    - `06-medium-snakes-and-ladders` (https://www.hackerrank.com/challenges/the-quickest-way-up)
    - `07-medium-kruskal-mst-really-special-subtree` (https://www.hackerrank.com/challenges/kruskalmstrsub)
    - `08-medium-dijkstra-shortest-reach-2` (https://www.hackerrank.com/challenges/dijkstrashortreach)
    - `09-hard-prim-mst-special-subtree` (https://www.hackerrank.com/challenges/primsmstsub)
    - `10-hard-jack-goes-to-rapture` (https://www.hackerrank.com/challenges/jack-goes-to-rapture)
  - `problem.md`: 난이도, HackerRank 출처 URL, 문제 설명, 입력/출력 형식, 제약 조건, 예제를 한국어로 포함
  - `explanation.md`: 접근 방식, 핵심 아이디어, 복잡도 분석, 대안적 접근을 한국어로 포함
  - `problem.go`: 핵심 함수 시그니처 + `// 여기에 코드를 작성하세요` 주석 + bufio 기반 main 함수
  - `answer.go`: 완전한 구현 + 단계별 한국어 주석 + bufio I/O + `[알고리즘 힌트]` 함수 주석
  - _요구사항: 10.1, 10.2, 10.3, 12.1, 12.2, 12.3, 12.4, 12.5, 13.1, 13.2, 13.3, 13.4, 13.5, 13.6, 14.1, 14.2, 14.3_

- [x] 13. 탐색 문제 10개 생성 (search/)
  - `algorithms/problems/search/` 아래에 10개 문제 폴더 생성
  - 각 문제 폴더에 `problem.md`, `explanation.md`, `problem.go`, `answer.go` 4개 파일 생성
  - 문제 목록:
    - `01-easy-ice-cream-parlor` (https://www.hackerrank.com/challenges/icecream-parlor)
    - `02-easy-missing-numbers` (https://www.hackerrank.com/challenges/missing-numbers)
    - `03-easy-sherlock-and-array` (https://www.hackerrank.com/challenges/sherlock-and-array)
    - `04-easy-pairs` (https://www.hackerrank.com/challenges/pairs)
    - `05-medium-connected-cells-in-a-grid` (https://www.hackerrank.com/challenges/connected-cell-in-a-grid)
    - `06-medium-count-luck` (https://www.hackerrank.com/challenges/count-luck)
    - `07-medium-hackerland-radio-transmitters` (https://www.hackerrank.com/challenges/hackerland-radio-transmitters)
    - `08-medium-minimum-loss` (https://www.hackerrank.com/challenges/minimum-loss)
    - `09-hard-knightl-on-chessboard` (https://www.hackerrank.com/challenges/knightl-on-chessboard)
    - `10-hard-red-knights-shortest-path` (https://www.hackerrank.com/challenges/red-knights-shortest-path)
  - `problem.md`: 난이도, HackerRank 출처 URL, 문제 설명, 입력/출력 형식, 제약 조건, 예제를 한국어로 포함
  - `explanation.md`: 접근 방식, 핵심 아이디어, 복잡도 분석, 대안적 접근을 한국어로 포함
  - `problem.go`: 핵심 함수 시그니처 + `// 여기에 코드를 작성하세요` 주석 + bufio 기반 main 함수
  - `answer.go`: 완전한 구현 + 단계별 한국어 주석 + bufio I/O + `[알고리즘 힌트]` 함수 주석
  - _요구사항: 11.1, 11.2, 11.3, 12.1, 12.2, 12.3, 12.4, 12.5, 13.1, 13.2, 13.3, 13.4, 13.5, 13.6, 14.1, 14.2, 14.3_

- [x] 14. 최종 체크포인트 — 전체 프로젝트 검증
  - 8개 서브카테고리 × 10문제 = 80개 문제 폴더가 모두 존재하는지 확인
  - 각 문제 폴더에 4개 파일이 모두 존재하는지 확인
  - 모든 Go 코드(`problem.go`, `answer.go`, `examples/*.go`)가 컴파일 가능한지 확인
  - 모든 마크다운 문서가 한국어로 작성되었는지 확인
  - 모든 `problem.md`에 HackerRank 출처 URL이 포함되었는지 확인
  - 문제가 있으면 사용자에게 질문

## 참고사항

- 모든 파일 경로는 `hackerrank-problems/` 루트 폴더 기준이다
- 기존 워크스페이스의 파일 패턴(예: `01-implementation-and-simulation/`)을 참조하여 일관된 스타일을 유지한다
- 각 서브카테고리 태스크(5~8, 10~13)는 독립적이므로 순서 변경 가능하나, 카테고리 구조(2, 3)가 먼저 완료되어야 한다
- 체크포인트에서 Go 코드 컴파일 확인 시 `go build` 또는 `go vet` 사용
