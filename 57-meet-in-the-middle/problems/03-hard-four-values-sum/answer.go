package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// fourValuesSum은 네 배열에서 각각 하나씩 골라 합이 0이 되는 조합의 수를 반환한다.
//
// [매개변수]
//   - a, b, c, d: 네 개의 정수 배열 (크기 동일)
//
// [반환값]
//   - int: A[i]+B[j]+C[k]+D[l]=0을 만족하는 (i,j,k,l) 쌍의 수
//
// [알고리즘 힌트]
//
//	Meet in the Middle 기법을 사용한다.
//	A+B의 모든 합(N²개)과 C+D의 모든 합(N²개)을 각각 구한다.
//	C+D 배열을 정렬한 뒤, A+B의 각 원소 x에 대해
//	C+D에서 -x의 개수를 이분 탐색(lower_bound, upper_bound)으로 센다.
//	전체 시간 복잡도: O(N² log N)
func fourValuesSum(a, b, c, d []int) int {
	n := len(a)

	// A+B의 모든 합을 구한다
	ab := make([]int, 0, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ab = append(ab, a[i]+b[j])
		}
	}

	// C+D의 모든 합을 구한다
	cd := make([]int, 0, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			cd = append(cd, c[i]+d[j])
		}
	}

	// cd를 정렬하여 이분 탐색 준비
	sort.Ints(cd)

	// ab의 각 원소에 대해 cd에서 보완값의 개수를 센다
	count := 0
	for _, x := range ab {
		target := -x
		lo := sort.SearchInts(cd, target)
		hi := sort.SearchInts(cd, target+1)
		count += hi - lo
	}

	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	d := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &d[i])
	}

	// 핵심 함수 호출
	fmt.Fprintln(writer, fourValuesSum(a, b, c, d))
}
