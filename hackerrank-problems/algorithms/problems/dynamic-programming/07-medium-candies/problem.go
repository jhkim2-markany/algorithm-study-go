package main

import (
	"bufio"
	"fmt"
	"os"
)

// candies는 조건을 만족하는 최소 사탕 수를 반환한다.
//
// [매개변수]
//   - n: 학생 수
//   - arr: 각 학생의 성적 배열
//
// [반환값]
//   - int64: 필요한 사탕의 최소 총 개수
func candies(n int, arr []int) int64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 학생 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 성적 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출 및 결과 출력
	result := candies(n, arr)
	fmt.Fprintln(writer, result)
}
