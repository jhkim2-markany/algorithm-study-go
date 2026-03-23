package main

import (
	"bufio"
	"fmt"
	"os"
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
func customSort(students []Student) []Student {
	// 여기에 코드를 작성하세요
	return nil
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
