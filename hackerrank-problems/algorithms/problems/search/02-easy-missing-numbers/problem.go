package main

import (
	"bufio"
	"fmt"
	"os"
)

// missingNumbers는 원본 배열에는 있지만 첫 번째 배열에는 빠진 숫자들을 반환한다.
//
// [매개변수]
//   - arr: 일부 숫자가 빠진 배열
//   - brr: 원본 배열
//
// [반환값]
//   - []int: 빠진 숫자들 (오름차순)
func missingNumbers(arr []int, brr []int) []int {
	// 여기에 코드를 작성하세요
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	var m int
	fmt.Fscan(reader, &m)
	brr := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &brr[i])
	}

	result := missingNumbers(arr, brr)
	for i, v := range result {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
