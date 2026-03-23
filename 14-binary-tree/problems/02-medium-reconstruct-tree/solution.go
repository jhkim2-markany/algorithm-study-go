package main

import (
	"bufio"
	"fmt"
	"os"
)

// 이진 트리 복원 - 전위 순회 + 중위 순회로 후위 순회 결과를 구한다
// 시간 복잡도: O(N)
// 공간 복잡도: O(N)

var preorder []int
var inorder []int
var inIdx map[int]int // 중위 순회에서 각 값의 인덱스
var result []int

// build 함수는 전위/중위 순회 결과로 트리를 복원하고 후위 순회를 수행한다
// preL, preR: 전위 순회에서의 범위
// inL, inR: 중위 순회에서의 범위
func build(preL, preR, inL, inR int) {
	if preL > preR {
		return
	}

	// 전위 순회의 첫 번째 원소가 현재 서브트리의 루트
	root := preorder[preL]

	// 중위 순회에서 루트의 위치를 찾는다
	rootInIdx := inIdx[root]

	// 왼쪽 서브트리의 크기
	leftSize := rootInIdx - inL

	// 왼쪽 서브트리 복원 (후위 순회이므로 왼쪽 먼저)
	build(preL+1, preL+leftSize, inL, rootInIdx-1)

	// 오른쪽 서브트리 복원
	build(preL+leftSize+1, preR, rootInIdx+1, inR)

	// 후위 순회: 자식을 모두 처리한 뒤 루트를 추가
	result = append(result, root)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 노드 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 전위 순회 결과 입력
	preorder = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &preorder[i])
	}

	// 중위 순회 결과 입력
	inorder = make([]int, n)
	inIdx = make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &inorder[i])
		// 중위 순회에서 각 값의 인덱스를 저장하여 O(1) 탐색
		inIdx[inorder[i]] = i
	}

	// 트리 복원 및 후위 순회 수행
	result = []int{}
	build(0, n-1, 0, n-1)

	// 결과 출력
	for i, v := range result {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
