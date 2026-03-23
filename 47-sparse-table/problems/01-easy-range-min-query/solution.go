package main

import (
	"bufio"
	"fmt"
	"os"
)

// Sparse Table을 이용한 구간 최솟값 쿼리 (RMQ)
// 전처리: O(N log N), 쿼리: O(1)

var sparse [17][100001]int // sparse[k][i] = 인덱스 i에서 길이 2^k 구간의 최솟값
var logTable [100002]int

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력: 배열 크기와 쿼리 수
	var n, m int
	fmt.Fscan(reader, &n, &m)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 로그 값 전처리
	for i := 2; i <= n; i++ {
		logTable[i] = logTable[i/2] + 1
	}

	// Sparse Table 구성: 길이 1인 구간 초기화
	for i := 0; i < n; i++ {
		sparse[0][i] = arr[i]
	}

	// 길이 2^k인 구간을 점화식으로 채운다
	maxK := logTable[n] + 1
	for k := 1; k < maxK; k++ {
		for i := 0; i+(1<<k)-1 < n; i++ {
			left := sparse[k-1][i]
			right := sparse[k-1][i+(1<<(k-1))]
			if left <= right {
				sparse[k][i] = left
			} else {
				sparse[k][i] = right
			}
		}
	}

	// 쿼리 처리
	for q := 0; q < m; q++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		// 1-indexed를 0-indexed로 변환한다
		l--
		r--
		// 구간 길이에 해당하는 로그 값을 구한다
		length := r - l + 1
		k := logTable[length]
		// 겹치는 두 구간의 최솟값을 합친다
		left := sparse[k][l]
		right := sparse[k][r-(1<<k)+1]
		if left <= right {
			fmt.Fprintln(writer, left)
		} else {
			fmt.Fprintln(writer, right)
		}
	}
}
