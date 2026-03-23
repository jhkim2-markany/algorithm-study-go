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

	// 결과 출력
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, students[i].name)
	}
}
