#!/bin/bash

# 프로젝트 구조 검증 스크립트
# 67개 알고리즘 폴더의 구조, 명명 규칙, 필수 파일 존재 여부, 난이도 분포를 검증한다.
# 실패 항목만 출력하여 빠르게 결과를 확인할 수 있다.

PASS=0
FAIL=0
ERRORS=""

pass() { PASS=$((PASS + 1)); }
fail() {
  FAIL=$((FAIL + 1))
  ERRORS="${ERRORS}  [FAIL] $1\n"
}

FOLDERS=(
  "01-implementation-and-simulation"
  "02-bruteforce"
  "03-sorting"
  "04-stack-and-queue"
  "05-hash"
  "06-prefix-sum"
  "07-math-and-number-theory"
  "08-binary-search"
  "09-parametric-search"
  "10-two-pointer-and-sliding-window"
  "11-greedy"
  "12-heap-and-priority-queue"
  "13-tree"
  "14-binary-tree"
  "15-graph-dfs"
  "16-graph-bfs"
  "17-backtracking"
  "18-divide-and-conquer"
  "19-dynamic-programming"
  "20-union-find"
  "21-shortest-path"
  "22-minimum-spanning-tree"
  "23-topological-sort"
  "24-graph-advanced"
  "25-segment-tree"
  "26-string-algorithm"
  "27-geometry"
  "28-combinatorics"
  "29-bitmask"
  "30-game-theory"
  "31-probability"
  "32-bitmask-dp"
  "33-maximum-flow"
  "34-primality-test"
  "35-offline-queries"
  "36-exponentiation-by-squaring"
  "37-knapsack"
  "38-dag"
  "39-coordinate-compression"
  "40-recursion"
  "41-euclidean-algorithm"
  "42-convex-hull"
  "43-bipartite-matching"
  "44-sieve-of-eratosthenes"
  "45-inclusion-exclusion"
  "46-lca"
  "47-sparse-table"
  "48-hashing"
  "49-modular-inverse"
  "50-floyd-warshall"
  "51-trie"
  "52-deque"
  "53-prime-factorization"
  "54-tree-dp"
  "55-lis"
  "56-sqrt-decomposition"
  "57-meet-in-the-middle"
  "58-zero-one-bfs"
  "59-flood-fill"
  "60-fft"
  "61-ternary-search"
  "62-euler-tour"
  "63-mcmf"
  "64-convex-hull-trick"
  "65-gaussian-elimination"
  "66-hld"
  "67-centroid-decomposition"
)

echo "=== 프로젝트 구조 검증 ==="

# 1~4. 폴더 존재, 명명 규칙, 필수 파일, 난이도 분포
for folder in "${FOLDERS[@]}"; do
  # 폴더 존재
  if [ -d "$folder" ]; then pass; else fail "$folder 폴더 누락"; continue; fi

  # 명명 규칙
  if echo "$folder" | grep -qE '^[0-9]{2}-[a-z]+(-[a-z]+)*$'; then pass; else fail "$folder 명명 규칙 위반"; fi

  # 필수 파일
  for f in README.md theory.md; do
    if [ -f "$folder/$f" ]; then pass; else fail "$folder/$f 누락"; fi
  done

  # examples/
  if [ -d "$folder/examples" ]; then
    go_count=$(find "$folder/examples" -name "*.go" | wc -l)
    if [ "$go_count" -ge 1 ]; then pass; else fail "$folder/examples/ 에 .go 파일 없음"; fi
  else
    fail "$folder/examples/ 누락"
  fi

  # problems/ 및 난이도 분포
  if [ -d "$folder/problems" ]; then
    has_easy=false; has_medium=false; has_hard=false
    for pdir in "$folder/problems"/*/; do
      pname=$(basename "$pdir")
      case "$pname" in *easy*) has_easy=true;; *medium*) has_medium=true;; *hard*) has_hard=true;; esac
      for req_file in problem.md solution.go answer.go explanation.md; do
        if [ -f "$pdir/$req_file" ]; then pass; else fail "$folder/problems/$pname/$req_file 누락"; fi
      done
    done
    if $has_easy && $has_medium && $has_hard; then pass; else fail "$folder 난이도 분포 불완전"; fi
  else
    fail "$folder/problems/ 누락"
  fi
done



# 결과 요약
TOTAL=$((PASS + FAIL))
echo ""
if [ -n "$ERRORS" ]; then
  echo "--- 실패 항목 ---"
  echo -e "$ERRORS"
fi
echo "=== 결과: 총 $TOTAL건 검사, 통과 $PASS, 실패 $FAIL ==="
if [ "$FAIL" -eq 0 ]; then
  echo "모든 검증 통과!"
else
  exit 1
fi
