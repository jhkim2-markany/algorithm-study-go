package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// 학생 정보를 저장하는 구조체
type Student struct {
	name           string
	kor, eng, math int
}

// customSort는 학생 배열을 다중 조건에 따라 정렬하여 반환한다.
//
// [매개변수]
//   - students: 학생 정보 배열 (각 학생은 이름, 국어, 영어, 수학 점수를 가짐)
//
// [반환값]
//   - []Student: 다중 조건으로 정렬된 학생 배열
//
// [알고리즘 힌트]
//
//	sort.SliceStable을 사용하여 안정 정렬을 수행한다.
//	정렬 조건은 우선순위 순서대로 비교한다:
//	  1. 국어 점수 내림차순 (높은 점수가 앞)
//	  2. 영어 점수 오름차순 (낮은 점수가 앞)
//	  3. 수학 점수 내림차순 (높은 점수가 앞)
//	  4. 이름 사전순 오름차순
//
//	다중 조건 비교 시, 현재 기준이 같으면 다음 기준으로 넘어간다.
//	시간복잡도: O(N log N)
//
//	예시: [{bob 90 70 80}, {alice 90 80 70}]
//	  → 국어 같음(90) → 영어 비교: bob(70) < alice(80) → bob이 앞
func customSort(students []Student) []Student {
	// 다중 조건 정렬
	sort.SliceStable(students, func(i, j int) bool {
		a, b := students[i], students[j]
		// 1. 국어 점수 내림차순
		if a.kor != b.kor {
			return a.kor > b.kor
		}
		// 2. 영어 점수 오름차순
		if a.eng != b.eng {
			return a.eng < b.eng
		}
		// 3. 수학 점수 내림차순
		if a.math != b.math {
			return a.math > b.math
		}
		// 4. 이름 사전순 오름차순
		return a.name < b.name
	})
	return students
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 학생 수 입력
	var n int
	fmt.Fscan(reader, &n)

	// 학생 정보 입력
	students := make([]Student, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &students[i].name, &students[i].kor, &students[i].eng, &students[i].math)
	}

	// 핵심 함수 호출
	sorted := customSort(students)

	// 결과 출력
	for i := 0; i < len(sorted); i++ {
		fmt.Fprintln(writer, sorted[i].name)
	}
}
