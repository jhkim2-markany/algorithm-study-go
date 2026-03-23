# 구현 계획: 코딩테스트 학습용 알고리즘 자료

## 개요

67개 알고리즘 유형의 학습 자료를 체계적으로 생성하는 정적 콘텐츠 저장소를 구축한다. 루트 README.md와 67개 알고리즘 폴더(각각 이론 문서, 예시 코드, 3개 문제/풀이/해설 포함)를 순서대로 생성하고, 검증 스크립트로 구조적 정확성을 확인한다. 기존 53개 알고리즘(01~53번)은 완료되었으며, 이론 문서 상세 보충(01~53번), Tier_1 6개(54~59번), Tier_2 8개(60~67번) 알고리즘을 추가한다.

## Tasks

- [x] 1. 프로젝트 기본 구조 설정
  - [x] 1.1 루트 README.md 생성
    - 프로젝트 목적, 대상 독자, 26개 알고리즘 유형 목록(폴더 링크 포함), 권장 학습 순서, Golang 개발 환경 설정 안내를 포함
    - 한국어로 작성
    - _Requirements: 6.1, 6.2, 6.3, 6.4, 6.5, 6.6_

  - [x] 1.2 26개 알고리즘 폴더 및 하위 디렉토리 생성
    - 01-implementation-and-simulation부터 26-string-algorithm까지 26개 폴더 생성
    - 각 폴더에 examples/, problems/ 하위 디렉토리 생성
    - 각 problems/ 내에 01-easy-*, 02-medium-*, 03-hard-* 문제 폴더 생성
    - _Requirements: 1.1, 1.2, 1.3, 1.4_

- [x] 2. 01-구현과 시뮬레이션 콘텐츠 생성
  - [x] 2.1 01-implementation-and-simulation/README.md 생성
    - 알고리즘 유형 이름, 설명, 폴더 내 파일 목록, 문제 목록(난이도 표기) 포함
    - _Requirements: 1.5_
  - [x] 2.2 01-implementation-and-simulation/theory.md 생성
    - 개념 정의, 동작 원리, 시간/공간 복잡도, 적합한 문제 유형 포함
    - 한국어로 작성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 2.3 01-implementation-and-simulation/examples/ 예시 코드 생성
    - package main, main() 함수, 한국어 주석, 표준 라이브러리만 사용
    - go run으로 실행 가능한 완전한 코드
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 2.4 01-implementation-and-simulation/problems/ 3개 문제 생성
    - 01-easy, 02-medium, 03-hard 각각 problem.md, solution.go, explanation.md 생성
    - problem.md: 문제 설명, 입출력 형식, 예제, 제약 조건, 난이도 표기
    - solution.go: 컴파일 가능한 Golang 풀이, 한국어 주석, 표준 입출력
    - explanation.md: 접근 방식, 핵심 아이디어, 복잡도 분석
    - _Requirements: 4.1, 4.2, 4.3, 4.4, 4.5, 4.6, 4.7, 4.8, 4.9, 5.1, 5.2, 5.3, 5.4, 5.5_

- [x] 3. 02-브루트포스 콘텐츠 생성
  - [x] 3.1 02-bruteforce/README.md 생성
    - _Requirements: 1.5_
  - [x] 3.2 02-bruteforce/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 3.3 02-bruteforce/examples/ 예시 코드 생성
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 3.4 02-bruteforce/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 4. 03-정렬 콘텐츠 생성
  - [x] 4.1 03-sorting/README.md 생성
    - _Requirements: 1.5_
  - [x] 4.2 03-sorting/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 4.3 03-sorting/examples/ 예시 코드 생성
    - 여러 정렬 변형을 별도 파일로 제공 (예: bubble_sort.go, merge_sort.go, quick_sort.go)
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5, 3.6_
  - [x] 4.4 03-sorting/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 5. 04-스택과 큐 콘텐츠 생성
  - [x] 5.1 04-stack-and-queue/README.md 생성
    - _Requirements: 1.5_
  - [x] 5.2 04-stack-and-queue/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 5.3 04-stack-and-queue/examples/ 예시 코드 생성
    - 스택과 큐 각각 별도 파일로 제공
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5, 3.6_
  - [x] 5.4 04-stack-and-queue/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 6. 05-해시 콘텐츠 생성
  - [x] 6.1 05-hash/README.md 생성
    - _Requirements: 1.5_
  - [x] 6.2 05-hash/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 6.3 05-hash/examples/ 예시 코드 생성
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 6.4 05-hash/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 7. 06-누적합 콘텐츠 생성
  - [x] 7.1 06-prefix-sum/README.md 생성
    - _Requirements: 1.5_
  - [x] 7.2 06-prefix-sum/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 7.3 06-prefix-sum/examples/ 예시 코드 생성
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 7.4 06-prefix-sum/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 8. 07-수학과 정수론 콘텐츠 생성
  - [x] 8.1 07-math-and-number-theory/README.md 생성
    - _Requirements: 1.5_
  - [x] 8.2 07-math-and-number-theory/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 8.3 07-math-and-number-theory/examples/ 예시 코드 생성
    - GCD, 소수 판별 등 주요 변형을 별도 파일로 제공
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5, 3.6_
  - [x] 8.4 07-math-and-number-theory/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 9. Checkpoint - 01~07번 알고리즘 콘텐츠 검증
  - Ensure all .go files compile without errors, ask the user if questions arise.

