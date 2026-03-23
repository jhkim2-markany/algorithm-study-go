package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// buildSuffixArray는 문자열의 접미사 배열을 구축한다.
//
// [매개변수]
//   - s: 입력 문자열
//
// [반환값]
//   - []int: 접미사 배열 (각 접미사의 시작 인덱스)
//
// [알고리즘 힌트]
//
//	O(N log²N) 방식으로 구현한다.
//	초기 순위를 각 문자의 ASCII 값으로 설정하고,
//	길이를 2배씩 늘려가며 (현재 순위, k칸 뒤 순위) 두 키로 정렬한다.
//	정렬 후 새로운 순위를 부여하고, 모든 순위가 고유하면 조기 종료한다.
func buildSuffixArray(s string) []int {
	n := len(s)
	sa := make([]int, n)
	rank := make([]int, n)
	tmp := make([]int, n)

	for i := 0; i < n; i++ {
		sa[i] = i
		rank[i] = int(s[i])
	}

	for k := 1; k < n; k *= 2 {
		kk := k
		sort.Slice(sa, func(i, j int) bool {
			if rank[sa[i]] != rank[sa[j]] {
				return rank[sa[i]] < rank[sa[j]]
			}
			ri, rj := -1, -1
			if sa[i]+kk < n {
				ri = rank[sa[i]+kk]
			}
			if sa[j]+kk < n {
				rj = rank[sa[j]+kk]
			}
			return ri < rj
		})

		tmp[sa[0]] = 0
		for i := 1; i < n; i++ {
			tmp[sa[i]] = tmp[sa[i-1]]
			ri1, ri0 := -1, -1
			rj1, rj0 := -1, -1
			ri0 = rank[sa[i-1]]
			rj0 = rank[sa[i]]
			if sa[i-1]+kk < n {
				ri1 = rank[sa[i-1]+kk]
			}
			if sa[i]+kk < n {
				rj1 = rank[sa[i]+kk]
			}
			if ri0 != rj0 || ri1 != rj1 {
				tmp[sa[i]]++
			}
		}
		copy(rank, tmp)

		if rank[sa[n-1]] == n-1 {
			break
		}
	}
	return sa
}

// buildLCP는 접미사 배열로부터 LCP 배열을 구축한다.
//
// [매개변수]
//   - s: 입력 문자열
//   - sa: 접미사 배열
//
// [반환값]
//   - []int: LCP 배열 (인접한 접미사 간의 최장 공통 접두사 길이)
//
// [알고리즘 힌트]
//
//	Kasai 알고리즘을 사용한다. O(N) 시간 복잡도.
//	접미사 배열의 역배열(순위 배열)을 구한 후,
//	이전 LCP 값을 활용하여 효율적으로 계산한다.
//	h 값은 최대 1만 감소하므로 전체 O(N)이 보장된다.
func buildLCP(s string, sa []int) []int {
	n := len(s)
	lcp := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < n; i++ {
		rank[sa[i]] = i
	}

	h := 0
	for i := 0; i < n; i++ {
		if rank[i] > 0 {
			j := sa[rank[i]-1]
			for i+h < n && j+h < n && s[i+h] == s[j+h] {
				h++
			}
			lcp[rank[i]] = h
			if h > 0 {
				h--
			}
		} else {
			h = 0
		}
	}
	return lcp
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var s string
	fmt.Fscan(reader, &s)

	// 핵심 함수 호출
	sa := buildSuffixArray(s)
	lcp := buildLCP(s, sa)

	// 접미사 배열 출력
	for i, v := range sa {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)

	// LCP 배열 출력
	for i, v := range lcp {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
