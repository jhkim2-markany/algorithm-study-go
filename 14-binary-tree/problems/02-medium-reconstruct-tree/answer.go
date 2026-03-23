package main

import (
	"bufio"
	"fmt"
	"os"
)

// reconstructPostorder는 전위 순회와 중위 순회 결과로 후위 순회 결과를 복원한다.
//
// [매개변수]
//   - preorder: 전위 순회 결과 배열
//   - inorder: 중위 순회 결과 배열
//
// [반환값]
//   - []int: 후위 순회 결과 배열
//
// [알고리즘 힌트]
//
//	전위 순회의 첫 번째 원소가 현재 서브트리의 루트이다.
//	중위 순회에서 루트의 위치를 찾아 왼쪽/오른쪽 서브트리를 분할한다.
//	후위 순회 순서(왼쪽 → 오른쪽 → 루트)로 재귀 호출하여 결과를 구성한다.
//	중위 순회의 인덱스를 맵에 저장하면 O(1)로 루트 위치를 찾을 수 있다.
func reconstructPostorder(preorder, inorder []int) []int {
	inIdx := make(map[int]int)
	for i, v := range inorder {
		inIdx[v] = i
	}

	var result []int
	var build func(preL, preR, inL, inR int)
	build = func(preL, preR, inL, inR int) {
		if preL > preR {
			return
		}
		root := preorder[preL]
		rootInIdx := inIdx[root]
		leftSize := rootInIdx - inL

		build(preL+1, preL+leftSize, inL, rootInIdx-1)
		build(preL+leftSize+1, preR, rootInIdx+1, inR)
		result = append(result, root)
	}

	n := len(preorder)
	build(0, n-1, 0, n-1)
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 노드 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 전위 순회 결과 입력
	preorder := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &preorder[i])
	}

	// 중위 순회 결과 입력
	inorder := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &inorder[i])
	}

	// 핵심 함수 호출
	result := reconstructPostorder(preorder, inorder)

	// 결과 출력
	for i, v := range result {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, v)
	}
	fmt.Fprintln(writer)
}