- [x] 10. 08-이진 탐색 콘텐츠 생성
  - [x] 10.1 08-binary-search/README.md 생성
    - _Requirements: 1.5_
  - [x] 10.2 08-binary-search/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 10.3 08-binary-search/examples/ 예시 코드 생성
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 10.4 08-binary-search/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 11. 09-파라메트릭 서치 콘텐츠 생성
  - [x] 11.1 09-parametric-search/README.md 생성
    - _Requirements: 1.5_
  - [x] 11.2 09-parametric-search/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 11.3 09-parametric-search/examples/ 예시 코드 생성
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 11.4 09-parametric-search/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 12. 10-투 포인터와 슬라이딩 윈도우 콘텐츠 생성
  - [x] 12.1 10-two-pointer-and-sliding-window/README.md 생성
    - _Requirements: 1.5_
  - [x] 12.2 10-two-pointer-and-sliding-window/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 12.3 10-two-pointer-and-sliding-window/examples/ 예시 코드 생성
    - 투 포인터와 슬라이딩 윈도우 각각 별도 파일로 제공
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5, 3.6_
  - [x] 12.4 10-two-pointer-and-sliding-window/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 13. 11-그리디 콘텐츠 생성
  - [x] 13.1 11-greedy/README.md 생성
    - _Requirements: 1.5_
  - [x] 13.2 11-greedy/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 13.3 11-greedy/examples/ 예시 코드 생성
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 13.4 11-greedy/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 14. 12-힙과 우선순위 큐 콘텐츠 생성
  - [x] 14.1 12-heap-and-priority-queue/README.md 생성
    - _Requirements: 1.5_
  - [x] 14.2 12-heap-and-priority-queue/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 14.3 12-heap-and-priority-queue/examples/ 예시 코드 생성
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 14.4 12-heap-and-priority-queue/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 15. Checkpoint - 08~12번 알고리즘 콘텐츠 검증
  - Ensure all .go files compile without errors, ask the user if questions arise.

- [x] 16. 13-트리 콘텐츠 생성
  - [x] 16.1 13-tree/README.md 생성
    - _Requirements: 1.5_
  - [x] 16.2 13-tree/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 16.3 13-tree/examples/ 예시 코드 생성
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 16.4 13-tree/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 17. 14-이진 트리 콘텐츠 생성
  - [x] 17.1 14-binary-tree/README.md 생성
    - _Requirements: 1.5_
  - [x] 17.2 14-binary-tree/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 17.3 14-binary-tree/examples/ 예시 코드 생성
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 17.4 14-binary-tree/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 18. 15-그래프 탐색 DFS 콘텐츠 생성
  - [x] 18.1 15-graph-dfs/README.md 생성
    - _Requirements: 1.5_
  - [x] 18.2 15-graph-dfs/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 18.3 15-graph-dfs/examples/ 예시 코드 생성
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 18.4 15-graph-dfs/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 19. 16-그래프 탐색 BFS 콘텐츠 생성
  - [x] 19.1 16-graph-bfs/README.md 생성
    - _Requirements: 1.5_
  - [x] 19.2 16-graph-bfs/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 19.3 16-graph-bfs/examples/ 예시 코드 생성
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 19.4 16-graph-bfs/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 20. 17-백트래킹 콘텐츠 생성
  - [x] 20.1 17-backtracking/README.md 생성
    - _Requirements: 1.5_
  - [x] 20.2 17-backtracking/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 20.3 17-backtracking/examples/ 예시 코드 생성
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 20.4 17-backtracking/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 21. 18-분할 정복 콘텐츠 생성
  - [x] 21.1 18-divide-and-conquer/README.md 생성
    - _Requirements: 1.5_
  - [x] 21.2 18-divide-and-conquer/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 21.3 18-divide-and-conquer/examples/ 예시 코드 생성
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 21.4 18-divide-and-conquer/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 22. Checkpoint - 13~18번 알고리즘 콘텐츠 검증
  - Ensure all .go files compile without errors, ask the user if questions arise.

