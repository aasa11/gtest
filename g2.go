package main

//learning golang exercise chapter 2

import (
	"errors"
	"fmt"
	"strconv"
)

type empty interface {
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 0
}

func q50(xs []float64) (avg float64) {
	sum := 0.0
	if len(xs) <= 0 {
		return
	}

	for _, v := range xs {
		sum += v
	}

	avg = sum / float64(len(xs))
	return
}

func do_q50() {
	fmt.Println("----do_q50----")
	x := [...]float64{14.45, 45.534, 867.34, 654.534}
	var xs []float64 = x[:]
	avg := q50(xs)
	fmt.Println("avg", avg)
}

func q60(a, b int) (c, d int) {
	if a > b {
		c, d = b, a
	} else {
		c, d = a, b
	}
	return
}

func do_q60() {
	fmt.Println("----do_q60----")
	a, b := 112, 45
	c, d := q60(a, b)
	fmt.Println(a, b, c, d)
}

type MyStack struct {
	st  [10]int
	cnt int
}

func (s *MyStack) push(k int) (err error) {
	if s.cnt >= 10 {
		err = errors.New("Out of stack capacity")
		return
	}
	s.st[s.cnt] = k
	s.cnt += 1
	return
}

func (s *MyStack) pop() (k int, err error) {
	if s.cnt <= 0 {
		err = errors.New("stack is empty")
		return
	}
	k = s.st[s.cnt-1]
	s.st[s.cnt-1] = 0
	s.cnt -= 1
	return
}

func (s MyStack) String() (ss string) {
	for i := 0; i < s.cnt; i++ {
		ss += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.st[i]) + "] "
	}
	ss += "\n"
	return
}

func do_q80() {
	var s MyStack

	fmt.Printf("s %v\n", s)
	for i := 0; i < 15; i++ {
		err := s.push(i)
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Println("push ", i)
		}
	}

	fmt.Printf("s %v\n", s)
	for i := 0; i < 15; i++ {
		k, err := s.pop()
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Println("pop ", k)
		}
	}

	fmt.Printf("s %v\n", s)
}

func q90(arg ...int) {
	for _, v := range arg {
		fmt.Println(v)
	}
}

func do_q90() {
	fmt.Println("----do_q90----")
	xs := [...]int{1, 2, 34, 5, 6, 7, 5, 4, 3}
	q90(xs[:]...)
}

func q100(n int) []int {
	x := make([]int, n)
	x[0], x[1] = 1, 1
	for i := 2; i < n; i++ {
		x[i] = x[i-1] + x[i-2]
	}
	return x
}

func do_q100() {
	fmt.Println("----do_q100----")
	v := q100(10)
	fmt.Println(10, v)
}

func f110(a empty) empty {
	fmt.Println(a)
	return nil
}

func q110(s []empty, f func(empty) empty) []empty {
	slen := len(s)
	xs := make([]empty, slen)
	for i := 0; i < slen; i++ {
		xs[i] = f(s[i])
	}
	return xs
}

func do_q110() {
	fmt.Println("----do_q110----")
	s := []empty{1, "dfa", 3.54, 4, 'a', 6, 7}
	ss := q110(s, f110)
	fmt.Println(ss)
}

func q120(s []int) (m int) {
	m = s[0]
	for _, v := range s {
		if v < m {
			m = v
		}
	}
	return
}

func do_q120() {
	fmt.Println("----do_q120----")
	s := []int{123, 54, 123, 7, 8779, 54, 34, 13}
	v := q120(s)
	fmt.Println(s, v)
}

func q130(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}

func do_q130() {
	fmt.Println("----do_q130----")
	s := []int{23, 54, 764, 523, 444, 6, 7, 8, 54, 33, 1}
	fmt.Println(s)
	fmt.Println(q130(s))
	fmt.Println(s)
}

func q140(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

func do_q140() {
	fmt.Println("----do_q140----")
	f := q140(10)
	fmt.Println(f(5))
}

func main() {
	//do_q50()
	//do_q60()
	//do_q80()
	//do_q90()
	//do_q100()
	//do_q110()
	//do_q120()
	//do_q130()
	do_q140()
}
