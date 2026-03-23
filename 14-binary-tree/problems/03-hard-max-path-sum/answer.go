package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// maxPathSum은 이진 트리에서 최대 경로 합을 반환한다.
//
// [매개변수]
//   - val: 각 노드의 값 배열 (1-indexed)
//   - left: 각 노드의 왼쪽 자식 배열 (-1이면 자식 없음)
//   - right: 각 노드의 오른쪽 자식 배열 (-1이면 자식 없음)
//   - root: 루트 노드 번호
//
// [반환값]
//   - int: 트리에서 임의의 경로의 최대 합
//
// [알고리즘 힌트]
//
//	후위 순회 기반 트리 DP를 사용한다.
//	각 노드에서 왼쪽/오른쪽 서브트리의 최대 이득을 구하고 (음수면 0),
//	현재 노드를 꼭짓점으로 하는 경로 합(왼쪽 이득 + 노드 값 + 오른쪽 이득)으로
//	전역 최댓값을 갱신한다.
//	부모에게는 한쪽 방향만 선택하여 반환한다.
func maxPathSum(val, left, right []int, root int) int {
	ans := math.MinInt64

	var maxGain func(node int) int
	maxGain = func(node int) int {
		if node == -1 {
			return 0
		}

		leftGain := maxGain(left[node])
		if leftGain < 0 {
			leftGain = 0
		}
		rightGain := maxGain(right[node])
		if rightGain < 0 {
			rightGain = 0
		}

		pathSum := val[node] + leftGain + rightGain
		if pathSum > ans {
			ans = pathSum
		}

		if leftGain > rightGain {
			return val[node] + leftGain
		}
		return val[node] + rightGain
	}

	maxGain(root)
	return ans
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 노드 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 노드 값 입력
	val := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &val[i])
	}

	// 왼쪽/오른쪽 자식 배열 초기화
	left := make([]int, n+1)
	right := make([]int, n+1)

	// 각 노드의 자식 정보 입력
	for i := 0; i < n; i++ {
		var node, l, r int
		fmt.Fscan(reader, &node, &l, &r)
		left[node] = l
		right[node] = r
	}

	// 핵심 함수 호출
	result := maxPathSum(val, left, right, 1)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
