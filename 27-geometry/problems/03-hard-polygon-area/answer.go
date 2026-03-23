package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// polygonArea는 다각형의 넓이를 구한다.
//
// [매개변수]
//   - x: 꼭짓점의 x 좌표 배열
//   - y: 꼭짓점의 y 좌표 배열
//   - n: 꼭짓점의 수
//
// [반환값]
//   - float64: 다각형의 넓이
//
// [알고리즘 힌트]
//
//	신발끈 공식(Shoelace Formula)을 사용한다.
//	넓이 = |Σ(x_i * y_{i+1} - x_{i+1} * y_i)| / 2
//	마지막 꼭짓점은 첫 번째 꼭짓점과 연결하여 순환한다.
//	외적을 누적한 후 절댓값을 취하고 2로 나눈다.
func polygonArea(x, y []int64, n int) float64 {
	var sum int64
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		sum += x[i]*y[j] - x[j]*y[i]
	}
	return math.Abs(float64(sum)) / 2.0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	x := make([]int64, n)
	y := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i], &y[i])
	}

	// 핵심 함수 호출
	area := polygonArea(x, y, n)

	// 소수점 첫째 자리까지 출력
	fmt.Fprintf(writer, "%.1f\n", area)
}
