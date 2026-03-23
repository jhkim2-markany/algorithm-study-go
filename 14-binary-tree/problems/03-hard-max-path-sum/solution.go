package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// 이진 트리의 최대 경로 합 - 후위 순회 기반 트리 DP
// 시간 복잡도: O(N)
// 공간 복잡도: O(N)

var val []int
var left []int
var right []int
var ans int

// maxGain 함수는 해당 노드를 포함하여 아래로 내려가는 최대 경로 합을 반환한다
// 동시에 해당 노드를 꼭짓점으로 하는 경로의 합으로 전역 최댓값을 갱신한다
func maxGain(node int) int {
	if node == -1 {
		return 0
	}

	// 왼쪽/오른쪽 서브트리에서 얻을 수 있는 최대 이득 (음수면 0으로 처리)
	leftGain := maxGain(left[node])
	if leftGain < 0 {
		leftGain = 0
	}
	rightGain := maxGain(right[node])
	if rightGain < 0 {
		rightGain = 0
	}

	// 현재 노드를 꼭짓점으로 하는 경로의 합
	// 왼쪽에서 올라와서 현재 노드를 거쳐 오른쪽으로 내려가는 경로
	pathSum := val[node] + leftGain + rightGain
	if pathSum > ans {
		ans = pathSum
	}

	// 부모에게 반환하는 값: 현재 노드 + 왼쪽/오른쪽 중 더 큰 쪽
	// 부모로 올라가는 경로이므로 한쪽만 선택해야 한다
	if leftGain > rightGain {
		return val[node] + leftGain
	}
	return val[node] + rightGain
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 노드 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 노드 값 입력
	val = make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &val[i])
	}

	// 왼쪽/오른쪽 자식 배열 초기화
	left = make([]int, n+1)
	right = make([]int, n+1)

	// 각 노드의 자식 정보 입력
	for i := 0; i < n; i++ {
		var node, l, r int
		fmt.Fscan(reader, &node, &l, &r)
		left[node] = l
		right[node] = r
	}

	// 전역 최댓값을 음의 무한대로 초기화
	ans = math.MinInt64

	// 루트에서 후위 순회 기반 탐색 시작
	maxGain(1)

	fmt.Fprintln(writer, ans)
}
