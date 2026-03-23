package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n    int
	arr  []int
	temp []int
)

// mergeSortCount 함수는 병합 정렬을 수행하면서 역전의 수를 센다
func mergeSortCount(lo, hi int) int64 {
	// 기저 조건: 원소가 하나이면 역전 없음
	if lo >= hi {
		return 0
	}

	mid := (lo + hi) / 2

	// 왼쪽과 오른쪽 부분의 역전 수를 재귀적으로 계산
	count := mergeSortCount(lo, mid)
	count += mergeSortCount(mid+1, hi)

	// 병합하면서 교차 역전 수를 계산
	count += mergeCount(lo, mid, hi)

	return count
}

// mergeCount 함수는 두 정렬된 부분 배열을 병합하면서 역전 수를 센다
func mergeCount(lo, mid, hi int) int64 {
	var count int64
	i, j, k := lo, mid+1, lo

	// 두 부분 배열을 비교하며 병합
	for i <= mid && j <= hi {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			// arr[i] > arr[j]이면 arr[i..mid]의 모든 원소가 arr[j]보다 크다
			count += int64(mid - i + 1)
			temp[k] = arr[j]
			j++
		}
		k++
	}

	// 남은 원소 복사
	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}
	for j <= hi {
		temp[k] = arr[j]
		j++
		k++
	}

	// 임시 배열의 결과를 원래 배열에 복사
	for idx := lo; idx <= hi; idx++ {
		arr[idx] = temp[idx]
	}

	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	fmt.Fscan(reader, &n)
	arr = make([]int, n)
	temp = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 병합 정렬 기반 역전 카운트
	result := mergeSortCount(0, n-1)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
