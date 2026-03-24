# 요구사항 문서

## 소개

HackerRank(<https://www.hackerrank.com/>)에서 백준 실버3~골드4 수준에 해당하는 문제를 선별하여, 자료구조와 알고리즘 두 대분류로 나눈 학습 가이드를 구성한다. 기존 번호 접두사(68- 등) 없이 영어 폴더명을 사용하며, 최상위 폴더 `hackerrank-problems/` 아래에 `data-structures/`와 `algorithms/` 하위 폴더를 둔다. 각 하위 카테고리(연결 리스트, 트리, 힙, 스택/큐, 그리디, DP, 그래프, 탐색)마다 약 10문제씩(하/중/상 혼합) 포함하며, 모든 문제는 HackerRank 출처 URL을 명시한다. 기존 스터디 가이드의 파일 패턴(problem.md, explanation.md, 문제파일.go, 정답.go)과 한국어 문서 작성 방식을 그대로 유지한다.

## 용어 정의

- **Study_Guide**: 워크스페이스 전체를 구성하는 Go 알고리즘 스터디 가이드 프로젝트
- **Root_Folder**: 최상위 프로젝트 폴더 `hackerrank-problems/`
- **Category_Folder**: 대분류 폴더 (`data-structures/`, `algorithms/`)
- **Subcategory_Folder**: 세부 주제 폴더 (예: `linked-list/`, `tree/`, `greedy/`, `dynamic-programming/`)
- **Problem_Folder**: 개별 문제를 담는 하위 폴더 (예: `01-easy-xxx/`, `05-medium-xxx/`, `09-hard-xxx/`)
- **Solution_Template**: 사용자가 직접 풀어볼 수 있도록 핵심 함수 본문이 비어 있는 Go 소스 파일 (`문제파일.go`)
- **Answer_File**: 상세 주석이 포함된 완전한 정답 Go 소스 파일 (`정답.go`)
- **Problem_Statement**: 문제 설명, 입출력 형식, 제약 조건, 예제를 포함하는 마크다운 파일 (`problem.md`)
- **Explanation_File**: 풀이 접근 방식, 핵심 아이디어, 복잡도 분석을 포함하는 마크다운 파일 (`explanation.md`)
- **HackerRank_Source**: 문제의 원본 출처인 HackerRank URL
- **Difficulty_Level**: 문제 난이도 등급 (하/중/상), 백준 실버3~골드4 범위에 대응

## 요구사항

### 요구사항 1: 최상위 폴더 구조 생성

**사용자 스토리:** 알고리즘 학습자로서, 자료구조와 알고리즘을 체계적으로 분류한 학습 가이드를 갖고 싶다. 이를 통해 주제별로 효율적으로 학습할 수 있다.

#### 인수 조건 1

1. THE Study_Guide SHALL 번호 접두사 없이 `hackerrank-problems/` Root_Folder를 생성한다
2. THE Root_Folder SHALL `README.md` 파일을 포함하며, 전체 프로젝트 개요, 카테고리 구성 안내, 전체 문제 수, 난이도 분포를 한국어로 포함한다
3. THE Root_Folder SHALL `data-structures/` Category_Folder를 포함한다
4. THE Root_Folder SHALL `algorithms/` Category_Folder를 포함한다

### 요구사항 2: 자료구조 카테고리 구조 생성

**사용자 스토리:** 알고리즘 학습자로서, 자료구조를 세부 주제별로 나누어 학습하고 싶다. 이를 통해 각 자료구조에 집중하여 깊이 있는 학습이 가능하다.

#### 인수 조건 2

1. THE `data-structures/` Category_Folder SHALL `README.md`, `theory.md`, `examples/`, `problems/` 구조를 포함한다
2. THE `data-structures/theory.md` SHALL 연결 리스트, 트리, 힙, 스택/큐 각각의 개념, 동작 원리, 시간/공간 복잡도 비교, 실전 팁, 자료구조 선택 기준을 한국어로 포함한다
3. THE `data-structures/examples/` SHALL 연결 리스트(`linked_list.go`), 트리(`tree.go`), 힙(`heap.go`)의 기본 동작을 보여주는 Go 예시 코드를 포함한다
4. THE `data-structures/problems/` SHALL 다음 4개의 Subcategory_Folder를 포함한다: `linked-list/`, `tree/`, `heap/`, `stack-queue/`

### 요구사항 3: 알고리즘 카테고리 구조 생성

**사용자 스토리:** 알고리즘 학습자로서, 코딩 테스트에서 자주 출제되는 알고리즘 유형별로 문제를 학습하고 싶다. 이를 통해 실전 코딩 테스트에 효과적으로 대비할 수 있다.

#### 인수 조건 3

1. THE `algorithms/` Category_Folder SHALL `README.md`, `theory.md`, `examples/`, `problems/` 구조를 포함한다
2. THE `algorithms/theory.md` SHALL 그리디, 동적 프로그래밍, 그래프, 탐색 각각의 개념, 동작 원리, 시간/공간 복잡도 비교, 실전 팁, 알고리즘 선택 기준을 한국어로 포함한다
3. THE `algorithms/examples/` SHALL 그리디(`greedy.go`), 동적 프로그래밍(`dp.go`), 그래프(`graph.go`), 탐색(`search.go`)의 기본 동작을 보여주는 Go 예시 코드를 포함한다
4. THE `algorithms/problems/` SHALL 다음 4개의 Subcategory_Folder를 포함한다: `greedy/`, `dynamic-programming/`, `graph/`, `search/`

### 요구사항 4: HackerRank 문제 선별 — 자료구조 (연결 리스트)

**사용자 스토리:** 알고리즘 학습자로서, HackerRank에서 연결 리스트 관련 문제 10개를 난이도 혼합으로 학습하고 싶다.

#### 인수 조건 4

1. THE `linked-list/` Subcategory_Folder SHALL 다음 10개의 HackerRank 문제를 포함한다:
   - `01-easy-print-the-elements-of-a-linked-list`: "Print the Elements of a Linked List" (<https://www.hackerrank.com/challenges/print-the-elements-of-a-linked-list>)
   - `02-easy-insert-a-node-at-the-tail`: "Insert a Node at the Tail of a Linked List" (<https://www.hackerrank.com/challenges/insert-a-node-at-the-tail-of-a-linked-list>)
   - `03-easy-insert-a-node-at-the-head`: "Insert a node at the head of a linked list" (<https://www.hackerrank.com/challenges/insert-a-node-at-the-head-of-a-linked-list>)
   - `04-easy-delete-a-node`: "Delete a Node" (<https://www.hackerrank.com/challenges/delete-a-node-from-a-linked-list>)
   - `05-easy-reverse-a-linked-list`: "Reverse a linked list" (<https://www.hackerrank.com/challenges/reverse-a-linked-list>)
   - `06-medium-compare-two-linked-lists`: "Compare two linked lists" (<https://www.hackerrank.com/challenges/compare-two-linked-lists>)
   - `07-medium-merge-two-sorted-linked-lists`: "Merge two sorted linked lists" (<https://www.hackerrank.com/challenges/merge-two-sorted-linked-lists>)
   - `08-medium-get-node-value`: "Get Node Value" (<https://www.hackerrank.com/challenges/get-the-value-of-the-node-at-a-specific-position-from-the-tail>)
   - `09-medium-insert-a-node-at-a-specific-position`: "Insert a node at a specific position in a linked list" (<https://www.hackerrank.com/challenges/insert-a-node-at-a-specific-position-in-a-linked-list>)
   - `10-hard-reverse-a-doubly-linked-list`: "Reverse a doubly linked list" (<https://www.hackerrank.com/challenges/reverse-a-doubly-linked-list>)
2. WHEN Problem_Statement를 작성할 때, THE Problem_Statement SHALL HackerRank_Source URL을 명시한다
3. THE 문제 목록 SHALL 하(Easy), 중(Medium), 상(Hard) 난이도를 혼합하여 구성한다

### 요구사항 5: HackerRank 문제 선별 — 자료구조 (트리)

**사용자 스토리:** 알고리즘 학습자로서, HackerRank에서 트리 관련 문제 10개를 난이도 혼합으로 학습하고 싶다.

#### 인수 조건 5

1. THE `tree/` Subcategory_Folder SHALL 다음 10개의 HackerRank 문제를 포함한다:
   - `01-easy-tree-preorder-traversal`: "Tree: Preorder Traversal" (<https://www.hackerrank.com/challenges/tree-preorder-traversal>)
   - `02-easy-tree-postorder-traversal`: "Tree: Postorder Traversal" (<https://www.hackerrank.com/challenges/tree-postorder-traversal>)
   - `03-easy-tree-inorder-traversal`: "Tree: Inorder Traversal" (<https://www.hackerrank.com/challenges/tree-inorder-traversal>)
   - `04-easy-tree-height-of-a-binary-tree`: "Tree: Height of a Binary Tree" (<https://www.hackerrank.com/challenges/tree-height-of-a-binary-tree>)
   - `05-easy-tree-level-order-traversal`: "Tree: Level Order Traversal" (<https://www.hackerrank.com/challenges/tree-level-order-traversal>)
   - `06-medium-tree-top-view`: "Tree: Top View" (<https://www.hackerrank.com/challenges/tree-top-view>)
   - `07-medium-binary-search-tree-insertion`: "Binary Search Tree: Insertion" (<https://www.hackerrank.com/challenges/binary-search-tree-insertion>)
   - `08-medium-binary-search-tree-lowest-common-ancestor`: "Binary Search Tree: Lowest Common Ancestor" (<https://www.hackerrank.com/challenges/binary-search-tree-lowest-common-ancestor>)
   - `09-medium-is-this-a-binary-search-tree`: "Is This a Binary Search Tree?" (<https://www.hackerrank.com/challenges/is-binary-search-tree>)
   - `10-hard-swap-nodes-algo`: "Swap Nodes [Algo]" (<https://www.hackerrank.com/challenges/swap-nodes-algo>)
2. WHEN Problem_Statement를 작성할 때, THE Problem_Statement SHALL HackerRank_Source URL을 명시한다
3. THE 문제 목록 SHALL 하(Easy), 중(Medium), 상(Hard) 난이도를 혼합하여 구성한다

### 요구사항 6: HackerRank 문제 선별 — 자료구조 (힙)

**사용자 스토리:** 알고리즘 학습자로서, HackerRank에서 힙(우선순위 큐) 관련 문제 10개를 난이도 혼합으로 학습하고 싶다.

#### 인수 조건 6

1. THE `heap/` Subcategory_Folder SHALL 다음 10개의 HackerRank 문제를 포함한다:
   - `01-easy-jesse-and-cookies`: "Jesse and Cookies" (<https://www.hackerrank.com/challenges/jesse-and-cookies>)
   - `02-easy-qheap1`: "QHEAP1" (<https://www.hackerrank.com/challenges/qheap1>)
   - `03-easy-minimum-average-waiting-time`: "Minimum Average Waiting Time" (<https://www.hackerrank.com/challenges/minimum-average-waiting-time>)
   - `04-medium-find-the-running-median`: "Find the Running Median" (<https://www.hackerrank.com/challenges/find-the-running-median>)
   - `05-medium-components-in-a-graph`: "Components in a graph" (<https://www.hackerrank.com/challenges/components-in-graph>)
   - `06-medium-kundu-and-tree`: "Kundu and Tree" (<https://www.hackerrank.com/challenges/kundu-and-tree>)
   - `07-medium-heap-full-sort`: "Heap Sort — Full Sort" (<https://www.hackerrank.com/challenges/heapsort>)
   - `08-hard-median-updates`: "Median Updates" (<https://www.hackerrank.com/challenges/median>)
   - `09-hard-kth-minimum`: "Kth Minimum in Range" (<https://www.hackerrank.com/challenges/kth-minimum-in-range>)
   - `10-hard-largest-rectangle`: "Largest Rectangle" (<https://www.hackerrank.com/challenges/largest-rectangle>)
2. WHEN Problem_Statement를 작성할 때, THE Problem_Statement SHALL HackerRank_Source URL을 명시한다
3. THE 문제 목록 SHALL 하(Easy), 중(Medium), 상(Hard) 난이도를 혼합하여 구성한다

### 요구사항 7: HackerRank 문제 선별 — 자료구조 (스택/큐)

**사용자 스토리:** 알고리즘 학습자로서, HackerRank에서 스택과 큐 관련 문제 10개를 난이도 혼합으로 학습하고 싶다.

#### 인수 조건 7

1. THE `stack-queue/` Subcategory_Folder SHALL 다음 10개의 HackerRank 문제를 포함한다:
   - `01-easy-maximum-element`: "Maximum Element" (<https://www.hackerrank.com/challenges/maximum-element>)
   - `02-easy-balanced-brackets`: "Balanced Brackets" (<https://www.hackerrank.com/challenges/balanced-brackets>)
   - `03-easy-equal-stacks`: "Equal Stacks" (<https://www.hackerrank.com/challenges/equal-stacks>)
   - `04-easy-queue-using-two-stacks`: "Queue using Two Stacks" (<https://www.hackerrank.com/challenges/queue-using-two-stacks>)
   - `05-medium-simple-text-editor`: "Simple Text Editor" (<https://www.hackerrank.com/challenges/simple-text-editor>)
   - `06-medium-waiter`: "Waiter" (<https://www.hackerrank.com/challenges/waiter>)
   - `07-medium-largest-rectangle`: "Largest Rectangle" (<https://www.hackerrank.com/challenges/largest-rectangle>)
   - `08-medium-poisonous-plants`: "Poisonous Plants" (<https://www.hackerrank.com/challenges/poisonous-plants>)
   - `09-hard-game-of-two-stacks`: "Game of Two Stacks" (<https://www.hackerrank.com/challenges/game-of-two-stacks>)
   - `10-hard-castle-on-the-grid`: "Castle on the Grid" (<https://www.hackerrank.com/challenges/castle-on-the-grid>)
2. WHEN Problem_Statement를 작성할 때, THE Problem_Statement SHALL HackerRank_Source URL을 명시한다
3. THE 문제 목록 SHALL 하(Easy), 중(Medium), 상(Hard) 난이도를 혼합하여 구성한다

### 요구사항 8: HackerRank 문제 선별 — 알고리즘 (그리디)

**사용자 스토리:** 알고리즘 학습자로서, HackerRank에서 그리디 알고리즘 관련 문제 10개를 난이도 혼합으로 학습하고 싶다.

#### 인수 조건 8

1. THE `greedy/` Subcategory_Folder SHALL 다음 10개의 HackerRank 문제를 포함한다:
   - `01-easy-minimum-absolute-difference-in-an-array`: "Minimum Absolute Difference in an Array" (<https://www.hackerrank.com/challenges/minimum-absolute-difference-in-an-array>)
   - `02-easy-marc-cakewalk`: "Marc's Cakewalk" (<https://www.hackerrank.com/challenges/marcs-cakewalk>)
   - `03-easy-grid-challenge`: "Grid Challenge" (<https://www.hackerrank.com/challenges/grid-challenge>)
   - `04-easy-luck-balance`: "Luck Balance" (<https://www.hackerrank.com/challenges/luck-balance>)
   - `05-medium-greedy-florist`: "Greedy Florist" (<https://www.hackerrank.com/challenges/greedy-florist>)
   - `06-medium-max-min`: "Max Min" (<https://www.hackerrank.com/challenges/angry-children>)
   - `07-medium-jim-and-the-orders`: "Jim and the Orders" (<https://www.hackerrank.com/challenges/jim-and-the-orders>)
   - `08-medium-permuting-two-arrays`: "Permuting Two Arrays" (<https://www.hackerrank.com/challenges/two-arrays>)
   - `09-hard-chief-hopper`: "Chief Hopper" (<https://www.hackerrank.com/challenges/chief-hopper>)
   - `10-hard-sherlock-and-minimax`: "Sherlock and MiniMax" (<https://www.hackerrank.com/challenges/sherlock-and-minimax>)
2. WHEN Problem_Statement를 작성할 때, THE Problem_Statement SHALL HackerRank_Source URL을 명시한다
3. THE 문제 목록 SHALL 하(Easy), 중(Medium), 상(Hard) 난이도를 혼합하여 구성한다

### 요구사항 9: HackerRank 문제 선별 — 알고리즘 (동적 프로그래밍)

**사용자 스토리:** 알고리즘 학습자로서, HackerRank에서 동적 프로그래밍 관련 문제 10개를 난이도 혼합으로 학습하고 싶다.

#### 인수 조건 9

1. THE `dynamic-programming/` Subcategory_Folder SHALL 다음 10개의 HackerRank 문제를 포함한다:
   - `01-easy-fibonacci-modified`: "Fibonacci Modified" (<https://www.hackerrank.com/challenges/fibonacci-modified>)
   - `02-easy-the-coin-change-problem`: "The Coin Change Problem" (<https://www.hackerrank.com/challenges/coin-change>)
   - `03-easy-equal`: "Equal" (<https://www.hackerrank.com/challenges/equal>)
   - `04-medium-sherlock-and-cost`: "Sherlock and Cost" (<https://www.hackerrank.com/challenges/sherlock-and-cost>)
   - `05-medium-sam-and-substrings`: "Sam and substrings" (<https://www.hackerrank.com/challenges/sam-and-substrings>)
   - `06-medium-abbreviation`: "Abbreviation" (<https://www.hackerrank.com/challenges/abbr>)
   - `07-medium-candies`: "Candies" (<https://www.hackerrank.com/challenges/candies>)
   - `08-medium-the-longest-common-subsequence`: "The Longest Common Subsequence" (<https://www.hackerrank.com/challenges/dynamic-programming-classics-the-longest-common-subsequence>)
   - `09-hard-knapsack`: "Knapsack" (<https://www.hackerrank.com/challenges/unbounded-knapsack>)
   - `10-hard-construct-the-array`: "Construct the Array" (<https://www.hackerrank.com/challenges/construct-the-array>)
2. WHEN Problem_Statement를 작성할 때, THE Problem_Statement SHALL HackerRank_Source URL을 명시한다
3. THE 문제 목록 SHALL 하(Easy), 중(Medium), 상(Hard) 난이도를 혼합하여 구성한다

### 요구사항 10: HackerRank 문제 선별 — 알고리즘 (그래프)

**사용자 스토리:** 알고리즘 학습자로서, HackerRank에서 그래프 관련 문제 10개를 난이도 혼합으로 학습하고 싶다.

#### 인수 조건 10

1. THE `graph/` Subcategory_Folder SHALL 다음 10개의 HackerRank 문제를 포함한다:
   - `01-easy-bfs-shortest-reach`: "BFS: Shortest Reach in a Graph" (<https://www.hackerrank.com/challenges/bfsshortreach>)
   - `02-easy-dfs-connected-cell-in-a-grid`: "DFS: Connected Cell in a Grid" (<https://www.hackerrank.com/challenges/ctci-connected-cell-in-a-grid>)
   - `03-easy-roads-and-libraries`: "Roads and Libraries" (<https://www.hackerrank.com/challenges/torque-and-development>)
   - `04-medium-journey-to-the-moon`: "Journey to the Moon" (<https://www.hackerrank.com/challenges/journey-to-the-moon>)
   - `05-medium-even-tree`: "Even Tree" (<https://www.hackerrank.com/challenges/even-tree>)
   - `06-medium-snakes-and-ladders`: "Snakes and Ladders: The Quickest Way Up" (<https://www.hackerrank.com/challenges/the-quickest-way-up>)
   - `07-medium-kruskal-mst-really-special-subtree`: "Kruskal (MST): Really Special Subtree" (<https://www.hackerrank.com/challenges/kruskalmstrsub>)
   - `08-medium-dijkstra-shortest-reach-2`: "Dijkstra: Shortest Reach 2" (<https://www.hackerrank.com/challenges/dijkstrashortreach>)
   - `09-hard-prim-mst-special-subtree`: "Prim's (MST): Special Subtree" (<https://www.hackerrank.com/challenges/primsmstsub>)
   - `10-hard-jack-goes-to-rapture`: "Jack goes to Rapture" (<https://www.hackerrank.com/challenges/jack-goes-to-rapture>)
2. WHEN Problem_Statement를 작성할 때, THE Problem_Statement SHALL HackerRank_Source URL을 명시한다
3. THE 문제 목록 SHALL 하(Easy), 중(Medium), 상(Hard) 난이도를 혼합하여 구성한다

### 요구사항 11: HackerRank 문제 선별 — 알고리즘 (탐색)

**사용자 스토리:** 알고리즘 학습자로서, HackerRank에서 탐색(이진 탐색, 완전 탐색 등) 관련 문제 10개를 난이도 혼합으로 학습하고 싶다.

#### 인수 조건 11

1. THE `search/` Subcategory_Folder SHALL 다음 10개의 HackerRank 문제를 포함한다:
   - `01-easy-ice-cream-parlor`: "Ice Cream Parlor" (<https://www.hackerrank.com/challenges/icecream-parlor>)
   - `02-easy-missing-numbers`: "Missing Numbers" (<https://www.hackerrank.com/challenges/missing-numbers>)
   - `03-easy-sherlock-and-array`: "Sherlock and Array" (<https://www.hackerrank.com/challenges/sherlock-and-array>)
   - `04-easy-pairs`: "Pairs" (<https://www.hackerrank.com/challenges/pairs>)
   - `05-medium-connected-cells-in-a-grid`: "Connected Cells in a Grid" (<https://www.hackerrank.com/challenges/connected-cell-in-a-grid>)
   - `06-medium-count-luck`: "Count Luck" (<https://www.hackerrank.com/challenges/count-luck>)
   - `07-medium-hackerland-radio-transmitters`: "Hackerland Radio Transmitters" (<https://www.hackerrank.com/challenges/hackerland-radio-transmitters>)
   - `08-medium-minimum-loss`: "Minimum Loss" (<https://www.hackerrank.com/challenges/minimum-loss>)
   - `09-hard-knightl-on-chessboard`: "KnightL on a Chessboard" (<https://www.hackerrank.com/challenges/knightl-on-chessboard>)
   - `10-hard-red-knights-shortest-path`: "Red Knight's Shortest Path" (<https://www.hackerrank.com/challenges/red-knights-shortest-path>)
2. WHEN Problem_Statement를 작성할 때, THE Problem_Statement SHALL HackerRank_Source URL을 명시한다
3. THE 문제 목록 SHALL 하(Easy), 중(Medium), 상(Hard) 난이도를 혼합하여 구성한다

### 요구사항 12: 문제 폴더 구조 생성

**사용자 스토리:** 알고리즘 학습자로서, 각 문제가 기존 문제 폴더와 동일한 파일 구조를 갖기를 원한다. 이를 통해 일관된 방식으로 문제를 풀고 학습할 수 있다.

#### 인수 조건 12

1. THE Problem_Folder SHALL `problem.md`, `explanation.md`, `문제파일.go`, `정답.go` 4개 파일을 포함한다
2. THE Problem_Folder SHALL `{번호}-{난이도}-{영문 문제명 kebab-case}` 형식으로 명명한다 (예: `01-easy-print-the-elements-of-a-linked-list`)
3. THE Problem_Statement SHALL 난이도, 출처 URL, 문제 설명, 입력 형식, 출력 형식, 제약 조건, 예제를 한국어로 포함한다
4. THE Explanation_File SHALL 접근 방식, 핵심 아이디어, 복잡도 분석, 대안적 접근을 한국어로 포함한다
5. THE 각 Subcategory_Folder 내 문제 번호 SHALL 01부터 10까지 순차적으로 매기되, 난이도는 하/중/상을 혼합하여 배치한다

### 요구사항 13: 문제파일(문제파일.go) 및 정답(정답.go) 파일 생성

**사용자 스토리:** 알고리즘 학습자로서, 빈 템플릿으로 직접 풀어보고 정답과 비교하고 싶다. 이를 통해 능동적으로 학습할 수 있다.

#### 인수 조건 13

1. THE Solution_Template SHALL 핵심 함수의 시그니처와 매개변수/반환값 주석을 포함하되, 함수 본문은 비워둔다
2. THE Solution_Template SHALL `// 여기에 코드를 작성하세요` 주석을 함수 본문에 포함한다
3. THE Solution_Template SHALL 입출력 처리 코드(main 함수)를 완전하게 포함하여 사용자가 핵심 로직만 작성하면 실행 가능하도록 한다
4. THE Answer_File SHALL 핵심 함수의 완전한 구현과 각 단계별 상세 주석을 포함한다
5. THE Answer_File SHALL `bufio.NewReader`와 `bufio.NewWriter`를 사용하여 효율적인 입출력을 처리한다
6. THE Answer_File SHALL 알고리즘 힌트를 함수 주석에 포함한다

### 요구사항 14: 한국어 문서 작성

**사용자 스토리:** 한국어 사용자로서, 모든 문서가 한국어로 작성되기를 원한다. 이를 통해 언어 장벽 없이 학습할 수 있다.

#### 인수 조건 14

1. THE Study_Guide SHALL 모든 마크다운 문서(README.md, theory.md, problem.md, explanation.md)를 한국어로 작성한다
2. THE Study_Guide SHALL Go 소스 코드의 주석을 한국어로 작성한다
3. THE Study_Guide SHALL 기존 문서의 문체와 용어 사용 방식을 일관되게 유지한다

### 요구사항 15: 예시 코드 작성

**사용자 스토리:** 알고리즘 학습자로서, 각 카테고리의 기본 동작을 보여주는 예시 코드를 통해 개념을 이해하고 싶다.

#### 인수 조건 15

1. THE `data-structures/examples/` SHALL 연결 리스트의 기본 연산(삽입, 삭제, 탐색, 출력)을 보여주는 `linked_list.go`를 포함한다
2. THE `data-structures/examples/` SHALL 이진 트리의 기본 연산(삽입, 순회)을 보여주는 `tree.go`를 포함한다
3. THE `data-structures/examples/` SHALL 힙(우선순위 큐)의 기본 연산(삽입, 추출)을 보여주는 `heap.go`를 포함한다
4. THE `algorithms/examples/` SHALL 그리디 알고리즘의 기본 패턴을 보여주는 `greedy.go`를 포함한다
5. THE `algorithms/examples/` SHALL 동적 프로그래밍의 기본 패턴(메모이제이션, 타뷸레이션)을 보여주는 `dp.go`를 포함한다
6. THE `algorithms/examples/` SHALL 그래프 탐색(BFS, DFS)의 기본 패턴을 보여주는 `graph.go`를 포함한다
7. THE `algorithms/examples/` SHALL 이진 탐색의 기본 패턴을 보여주는 `search.go`를 포함한다
8. THE 모든 예시 코드 SHALL 각 함수에 한국어 주석으로 동작 설명을 포함한다
9. THE 모든 예시 코드 SHALL `main` 함수에서 예시 실행 결과를 출력한다
10. THE 모든 예시 코드 SHALL 시간 복잡도와 공간 복잡도를 파일 상단 주석에 명시한다
