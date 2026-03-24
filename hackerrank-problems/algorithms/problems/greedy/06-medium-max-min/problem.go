package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// maxMin은 배열에서 K개의 원소를 선택하여 최소 불공정도를 반환한다.
//
// [매개변수]
//   - k: 선택할 원소의 수
//   - arr: 정수 배열
//
// [반환값]
//   - int: 최소 불공정도
func maxMin(k int, arr []int) int {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// N, K 입력
	var n, k int
	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &k)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출 및 결과 출력
	result := maxMin(k, arr)
	fmt.Fprintln(writer, result)

	// sort 패키지 사용을 위한 임포트 유지
	_ = sort.Ints
}
