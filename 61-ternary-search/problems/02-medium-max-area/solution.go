package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// maxArea는 길이 l의 철사를 정사각형과 원으로 나눌 때,
// 삼분 탐색으로 넓이 합의 최솟값 위치를 찾고 양 끝점과 비교하여
// 최대 넓이를 달성하는 분할 지점과 그 넓이를 반환한다.
//
// [매개변수]
//   - l: 철사의 전체 길이
//
// [반환값]
//   - float64: 최대 넓이를 달성하는 정사각형 사용 길이 x
//   - float64: 해당 분할에서의 최대 넓이
func maxArea(l float64) (float64, float64) {
	// 여기에 코드를 작성하세요
	_ = math.Pi
	return 0, 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var l float64
	fmt.Fscan(reader, &l)

	bestX, bestArea := maxArea(l)
	fmt.Fprintf(writer, "%.6f\n", bestX)
	fmt.Fprintf(writer, "%.6f\n", bestArea)
}
