package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// buildSuffixArray 함수는 문자열의 접미사 배열을 구축한다
// O(N log²N) 방식으로 구현한다
func buildSuffixArray(s string) []int {
	n := len(s)
	sa := make([]int, n)
	rank := make([]int, n)
	tmp := make([]int, n)

	// 초기 순위: 각 문자의 ASCII 값
	for i := 0; i < n; i++ {
		sa[i] = i
		rank[i] = int(s[i])
	}

	// 길이를 2배씩 늘려가며 정렬한다
	for k := 1; k < n; k *= 2 {
		kk := k
		// 두 개의 키(현재 순위, k칸 뒤 순위)를 기준으로 정렬한다
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

		// 새로운 순위를 부여한다
		tmp[sa[0]] = 0
		for i := 1; i < n; i++ {
			tmp[sa[i]] = tmp[sa[i-1]]
			// 이전 원소와 다르면 순위를 증가시킨다
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

		// 모든 순위가 고유하면 조기 종료한다
		if rank[sa[n-1]] == n-1 {
			break
		}
	}
	return sa
}

// buildLCP 함수는 접미사 배열로부터 LCP 배열을 구축한다 (Kasai 알고리즘)
// 시간 복잡도: O(N)
func buildLCP(s string, sa []int) []int {
	n := len(s)
	lcp := make([]int, n)
	rank := make([]int, n)

	// 접미사 배열의 역배열(순위 배열)을 구한다
	for i := 0; i < n; i++ {
		rank[sa[i]] = i
	}

	// Kasai 알고리즘: 이전 LCP 값을 활용하여 효율적으로 계산한다
	h := 0
	for i := 0; i < n; i++ {
		if rank[i] > 0 {
			// 접미사 배열에서 바로 앞에 있는 접미사와 비교한다
			j := sa[rank[i]-1]
			// 공통 접두사 길이를 계산한다
			for i+h < n && j+h < n && s[i+h] == s[j+h] {
				h++
			}
			lcp[rank[i]] = h
			// 다음 반복을 위해 h를 1 감소시킨다
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

	// 문자열 입력
	var s string
	fmt.Fscan(reader, &s)

	// 접미사 배열 구성
	sa := buildSuffixArray(s)

	// LCP 배열 계산
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
