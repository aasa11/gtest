package main

import "fmt"

func main() {
	list := []string{"a", "b", "c", "a", "d", "d", "d", "f", "e"}
	first := list[0]

	fmt.Printf("%s ", first)
	for _, v := range list[1:] {
		if first != v {
			fmt.Printf("%s ", v)
			first = v
		}
	}
	fmt.Println()
}
