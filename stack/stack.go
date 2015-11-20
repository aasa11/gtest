package stack

import (
	//"errors"
	"fmt"
)

const STACK_COUNT int = 10

type Stack struct {
	data [STACK_COUNT]int
	idx  int
}

func (s *Stack) Push(a int) {
	if s.idx >= STACK_COUNT {
		return
	} else {
		s.data[s.idx] = a
		s.idx += 1
	}
	return
}

func (s *Stack) Pop() (a int) {
	if s.idx <= 0 {
		return
	} else {
		a = s.data[s.idx-1]
		s.data[s.idx-1] = 0
		s.idx -= 1
	}
	return
}

func (s *Stack) Count() (num int) {
	num = s.idx
	return
}

func (s *Stack) Pt() {
	fmt.Printf("Stack is : ")
	for i := 0; i < s.idx; i++ {
		fmt.Printf("%d ", s.data[i])
	}
	fmt.Println()
}
