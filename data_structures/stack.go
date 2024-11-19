// Stack implementation in Go
// https://github.com/Hemant-Jain-Author/Data-Structures-Algorithms-In-Go/blob/master/Stack/Stack.go

package main

import "fmt"

type Stack struct {
	stk []any
}

func (s *Stack) Push(data interface{}) {
	s.stk = append(s.stk, data)
}

func (s *Stack) Pop() interface{} {
	n := len(s.stk)
	value := s.stk[n-1]
	s.stk = s.stk[:n-1]
	return value
}

func (s *Stack) Top() interface{} {
	n := len(s.stk)
	return s.stk[n-1]
}

func (s Stack) Size() int {
	return len(s.stk)
}

func (s Stack) IsEmpty() bool {
	return len(s.stk) == 0
}

func (s Stack) Print() {
	fmt.Println(s.stk)
}

func main() {
	stk := &Stack{}
	stk.Push("1x")
	stk.Push("2x")
	stk.Push("3x")
	fmt.Println(stk.Top())
	fmt.Println(stk.Size())
	fmt.Println(stk.IsEmpty())
	stk.Print()
	stk.Pop()
	stk.Print()
	stk.Pop()
	stk.Print()
	stk.Top()
}
