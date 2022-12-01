package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

	elfs := []int{0}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		cal, err := strconv.Atoi(line)
		if err != nil {
			elfs = append(elfs, 0)
		} else {
			elfs[len(elfs)-1] += cal
		}
	}
	sort.Ints(elfs)

	fmt.Printf("Elfs: %+v\n", elfs)
	fmt.Printf("Max: %+v\n", elfs[len(elfs)-1])
	fmt.Printf("Top3: %+v\n", elfs[len(elfs)-1]+elfs[len(elfs)-2]+elfs[len(elfs)-3])
}
