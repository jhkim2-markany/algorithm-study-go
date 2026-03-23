package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// minMaxDistance는 x축 위의 점 P(t, 0)에서 주어진 점들까지의 최대 거리를
// 최소화하는 t와 그때의 최대 거리를 삼분 탐색으로 구한다.
//
// [매개변수]
//   - x, y: 각 점의 x좌표, y좌표 배열
//   - n: 점의 개수
//
// [반환값]
//   - float64: 최대 거리를 최소화하는 t 좌표
//   - float64: 해당 t에서의 최대 거리
func minMaxDistance(x, y []float64, n int) (float64, float64) {
	// 여기에 코드를 작성하세요
	_ = math.Sqrt
	return 0, 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i], &y[i])
	}

	t, dist := minMaxDistance(x, y, n)
	fmt.Fprintf(writer, "%.6f\n", t)
	fmt.Fprintf(writer, "%.6f\n", dist)
}
