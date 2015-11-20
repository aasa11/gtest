package stack

import (
	"testing"
)

var s Stack

func TestStack1(t *testing.T) {
	for i := 0; i < 10; i++ {
		if s.Push(i) != nil {
			t.Log("stack add err")
			t.Fail()
		}
	}
}

func TestStack2(t *testing.T) {
	for j := 10; j < 15; j++ {
		if s.Push(j) == nil {
			t.Log("stack add fail")
			t.Fail()
		}
	}
}

func TestStack3(t *testing.T) {
	if s.Count() != 10 {
		t.Log("Stack count err")
		t.Fail()
	}
}

func TestStack4(t *testing.T) {
	for i := 0; i < 10; i++ {
		k, err := s.Pop()
		if k != 9-i || err != nil {
			t.Log("Stack pop err")
			t.Fail()
		}
	}
}

func TestStack5(t *testing.T) {
	for j := 10; j < 15; j++ {
		_, err := s.Pop()
		if err == nil {
			t.Log("stack pop fail")
			t.Fail()
		}
	}
}
