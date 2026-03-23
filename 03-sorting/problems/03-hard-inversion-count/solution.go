package main

import (
	"bufio"
	"fmt"
	"os"
)

// 역전 카운트를 세면서 병합 정렬을 수행하는 변수
var count int64

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

	// 병합 정렬을 수행하면서 역전 개수를 센다
	count = 0
	temp := make([]int, n)
	mergeSortCount(arr, temp, 0, n-1)

	// 결과 출력
	fmt.Fprintln(writer, count)
}

// mergeSortCount 함수는 병합 정렬을 수행하면서 역전 개수를 누적한다.
func mergeSortCount(arr, temp []int, left, right int) {
	if left >= right {
		return
	}

	mid := (left + right) / 2

	// 왼쪽 절반 정렬
	mergeSortCount(arr, temp, left, mid)
	// 오른쪽 절반 정렬
	mergeSortCount(arr, temp, mid+1, right)
	// 두 정렬된 부분을 병합하면서 역전 개수를 센다
	mergeCount(arr, temp, left, mid, right)
}

// mergeCount 함수는 두 정렬된 부분 배열을 병합하면서 역전 개수를 센다.
// 왼쪽 배열의 원소가 오른쪽 배열의 원소보다 클 때 역전이 발생한다.
func mergeCount(arr, temp []int, left, mid, right int) {
	i := left    // 왼쪽 배열 포인터
	j := mid + 1 // 오른쪽 배열 포인터
	k := left    // 임시 배열 포인터

	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			// 왼쪽 배열의 i번째 원소가 오른쪽 배열의 j번째 원소보다 크면
			// i부터 mid까지의 모든 원소가 arr[j]보다 크므로 역전 개수 추가
			count += int64(mid - i + 1)
			temp[k] = arr[j]
			j++
		}
		k++
	}

	// 왼쪽 배열의 남은 원소 복사
	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}

	// 오른쪽 배열의 남은 원소 복사
	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}

	// 임시 배열의 결과를 원래 배열에 복사
	for idx := left; idx <= right; idx++ {
		arr[idx] = temp[idx]
	}
}
