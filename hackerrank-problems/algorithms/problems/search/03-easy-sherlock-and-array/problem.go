package main

import (
	"bufio"
	"fmt"
	"os"
)

// balancedSums는 왼쪽 합과 오른쪽 합이 같은 원소가 존재하는지 판별한다.
//
// [매개변수]
//   - arr: 정수 배열
//
// [반환값]
//   - string: "YES" 또는 "NO"
func balancedSums(arr []int) string {
	// 여기에 코드를 작성하세요
	return ""
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		var n int
		fmt.Fscan(reader, &n)

		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &arr[i])
		}

		fmt.Fprintln(writer, balancedSums(arr))
	}
}
