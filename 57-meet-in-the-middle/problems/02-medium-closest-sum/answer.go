package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// closestSubsetSum은 Meet in the Middle 기법으로 목표값에 가장 가까운 부분집합 합을 반환한다.
//
// [매개변수]
//   - arr: 정수 배열
//   - s: 목표값
//
// [반환값]
//   - int: 목표값에 가장 가까운 부분집합 합
//
// [알고리즘 힌트]
//   - 배열을 반으로 분할하여 각 절반의 모든 부분집합 합을 열거한다
//   - 한쪽(sumB)을 정렬한 뒤, 다른 쪽(sumA)의 각 원소에 대해 need = s - a를 이분 탐색한다
//   - idx 위치(need 이상인 가장 작은 값)와 idx-1 위치(need 미만인 가장 큰 값)를 모두 확인한다
//   - |total - s|가 최소인 합을 추적한다
func closestSubsetSum(arr []int, s int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

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

	bestDiff := int(1e18)
	bestSum := 0

	for _, a := range sumA {
		need := s - a
		idx := sort.SearchInts(sumB, need)

		if idx < len(sumB) {
			total := a + sumB[idx]
			if abs(total-s) < bestDiff {
				bestDiff = abs(total - s)
				bestSum = total
			}
		}

		if idx > 0 {
			total := a + sumB[idx-1]
			if abs(total-s) < bestDiff {
				bestDiff = abs(total - s)
				bestSum = total
			}
		}
	}

	return bestSum
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

	fmt.Fprintln(writer, closestSubsetSum(arr, s))
}
