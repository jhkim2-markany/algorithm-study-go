package main

import (
	"bufio"
	"fmt"
	"os"
)

// TextEditor는 undo 기능을 지원하는 간단한 텍스트 편집기이다.
type TextEditor struct {
	text    string
	history []string
}

// Append는 문자열 w를 현재 텍스트 끝에 추가한다.
//
// [매개변수]
//   - w: 추가할 문자열
func (e *TextEditor) Append(w string) {
	// 여기에 코드를 작성하세요
}

// Delete는 현재 텍스트의 마지막 k개 문자를 삭제한다.
//
// [매개변수]
//   - k: 삭제할 문자 수
func (e *TextEditor) Delete(k int) {
	// 여기에 코드를 작성하세요
}

// Print는 현재 텍스트의 k번째 문자를 반환한다 (1-indexed).
//
// [매개변수]
//   - k: 문자 위치 (1-indexed)
//
// [반환값]
//   - byte: k번째 문자
func (e *TextEditor) Print(k int) byte {
	// 여기에 코드를 작성하세요
	return 0
}

// Undo는 마지막 연산을 취소한다.
func (e *TextEditor) Undo() {
	// 여기에 코드를 작성하세요
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 연산 개수 입력
	var q int
	fmt.Fscan(reader, &q)

	editor := &TextEditor{}

	for i := 0; i < q; i++ {
		var op int
		fmt.Fscan(reader, &op)
		switch op {
		case 1:
			var w string
			fmt.Fscan(reader, &w)
			editor.Append(w)
		case 2:
			var k int
			fmt.Fscan(reader, &k)
			editor.Delete(k)
		case 3:
			var k int
			fmt.Fscan(reader, &k)
			fmt.Fprintf(writer, "%c\n", editor.Print(k))
		case 4:
			editor.Undo()
		}
	}
}
