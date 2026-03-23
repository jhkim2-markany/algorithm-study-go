package main

import (
	"bufio"
	"fmt"
	"os"
)

// collectAllExpected는 n종류의 아이템을 모두 모으기 위한 기대 횟수를 반환한다.
//
// [매개변수]
//   - n: 아이템 종류의 수 (1 이상, 최대 20)
//   - prob: 각 아이템이 뽑힐 확률 배열 (길이 n, 합이 1.0)
//
// [반환값]
//   - float64: 모든 아이템을 모으기까지의 기대 횟수
func collectAllExpected(n int, prob []float64) float64 {
	// 여기에 코드를 작성하세요
	return 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	prob := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &prob[i])
	}

	fmt.Fprintf(writer, "%.6f\n", collectAllExpected(n, prob))
}
