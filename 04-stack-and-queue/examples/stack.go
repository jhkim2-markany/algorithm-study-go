package main

import "fmt"

// 스택(Stack) - 후입선출(LIFO) 자료구조 기본 구현
// 슬라이스를 활용하여 Push, Pop, Peek 연산을 구현한다.
// 시간 복잡도: 모든 연산 O(1)
// 공간 복잡도: O(N)

// Stack 구조체 정의
type Stack struct {
	data []int
}

// Push - 스택의 맨 위에 원소를 추가한다
func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

// Pop - 스택의 맨 위 원소를 제거하고 반환한다
func (s *Stack) Pop() int {
	n := len(s.data)
	val := s.data[n-1]
	s.data = s.data[:n-1]
	return val
}

// Peek - 스택의 맨 위 원소를 제거하지 않고 확인한다
func (s *Stack) Peek() int {
	return s.data[len(s.data)-1]
}

// IsEmpty - 스택이 비어 있는지 확인한다
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Size - 스택에 저장된 원소의 개수를 반환한다
func (s *Stack) Size() int {
	return len(s.data)
}

func main() {
	// 스택 생성 및 기본 연산 테스트
	s := &Stack{}

	fmt.Println("=== 스택 기본 연산 ===")

	// Push 연산: 원소 삽입
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Printf("Push 10, 20, 30 후 크기: %d\n", s.Size())
	fmt.Printf("Peek (맨 위 원소): %d\n", s.Peek())

	// Pop 연산: 원소 제거 (LIFO 순서로 나온다)
	fmt.Printf("Pop: %d\n", s.Pop())
	fmt.Printf("Pop: %d\n", s.Pop())
	fmt.Printf("Pop 후 크기: %d\n", s.Size())
	fmt.Printf("비어 있는가: %v\n", s.IsEmpty())

	// 마지막 원소 Pop
	fmt.Printf("Pop: %d\n", s.Pop())
	fmt.Printf("비어 있는가: %v\n", s.IsEmpty())

	// 스택 활용 예시: 문자열 뒤집기
	fmt.Println("\n=== 스택 활용: 문자열 뒤집기 ===")
	str := "HELLO"
	charStack := &Stack{}
	for _, ch := range str {
		charStack.Push(int(ch))
	}
	fmt.Printf("원본: %s\n", str)
	fmt.Print("뒤집기: ")
	for !charStack.IsEmpty() {
		fmt.Print(string(rune(charStack.Pop())))
	}
	fmt.Println()
}
