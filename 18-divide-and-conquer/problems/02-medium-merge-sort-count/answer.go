package main

import (
	"bufio"
	"fmt"
	"os"
)

// countInversions는 병합 정렬을 이용하여 배열의 역전 수를 반환한다.
//
// [매개변수]
//   - arr: 정수 배열 (함수 내에서 정렬됨)
//
// [반환값]
//   - int64: 역전의 수
//
// [알고리즘 힌트]
//
//	병합 정렬을 수행하면서 역전 수를 센다.
//	두 정렬된 부분 배열을 병합할 때, 오른쪽 원소가 먼저 선택되면
//	왼쪽에 남은 원소 수만큼 역전이 발생한다.
//	왼쪽/오른쪽 부분의 역전 수 + 교차 역전 수를 합산한다.
func countInversions(arr []int) int64 {
	temp := make([]int, len(arr))

	var mergeSortCount func(lo, hi int) int64
	mergeSortCount = func(lo, hi int) int64 {
		if lo >= hi {
			return 0
		}
		mid := (lo + hi) / 2
		count := mergeSortCount(lo, mid)
		count += mergeSortCount(mid+1, hi)

		i, j, k := lo, mid+1, lo
		for i <= mid && j <= hi {
			if arr[i] <= arr[j] {
				temp[k] = arr[i]
				i++
			} else {
				count += int64(mid - i + 1)
				temp[k] = arr[j]
				j++
			}
			k++
		}
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
		for idx := lo; idx <= hi; idx++ {
			arr[idx] = temp[idx]
		}
		return count
	}

	return mergeSortCount(0, len(arr)-1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력 처리
	var n int
	fmt.Fscan(reader, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	// 핵심 함수 호출
	result := countInversions(arr)

	// 결과 출력
	fmt.Fprintln(writer, result)
}
