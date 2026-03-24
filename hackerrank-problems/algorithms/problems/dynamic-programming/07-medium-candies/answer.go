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
//
// [알고리즘 힌트]
//
//	왼쪽→오른쪽, 오른쪽→왼쪽 두 번 스캔하여
//	양쪽 이웃 조건을 모두 만족하는 최소 사탕 수를 구한다.
func candies(n int, arr []int) int64 {
	// 모든 학생에게 1개씩 배정
	candy := make([]int, n)
	for i := range candy {
		candy[i] = 1
	}

	// 왼쪽에서 오른쪽으로 스캔: 왼쪽 이웃보다 성적이 높으면 +1
	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			candy[i] = candy[i-1] + 1
		}
	}

	// 오른쪽에서 왼쪽으로 스캔: 오른쪽 이웃보다 성적이 높으면 최댓값 갱신
	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			if candy[i+1]+1 > candy[i] {
				candy[i] = candy[i+1] + 1
			}
		}
	}

	// 전체 합 계산
	var total int64
	for _, c := range candy {
		total += int64(c)
	}

	return total
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
