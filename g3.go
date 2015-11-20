package main

import (
	"bufio"
	"fmt"
	stack "github.com/aasa11/gtest/stack"
	"os"
	"strconv"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var st stack.Stack

func main() {
	for {
		s, err := reader.ReadString('\n')
		var token string
		if err != nil {
			fmt.Println("err", err)
			return
		}
		for _, c := range s {
			switch {
			case c >= '0' && c <= '9':
				token += string(c)
			case c == ' ':
				if token == "" {
					continue
				}

				r, _ := strconv.Atoi(token)
				st.Push(r)
				token = ""
				st.Pt()
			case c == '+':
				a := st.Pop() + st.Pop()
				fmt.Println(a)
				st.Push(a)
				st.Pt()
			case c == '*':
				a := st.Pop() * st.Pop()
				fmt.Println(a)
				st.Push(a)
				st.Pt()
			case c == '-':
				b := st.Pop()
				a := st.Pop()
				c := a - b
				fmt.Println(c)
				st.Push(c)
				st.Pt()
			case c == '/':
				b := st.Pop()
				a := st.Pop()
				c := a / b
				fmt.Println(c)
				st.Push(c)
				st.Pt()
			case c == 'q':
				fmt.Println("bye")
				return
			default:
				//return
			}
		}
	}
}
