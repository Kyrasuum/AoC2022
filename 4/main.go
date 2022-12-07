package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var ()

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error opening input file: %+v\n", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	num := 0
	for _, line := range lines {
		if len(line) < 3 {
			continue
		}
		sections := strings.Split(line, ",")
		strpos := append(strings.Split(sections[0], "-"), strings.Split(sections[1], "-")...)
		pos := []int{}
		for _, str := range strpos {
			i, _ := strconv.Atoi(str)
			pos = append(pos, i)
		}
		if (pos[0] >= pos[2] && pos[0] <= pos[3]) || (pos[1] >= pos[2] && pos[1] <= pos[3] || (pos[0] <= pos[2] && pos[1] >= pos[3]) || (pos[0] >= pos[2] && pos[1] <= pos[3])) {
			num++
		}
	}
	fmt.Printf("%+v\n", num)
}
