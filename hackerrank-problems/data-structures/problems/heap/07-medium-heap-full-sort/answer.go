package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// heapify는 인덱스 i에서 최대 힙 성질을 복원한다.
//
// [매개변수]
//   - arr: 배열
//   - n: 힙 크기
//   - i: 시작 인덱스
func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// 왼쪽 자식이 루트보다 크면 갱신
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// 오른쪽 자식이 현재 최대보다 크면 갱신
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// 최대가 루트가 아니면 교환 후 재귀 호출
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// heapSort는 힙 정렬을 수행하며 각 단계의 배열 상태를 반환한다.
//
// [매개변수]
//   - arr: 정렬할 정수 배열
//
// [반환값]
//   - [][]int: 각 교환 후 배열 상태
//
// [알고리즘 힌트]
//
//	1단계: Build Max Heap — 배열의 중간부터 역순으로 Heapify
//	2단계: 루트(최댓값)를 끝으로 보내고 힙 크기를 줄여가며 정렬
func heapSort(arr []int) [][]int {
	n := len(arr)
	var steps [][]int

	// 1단계: 최대 힙 구성 (Build Max Heap)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 2단계: 루트를 끝으로 보내며 정렬
	for i := n - 1; i > 0; i-- {
		// 루트(최댓값)와 마지막 원소 교환
		arr[0], arr[i] = arr[i], arr[0]

		// 힙 크기를 줄이고 루트에서 Heapify
		heapify(arr, i, 0)

		// 현재 배열 상태 저장
		snapshot := make([]int, n)
		copy(snapshot, arr)
		steps = append(steps, snapshot)
	}

	return steps
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 배열 크기 입력
	var n int
	fmt.Fscan(reader, &n)

	// 배열 입력
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	steps := heapSort(arr)

	// 결과 출력
	for _, step := range steps {
		strs := make([]string, len(step))
		for i, v := range step {
			strs[i] = fmt.Sprintf("%d", v)
		}
		fmt.Fprintln(writer, strings.Join(strs, " "))
	}
}
