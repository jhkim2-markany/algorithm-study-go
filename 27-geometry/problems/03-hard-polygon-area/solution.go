package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 꼭짓점 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 꼭짓점 좌표 입력
	x := make([]int64, n)
	y := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i], &y[i])
	}

	// 신발끈 공식 (Shoelace Formula)으로 다각형 넓이 계산
	// 넓이 = |Σ(x_i * y_{i+1} - x_{i+1} * y_i)| / 2
	var sum int64
	for i := 0; i < n; i++ {
		// 다음 꼭짓점 인덱스 (마지막 점은 첫 번째 점과 연결)
		j := (i + 1) % n
		// 외적 누적: x_i * y_j - x_j * y_i
		sum += x[i]*y[j] - x[j]*y[i]
	}

	// 절댓값을 취하고 2로 나누어 넓이를 구한다
	area := math.Abs(float64(sum)) / 2.0

	// 소수점 첫째 자리까지 출력
	fmt.Fprintf(writer, "%.1f\n", area)
}
