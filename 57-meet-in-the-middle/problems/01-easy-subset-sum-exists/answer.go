package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// subsetSumExists는 Meet in the Middle 기법으로 부분집합 합이 목표값과 같은지 판별한다.
//
// [매개변수]
//   - arr: 정수 배열
//   - s: 목표 합
//
// [반환값]
//   - bool: 부분집합 합이 s와 같은 경우 true
//
// [알고리즘 힌트]
//   - 배열을 반으로 분할하여 각 절반의 모든 부분집합 합을 열거한다 (2^(N/2)개씩)
//   - 한쪽(sumB)을 정렬한 뒤, 다른 쪽(sumA)의 각 원소에 대해 보완값(s-a)을 이분 탐색한다
//   - 시간 복잡도: O(2^(N/2) × log(2^(N/2))) = O(N × 2^(N/2))
func subsetSumExists(arr []int, s int) bool {
	enumSums := func(a []int) []int {
		n := len(a)
		sums := make([]int, 0, 1<<n)
		for mask := 0; mask < (1 << n); mask++ {
			sum := 0
			for i := 0; i < n; i++ {
				if mask&(1<<i) != 0 {
					sum += a[i]
				}
			}
			sums = append(sums, sum)
		}
		return sums
	}

	half := len(arr) / 2
	sumA := enumSums(arr[:half])
	sumB := enumSums(arr[half:])

	sort.Ints(sumB)

	for _, a := range sumA {
		need := s - a
		idx := sort.SearchInts(sumB, need)
		if idx < len(sumB) && sumB[idx] == need {
			return true
		}
	}

	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, s int
	fmt.Fscan(reader, &n, &s)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	if subsetSumExists(arr, s) {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
