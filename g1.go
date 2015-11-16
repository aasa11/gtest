package main

//learning golang exercise : charpter 1

import (
	"fmt"
	"unicode/utf8"
)

func q11() {
	fmt.Println("----q11----")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func q12() {
	fmt.Println("----q12----")
	i := 0
LABEL:
	if i >= 10 {
		return
	}

	fmt.Println(i)
	i += 1
	goto LABEL
}

func q13() {
	fmt.Println("----q13----")
	arr := [...]int{1, 2, 354, 6, 7, 6, 544, 5, 234, 5, 3}
	for _, v := range arr {
		fmt.Println(v)
	}
}

func q21() {
	fmt.Println("----g21")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("FizzBuzz ")
		} else if i%3 == 0 {
			fmt.Printf("Fizz ")
		} else if i%5 == 0 {
			fmt.Printf("Buzz ")
		} else {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("")
}

func q31() {
	fmt.Println("----q31----")
	s := "A"
	for i := 0; i < 100; i++ {
		fmt.Println(s)
		s += "A"
	}
}

func q32() {
	fmt.Println("----q32----")
	s := "fasdfas ewwe dfd tr dk"
	fmt.Println("s len", len([]byte(s)))
	fmt.Println(" s rune len", utf8.RuneCount([]byte(s)))
}

func q33() {
	fmt.Println("----q33----")
	s := "fasdfas ewwe dfd tr dk"
	ss := []rune(s)
	fmt.Printf("%s\n", string(ss))
	copy(ss[4:4+3], []rune("abc"))
	fmt.Printf("%s\n", string(ss))
}

func q34() {
	fmt.Println("----q4----")
	s := "fadfasf afdaer btyytjhj gfdg"
	ss := []rune(s)
	fmt.Printf("%s\n", string(ss))
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
	fmt.Printf("%s\n", string(ss))
}

func q41() {
	a := [...]float64{123.534, 6746.8678, 4234.756, 4234.867, 9789.543}

	sum := 0.0
	for _, j := range a {
		sum += j
	}
	fmt.Printf("%v\n", sum/float64(len(a)))

}

func main() {
	fmt.Println("This is the test g1.go")
	//q11()
	//q12()
	//q13()
	//q21()
	//q31()
	//q32()
	//q33()
	//q34()
	q41()
}