- [x] 23. 19-동적 프로그래밍 콘텐츠 생성
  - [x] 23.1 19-dynamic-programming/README.md 생성
    - _Requirements: 1.5_
  - [x] 23.2 19-dynamic-programming/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 23.3 19-dynamic-programming/examples/ 예시 코드 생성
    - Top-down, Bottom-up 등 주요 변형을 별도 파일로 제공
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5, 3.6_
  - [x] 23.4 19-dynamic-programming/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 24. 20-유니온 파인드 콘텐츠 생성
  - [x] 24.1 20-union-find/README.md 생성
    - _Requirements: 1.5_
  - [x] 24.2 20-union-find/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 24.3 20-union-find/examples/ 예시 코드 생성
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 24.4 20-union-find/problems/ 3개 문제 생성

- [x] 25. 21-최단 경로 콘텐츠 생성
  - [x] 25.1 21-shortest-path/README.md 생성
    - _Requirements: 1.5_
  - [x] 25.2 21-shortest-path/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 25.3 21-shortest-path/examples/ 예시 코드 생성
    - Dijkstra, Bellman-Ford, Floyd-Warshall 등 주요 변형을 별도 파일로 제공
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5, 3.6_
  - [x] 25.4 21-shortest-path/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 26. 22-최소 신장 트리 콘텐츠 생성
  - [x] 26.1 22-minimum-spanning-tree/README.md 생성
    - _Requirements: 1.5_
  - [x] 26.2 22-minimum-spanning-tree/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 26.3 22-minimum-spanning-tree/examples/ 예시 코드 생성
    - Kruskal, Prim 등 주요 변형을 별도 파일로 제공
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5, 3.6_
  - [x] 26.4 22-minimum-spanning-tree/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 27. 23-위상 정렬 콘텐츠 생성
  - [x] 27.1 23-topological-sort/README.md 생성
    - _Requirements: 1.5_
  - [x] 27.2 23-topological-sort/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 27.3 23-topological-sort/examples/ 예시 코드 생성
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 27.4 23-topological-sort/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 28. 24-그래프 알고리즘 기타 콘텐츠 생성
  - [x] 28.1 24-graph-advanced/README.md 생성
    - _Requirements: 1.5_
  - [x] 28.2 24-graph-advanced/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 28.3 24-graph-advanced/examples/ 예시 코드 생성
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 28.4 24-graph-advanced/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 29. 25-세그먼트 트리 콘텐츠 생성
  - [x] 29.1 25-segment-tree/README.md 생성
    - _Requirements: 1.5_
  - [x] 29.2 25-segment-tree/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 29.3 25-segment-tree/examples/ 예시 코드 생성
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 29.4 25-segment-tree/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 30. 26-문자열 알고리즘 콘텐츠 생성
  - [x] 30.1 26-string-algorithm/README.md 생성
    - _Requirements: 1.5_
  - [x] 30.2 26-string-algorithm/theory.md 생성
    - _Requirements: 2.1, 2.2, 2.3, 2.4, 2.5, 2.6_
  - [x] 30.3 26-string-algorithm/examples/ 예시 코드 생성
    - _Requirements: 3.1, 3.2, 3.3, 3.4, 3.5_
  - [x] 30.4 26-string-algorithm/problems/ 3개 문제 생성
    - _Requirements: 4.1~4.9, 5.1~5.5_

- [x] 31. Checkpoint - 19~26번 알고리즘 콘텐츠 검증
  - Ensure all .go files compile without errors, ask the user if questions arise.

