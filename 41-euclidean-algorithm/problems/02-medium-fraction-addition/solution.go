package main

import (
	"bufio"
	"fmt"
	"os"
)

// addFractions는 여러 분수의 합을 기약분수로 반환한다.
//
// [매개변수]
//   - nums: 각 분수의 분자 배열
//   - dens: 각 분수의 분모 배열
//
// [반환값]
//   - int: 결과 기약분수의 분자
//   - int: 결과 기약분수의 분모
func addFractions(nums, dens []int) (int, int) {
	// 여기에 코드를 작성하세요
	return 0, 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	nums := make([]int, n)
	dens := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i], &dens[i])
	}

	num, den := addFractions(nums, dens)
	fmt.Fprintln(writer, num, den)
}
