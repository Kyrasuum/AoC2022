package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	cycle  = 0
	x      = 1
	signal = 0
)

func nextCycle() {
	if x-1 <= cycle%40 && x+1 >= cycle%40 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
	cycle++
	if (cycle+20)%40 == 0 {
		signal += cycle * x
	}
	if (cycle)%40 == 0 {
		fmt.Println()
	}
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error opening input file: %+v\n", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if len(line) < 1 {
			continue
		}

		cmds := strings.Split(line, " ")
		switch cmds[0] {
		case "noop":
			nextCycle()
		case "addx":
			nextCycle()
			num, err := strconv.Atoi(cmds[1])
			if err != nil {
				continue
			}
			nextCycle()
			x += num
		}
	}
	fmt.Printf("%d\n", signal)
}