- [x] 32. 검증 스크립트 작성
  - [x] 32.1 validate.sh 셸 스크립트 생성
    - 26개 알고리즘 폴더 존재 여부 확인
    - 각 폴더 내 필수 파일(README.md, theory.md, examples/*.go, problems/ 3개 문제 폴더) 존재 여부 확인
    - 폴더/파일 명명 규칙 준수 확인
    - 난이도 분포(easy, medium, hard) 균형 확인
    - 모든 .go 파일 컴파일 가능 여부 확인 (go build)
    - _Requirements: 7.1, 7.5_
  - [x] 32.2 Go 속성 기반 테스트 작성 (validate/ 패키지)
    - **Property 1: 폴더 구조 완전성** - 모든 알고리즘 폴더가 필수 파일/디렉토리를 포함하는지 검증
    - **Validates: Requirements 1.5, 2.1, 3.1, 4.1, 4.6, 5.1**
  - [x] 32.3 Go 속성 기반 테스트 작성 - 폴더 명명 규칙
    - **Property 2: 폴더 명명 규칙** - 모든 폴더명이 두 자리 번호 + 케밥 케이스 패턴을 만족하는지 검증
    - **Validates: Requirements 1.3**
  - [x] 32.4 Go 속성 기반 테스트 작성 - 이론 문서 구조
    - **Property 3: 이론 문서 구조 완전성** - 모든 theory.md가 필수 섹션을 포함하는지 검증
    - **Validates: Requirements 2.2, 2.3, 2.4, 2.5**
  - [x] 32.5 Go 속성 기반 테스트 작성 - 문제 파일 구조
    - **Property 4: 문제 파일 구조 완전성** - 모든 problem.md가 필수 섹션을 포함하는지 검증
    - **Validates: Requirements 4.2, 4.3, 4.4, 4.5, 4.8**
  - [x] 32.6 Go 속성 기반 테스트 작성 - 해설 문서 구조
    - **Property 5: 해설 문서 구조 완전성** - 모든 explanation.md가 필수 섹션을 포함하는지 검증
    - **Validates: Requirements 5.2, 5.3, 5.4**
  - [x] 32.7 Go 속성 기반 테스트 작성 - 난이도 분포
    - **Property 6: 난이도 분포 균형** - 모든 problems/ 디렉토리가 easy, medium, hard를 각각 포함하는지 검증
    - **Validates: Requirements 4.9**
  - [x] 32.8 Go 속성 기반 테스트 작성 - Go 코드 컴파일
    - **Property 7: Go 코드 컴파일 가능성** - 모든 .go 파일이 go build로 컴파일되는지 검증
    - **Validates: Requirements 3.3, 4.7, 7.1, 7.5**
  - [x] 32.9 Go 속성 기반 테스트 작성 - Go 코드 규칙
    - **Property 8: Go 코드 규칙 준수** - 모든 .go 파일이 package main, func main(), 한국어 주석, 표준 라이브러리만 사용 규칙을 만족하는지 검증
    - **Validates: Requirements 3.2, 3.4, 3.5, 7.2, 7.3, 7.4**

- [x] 33. Final checkpoint - 전체 프로젝트 검증 (01~26번)
  - validate.sh 실행하여 전체 구조 검증
  - Ensure all tests pass, ask the user if questions arise.

- [x] 34. 우선순위_1 확장 알고리즘 콘텐츠 생성 (27~29번)
  - [x] 34.1 27-geometry (기하학) 콘텐츠 생성
    - README.md: 알고리즘 유형 이름, 설명, 파일 목록, 문제 목록
    - theory.md: 개념 정의, 동작 원리, 복잡도, 적합한 문제 유형, 선수 학습 안내 (관련 기존 폴더 참조)
    - examples/: CCW, 외적 등 기본 기하 연산 예시 코드 (package main, 한국어 주석, 표준 라이브러리)
    - problems/: 01-easy, 02-medium, 03-hard 각각 problem.md, solution.go, explanation.md
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 34.2 28-combinatorics (조합론) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 07-math-and-number-theory 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 34.3 29-bitmask (비트마스킹) 콘텐츠 생성
    - README.md, theory.md, examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_

- [x] 35. Checkpoint - 우선순위_1 확장 알고리즘 (27~29번) 검증
  - Ensure all tests pass, ask the user if questions arise.

- [x] 36. 우선순위_2 확장 알고리즘 콘텐츠 생성 - 전반부 (30~34번)
  - [x] 36.1 30-game-theory (게임 이론) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 19-dynamic-programming 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 36.2 31-probability (확률론) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 07-math-and-number-theory, 19-dynamic-programming 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 36.3 32-bitmask-dp (비트필드 DP) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 19-dynamic-programming, 29-bitmask 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 36.4 33-maximum-flow (최대 유량) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 15-graph-dfs, 16-graph-bfs, 24-graph-advanced 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 36.5 34-primality-test (소수 판정) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 07-math-and-number-theory 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_

- [x] 37. 우선순위_2 확장 알고리즘 콘텐츠 생성 - 후반부 (35~39번)
  - [x] 37.1 35-offline-queries (오프라인 쿼리) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 03-sorting, 25-segment-tree 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 37.2 36-exponentiation-by-squaring (분할 정복 거듭제곱) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 18-divide-and-conquer, 07-math-and-number-theory 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 37.3 37-knapsack (배낭 문제) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 19-dynamic-programming 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 37.4 38-dag (DAG) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 15-graph-dfs, 23-topological-sort, 19-dynamic-programming 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 37.5 39-coordinate-compression (좌표 압축) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 03-sorting 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_

- [x] 38. Checkpoint - 우선순위_2 확장 알고리즘 (30~39번) 검증
  - Ensure all tests pass, ask the user if questions arise.

- [x] 39. 우선순위_3 확장 알고리즘 콘텐츠 생성 - 1차 (40~44번)
  - [x] 39.1 40-recursion (재귀) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 18-divide-and-conquer 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 39.2 41-euclidean-algorithm (유클리드 호제법) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 07-math-and-number-theory 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 39.3 42-convex-hull (볼록 껍질) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 27-geometry, 03-sorting 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 39.4 43-bipartite-matching (이분 매칭) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 15-graph-dfs, 33-maximum-flow 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 39.5 44-sieve-of-eratosthenes (에라토스테네스의 체) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 07-math-and-number-theory 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_

- [x] 40. 우선순위_3 확장 알고리즘 콘텐츠 생성 - 2차 (45~49번)
  - [x] 40.1 45-inclusion-exclusion (포함 배제의 원리) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 28-combinatorics, 07-math-and-number-theory 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 40.2 46-lca (LCA) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 13-tree, 15-graph-dfs 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 40.3 47-sparse-table (희소 배열) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 06-prefix-sum, 25-segment-tree 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 40.4 48-hashing (해싱) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 05-hash, 26-string-algorithm 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 40.5 49-modular-inverse (모듈로 곱셈 역원) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 07-math-and-number-theory, 36-exponentiation-by-squaring 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_

- [x] 41. Checkpoint - 우선순위_3 확장 알고리즘 1차/2차 (40~49번) 검증
  - Ensure all tests pass, ask the user if questions arise.

- [x] 42. 우선순위_3 확장 알고리즘 콘텐츠 생성 - 3차 (50~53번)
  - [x] 42.1 50-floyd-warshall (플로이드-워셜) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 21-shortest-path 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 42.2 51-trie (트라이) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 13-tree, 26-string-algorithm 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 42.3 52-deque (덱) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 04-stack-and-queue 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_
  - [x] 42.4 53-prime-factorization (소인수분해) 콘텐츠 생성
    - README.md, theory.md (선수 학습: 07-math-and-number-theory, 34-primality-test 참조), examples/, problems/ 3개 문제
    - _Requirements: 8.1, 8.2, 8.3, 8.5, 8.6, 1.5, 2.1~2.6, 3.1~3.5, 4.1~4.9, 5.1~5.5_

- [x] 43. Checkpoint - 우선순위_3 확장 알고리즘 3차 (50~53번) 검증
  - Ensure all tests pass, ask the user if questions arise.

- [x] 44. 루트 README.md 업데이트
  - 53개 알고리즘 유형 목록으로 확장 (27~53번 폴더 링크 추가)
  - 기존 알고리즘(01~26번)과 확장 알고리즘(27~53번) 구분 표시
  - 확장 알고리즘을 우선순위_1, 우선순위_2, 우선순위_3 그룹별로 분류하여 표시
  - 권장 학습 순서에 확장 알고리즘 반영
  - _Requirements: 9.1, 9.2, 6.3, 6.4_

- [x] 45. validate.sh 업데이트
  - FOLDERS 배열에 27~53번 폴더 27개 추가 (총 53개)
  - Go 컴파일 검증(go build) 로직 제거
  - 53개 폴더 전체에 대해 구조, 명명 규칙, 필수 파일, 난이도 분포 검증
  - _Requirements: 9.4, 7.1, 7.5_

- [x] 46. Go 속성 기반 테스트 업데이트 (validate/ 패키지)
  - [x] 46.1 테스트 대상 폴더 목록을 53개로 확장
    - 모든 속성 기반 테스트(Property 1~6)가 53개 폴더를 대상으로 검증하도록 수정
    - _Requirements: 8.2, 9.4_
  - [x] 46.2 Property 7 (Go 코드 컴파일 테스트) 제거
    - 기존 TestGoCodeCompilation 테스트 함수 삭제
    - Property 7은 Go 코드 규칙 준수(package main, func main(), 한국어 주석, 표준 라이브러리)만 검증
    - _Requirements: 3.2, 3.4, 3.5, 7.2, 7.3_
  - [x] 46.3 확장 알고리즘 선수 학습 참조 검증 테스트 추가
    - 확장 알고리즘(27~53번)의 theory.md에 선수 학습 폴더 참조가 포함되어 있는지 검증
    - _Requirements: 8.6, 9.3_

- [x] 47. Final checkpoint - 전체 프로젝트 검증 (53개 폴더)
  - validate.sh 실행하여 53개 폴더 전체 구조 검증
  - Go 속성 기반 테스트 실행하여 모든 속성 검증
  - Ensure all tests pass, ask the user if questions arise.

- [x] 48. 기존 01~53번 이론 문서 상세 보충 - 전반부 (01~13번)
  - [x] 48.1 01-implementation-and-simulation ~ 07-math-and-number-theory 이론 문서 보충
    - 각 폴더의 theory.md에 "단계별 추적(Trace)" 섹션 추가: 구체적인 입력 예시를 통한 알고리즘 실행 과정을 ASCII 다이어그램이나 텍스트로 시각화
    - "실전 팁" 섹션 추가: 코딩테스트 활용 노하우, 자주 하는 실수, 엣지 케이스 주의사항
    - "관련 알고리즘 비교" 섹션 추가: 유사한 알고리즘과의 차이점 및 선택 기준 표
    - 기존 내용(개념, 동작 원리, 복잡도, 적합한 문제 유형)은 유지하면서 보충
    - _Requirements: 12.1, 12.2, 12.3, 12.4, 12.5, 12.8_
  - [x] 48.2 08-binary-search ~ 13-tree 이론 문서 보충
    - 각 폴더의 theory.md에 "단계별 추적(Trace)", "실전 팁", "관련 알고리즘 비교" 섹션 추가
    - 기존 내용 유지하면서 보충
    - _Requirements: 12.1, 12.2, 12.3, 12.4, 12.5, 12.8_

- [x] 49. 기존 01~53번 이론 문서 상세 보충 - 중반부 (14~26번)
  - [x] 49.1 14-binary-tree ~ 19-dynamic-programming 이론 문서 보충
    - 각 폴더의 theory.md에 "단계별 추적(Trace)", "실전 팁", "관련 알고리즘 비교" 섹션 추가
    - 기존 내용 유지하면서 보충
    - _Requirements: 12.1, 12.2, 12.3, 12.4, 12.5, 12.8_
  - [x] 49.2 20-union-find ~ 26-string-algorithm 이론 문서 보충
    - 각 폴더의 theory.md에 "단계별 추적(Trace)", "실전 팁", "관련 알고리즘 비교" 섹션 추가
    - 기존 내용 유지하면서 보충
    - _Requirements: 12.1, 12.2, 12.3, 12.4, 12.5, 12.8_

- [x] 50. 기존 01~53번 이론 문서 상세 보충 - 후반부 (27~39번)
  - [x] 50.1 27-geometry ~ 33-maximum-flow 이론 문서 보충
    - 각 폴더의 theory.md에 "단계별 추적(Trace)", "실전 팁", "관련 알고리즘 비교" 섹션 추가
    - 기존 내용 유지하면서 보충
    - _Requirements: 12.1, 12.2, 12.3, 12.4, 12.5, 12.8_
  - [x] 50.2 34-primality-test ~ 39-coordinate-compression 이론 문서 보충
    - 각 폴더의 theory.md에 "단계별 추적(Trace)", "실전 팁", "관련 알고리즘 비교" 섹션 추가
    - 기존 내용 유지하면서 보충
    - _Requirements: 12.1, 12.2, 12.3, 12.4, 12.5, 12.8_

- [x] 51. 기존 01~53번 이론 문서 상세 보충 - 마지막 (40~53번)
  - [x] 51.1 40-recursion ~ 46-lca 이론 문서 보충
    - 각 폴더의 theory.md에 "단계별 추적(Trace)", "실전 팁", "관련 알고리즘 비교" 섹션 추가
    - 기존 내용 유지하면서 보충
    - _Requirements: 12.1, 12.2, 12.3, 12.4, 12.5, 12.8_
  - [x] 51.2 47-sparse-table ~ 53-prime-factorization 이론 문서 보충
    - 각 폴더의 theory.md에 "단계별 추적(Trace)", "실전 팁", "관련 알고리즘 비교" 섹션 추가
    - 기존 내용 유지하면서 보충
    - _Requirements: 12.1, 12.2, 12.3, 12.4, 12.5, 12.8_

- [x] 52. Checkpoint - 이론 문서 상세 보충 검증 (01~53번)
  - 모든 01~53번 폴더의 theory.md에 "단계별 추적" 또는 "Trace", "실전 팁", "관련 알고리즘 비교" 섹션이 존재하는지 확인
  - 기존 내용(개념, 동작 원리, 복잡도, 적합한 문제 유형)이 유지되었는지 확인
  - Ensure all tests pass, ask the user if questions arise.

- [x] 53. Tier_1 알고리즘 콘텐츠 생성 (54~59번)
  - [x] 53.1 54-tree-dp (트리 DP) 콘텐츠 생성
    - README.md: 알고리즘 유형 이름, 설명, 파일 목록, 문제 목록
    - theory.md: 개념 정의, 동작 원리, 복잡도, 적합한 문제 유형, 단계별 추적(Trace), 실전 팁, 관련 알고리즘 비교, 선수 학습 안내 (13-tree, 19-dynamic-programming)
    - examples/: 서브트리 합, 리루팅 등 기본 구현 예시 코드 (package main, 한국어 주석, 표준 라이브러리)
    - problems/: 01-easy, 02-medium, 03-hard 각각 problem.md, solution.go, explanation.md
    - _Requirements: 10.1, 10.2, 10.3, 10.4, 10.5, 10.6, 12.9, 2.1~2.12, 3.1~3.6, 4.1~4.9, 5.1~5.6_
  - [x] 53.2 55-lis (LIS 최장 증가 부분 수열) 콘텐츠 생성
    - README.md, theory.md (상세 포함, 선수 학습: 08-binary-search, 19-dynamic-programming), examples/, problems/ 3개
    - _Requirements: 10.1, 10.2, 10.3, 10.4, 10.5, 10.6, 12.9, 2.1~2.12, 3.1~3.6, 4.1~4.9, 5.1~5.6_
  - [x] 53.3 56-sqrt-decomposition (제곱근 분할법) 콘텐츠 생성
    - README.md, theory.md (상세 포함, 선수 학습: 35-offline-queries, 25-segment-tree), examples/, problems/ 3개
    - _Requirements: 10.1, 10.2, 10.3, 10.4, 10.5, 10.6, 12.9, 2.1~2.12, 3.1~3.6, 4.1~4.9, 5.1~5.6_
  - [x] 53.4 57-meet-in-the-middle (중간에서 만나기) 콘텐츠 생성
    - README.md, theory.md (상세 포함, 선수 학습: 02-bruteforce, 03-sorting), examples/, problems/ 3개
    - _Requirements: 10.1, 10.2, 10.3, 10.4, 10.5, 10.6, 12.9, 2.1~2.12, 3.1~3.6, 4.1~4.9, 5.1~5.6_
  - [x] 53.5 58-zero-one-bfs (0-1 BFS) 콘텐츠 생성
    - README.md, theory.md (상세 포함, 선수 학습: 16-graph-bfs, 52-deque, 21-shortest-path), examples/, problems/ 3개
    - _Requirements: 10.1, 10.2, 10.3, 10.4, 10.5, 10.6, 12.9, 2.1~2.12, 3.1~3.6, 4.1~4.9, 5.1~5.6_
  - [x] 53.6 59-flood-fill (플러드 필) 콘텐츠 생성
    - README.md, theory.md (상세 포함, 선수 학습: 15-graph-dfs, 16-graph-bfs), examples/, problems/ 3개
    - _Requirements: 10.1, 10.2, 10.3, 10.4, 10.5, 10.6, 12.9, 2.1~2.12, 3.1~3.6, 4.1~4.9, 5.1~5.6_

- [x] 54. Checkpoint - Tier_1 알고리즘 (54~59번) 검증
  - 54~59번 폴더의 구조 완전성 확인 (README.md, theory.md, examples/, problems/)
  - theory.md에 단계별 추적, 실전 팁, 관련 알고리즘 비교, 선수 학습 참조가 포함되어 있는지 확인
  - 각 problems/ 디렉토리에 easy, medium, hard 난이도가 모두 존재하는지 확인
  - Ensure all tests pass, ask the user if questions arise.

- [x] 55. Tier_2 알고리즘 콘텐츠 생성 - 전반부 (60~63번)
  - [x] 55.1 60-fft (FFT 고속 푸리에 변환) 콘텐츠 생성
    - README.md, theory.md (상세 포함, 선수 학습: 07-math-and-number-theory, 18-divide-and-conquer), examples/, problems/ 3개
    - _Requirements: 11.1, 11.2, 11.3, 11.4, 11.5, 11.6, 12.9, 2.1~2.12, 3.1~3.6, 4.1~4.9, 5.1~5.6_
  - [x] 55.2 61-ternary-search (삼분 탐색) 콘텐츠 생성
    - README.md, theory.md (상세 포함, 선수 학습: 08-binary-search), examples/, problems/ 3개
    - _Requirements: 11.1, 11.2, 11.3, 11.4, 11.5, 11.6, 12.9, 2.1~2.12, 3.1~3.6, 4.1~4.9, 5.1~5.6_
  - [x] 55.3 62-euler-tour (오일러 경로 테크닉) 콘텐츠 생성
    - README.md, theory.md (상세 포함, 선수 학습: 13-tree, 25-segment-tree, 15-graph-dfs), examples/, problems/ 3개
    - _Requirements: 11.1, 11.2, 11.3, 11.4, 11.5, 11.6, 12.9, 2.1~2.12, 3.1~3.6, 4.1~4.9, 5.1~5.6_
  - [x] 55.4 63-mcmf (최소 비용 최대 유량) 콘텐츠 생성
    - README.md, theory.md (상세 포함, 선수 학습: 33-maximum-flow, 21-shortest-path), examples/, problems/ 3개
    - _Requirements: 11.1, 11.2, 11.3, 11.4, 11.5, 11.6, 12.9, 2.1~2.12, 3.1~3.6, 4.1~4.9, 5.1~5.6_

- [x] 56. Tier_2 알고리즘 콘텐츠 생성 - 후반부 (64~67번)
  - [x] 56.1 64-convex-hull-trick (볼록 껍질 트릭) 콘텐츠 생성
    - README.md, theory.md (상세 포함, 선수 학습: 19-dynamic-programming, 42-convex-hull), examples/, problems/ 3개
    - _Requirements: 11.1, 11.2, 11.3, 11.4, 11.5, 11.6, 12.9, 2.1~2.12, 3.1~3.6, 4.1~4.9, 5.1~5.6_
  - [x] 56.2 65-gaussian-elimination (가우스 소거법) 콘텐츠 생성
    - README.md, theory.md (상세 포함, 선수 학습: 07-math-and-number-theory), examples/, problems/ 3개
    - _Requirements: 11.1, 11.2, 11.3, 11.4, 11.5, 11.6, 12.9, 2.1~2.12, 3.1~3.6, 4.1~4.9, 5.1~5.6_
  - [x] 56.3 66-hld (HLD) 콘텐츠 생성
    - README.md, theory.md (상세 포함, 선수 학습: 13-tree, 25-segment-tree, 15-graph-dfs), examples/, problems/ 3개
    - _Requirements: 11.1, 11.2, 11.3, 11.4, 11.5, 11.6, 12.9, 2.1~2.12, 3.1~3.6, 4.1~4.9, 5.1~5.6_
  - [x] 56.4 67-centroid-decomposition (센트로이드 분할) 콘텐츠 생성
    - README.md, theory.md (상세 포함, 선수 학습: 13-tree, 18-divide-and-conquer), examples/, problems/ 3개
    - _Requirements: 11.1, 11.2, 11.3, 11.4, 11.5, 11.6, 12.9, 2.1~2.12, 3.1~3.6, 4.1~4.9, 5.1~5.6_

- [x] 57. Checkpoint - Tier_2 알고리즘 (60~67번) 검증
  - 60~67번 폴더의 구조 완전성 확인 (README.md, theory.md, examples/, problems/)
  - theory.md에 단계별 추적, 실전 팁, 관련 알고리즘 비교, 선수 학습 참조가 포함되어 있는지 확인
  - 각 problems/ 디렉토리에 easy, medium, hard 난이도가 모두 존재하는지 확인
  - Ensure all tests pass, ask the user if questions arise.

- [x] 58. 루트 README.md 업데이트
  - [x] 58.1 알고리즘 유형 목록 확장
    - 67개 알고리즘 유형 목록으로 확장 (54~67번 폴더 링크 추가)
    - Tier_1(54~59번), Tier_2(60~67번) 그룹 구분 표시
    - _Requirements: 9.1, 9.2, 6.3_
  - [x] 58.2 권장 학습 순서 업데이트
    - 권장 학습 순서에 Tier_1, Tier_2 반영
    - Tier_1은 기존 53개 이후, Tier_2는 Tier_1 이후 학습 권장
    - _Requirements: 6.4, 9.2_

- [x] 59. validate.sh 및 Go 테스트 업데이트
  - [x] 59.1 validate.sh 업데이트
    - FOLDERS 배열에 54~67번 폴더 14개 추가 (총 67개)
    - 67개 폴더 전체에 대해 구조, 명명 규칙, 필수 파일, 난이도 분포 검증
    - _Requirements: 9.4_
  - [x] 59.2 Go 속성 기반 테스트 업데이트 (validate/ 패키지)
    - 테스트 대상 폴더 목록을 67개로 확장
    - Property 8 (이론 문서 상세 보충 완전성) 테스트 추가: 모든 theory.md에 "단계별 추적" 또는 "Trace", "실전 팁", "관련 알고리즘 비교" 섹션 존재 여부 검증
    - **Property 8: 이론 문서 상세 보충 완전성**
    - **Validates: Requirements 12.2, 12.4, 12.5, 12.9**
    - _Requirements: 9.4, 12.2, 12.4, 12.5_
  - [x] 59.3 Tier_1/Tier_2 선수 학습 참조 검증 테스트 추가
    - 54~59번(Tier_1) 알고리즘의 theory.md에 지정된 선수 학습 폴더가 참조되어 있는지 검증
    - 60~67번(Tier_2) 알고리즘의 theory.md에 지정된 선수 학습 폴더가 참조되어 있는지 검증
    - _Requirements: 10.5, 10.6, 11.5, 11.6_

- [x] 60. Final checkpoint - 전체 프로젝트 검증 (67개 폴더)
  - validate.sh 실행하여 67개 폴더 전체 구조 검증
  - Go 속성 기반 테스트 실행하여 모든 속성 검증 (Property 1~8)
  - Ensure all tests pass, ask the user if questions arise.

## Notes

- `*` 표시된 태스크는 선택 사항이며 빠른 MVP를 위해 건너뛸 수 있음
- 각 태스크는 특정 요구사항을 참조하여 추적 가능성을 보장
- 체크포인트는 점진적 검증을 보장
- 속성 기반 테스트는 설계 문서의 정확성 속성을 검증
- 모든 Go 코드는 `go run`으로 즉시 실행 가능해야 함
- 모든 문서는 한국어로 작성
- 확장 알고리즘(27~53번)은 우선순위_1 → 우선순위_2 → 우선순위_3 순서로 생성
- 확장 알고리즘의 theory.md에는 관련 기존 폴더를 선수 학습으로 참조
- Tier_1(54~59번)은 코테 실전 필수 유형, Tier_2(60~67번)은 코테 상위권 차별화 유형
- 48~51번 태스크는 기존 01~53번 이론 문서에 상세 보충 섹션을 추가하는 작업
- 54~67번 알고리즘의 theory.md는 처음부터 상세 섹션(Trace, 실전 팁, 관련 알고리즘 비교)을 포함하여 작성

