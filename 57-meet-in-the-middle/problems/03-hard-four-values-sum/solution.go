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
func fourValuesSum(a, b, c, d []int) int {
	// 여기에 코드를 작성하세요
	_ = sort.SearchInts
	return 0
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
