package main

import (
	"fmt"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ps := exec.Command("ps", "-e", "-opid,ppid,comm")
	output, err := ps.Output()
	if err != nil {
		fmt.Println("err occured", err)
		return
	}
	child := make(map[int][]int)
	for i, s := range strings.Split(string(output), "\n") {
		if i == 0 {
			continue
		}
		if len(s) <= 0 {
			continue
		}
		f := strings.Fields(s)
		fpp, err1 := strconv.Atoi(f[1])
		if err1 != nil {
			fmt.Println("get ppid err", err, s)
			continue
		}
		fp, err2 := strconv.Atoi(f[0])
		if err2 != nil {
			fmt.Println("get pid err", err, s)
			continue
		}

		child[fpp] = append(child[fpp], fp)
	}

	schild := make([]int, len(child))
	i := 0
	for k, _ := range child {
		schild[i] = k
		i++
	}
	sort.Ints(schild)
	for _, ppid := range schild {
		fmt.Printf("Pid, %d has %d child", ppid, len(child[ppid]))
		fmt.Printf(" : %v\n", child[ppid])
	}
}
