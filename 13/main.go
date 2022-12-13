package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var ()

func matching(sub string) int {
	depth := 1
	index := 0
	for depth > 0 && index < len(sub) {
		switch sub[index] {
		case ']':
			depth--
		case '[':
			depth++
		}
		index++
	}
	//handle malformed arrays gracefully
	if depth != 0 {
		return index + 1
	}
	return index
}

func load(sub string) interface{} {
	//attempt to load number
	num, err := strconv.Atoi(sub)
	if err == nil {
		return num
	}

	i := 0
	arr := []interface{}{}
	//trim brackets
	sub = sub[1:matching(sub[1:])]
	for len(sub) > 0 {
		j := strings.Index(sub, ",")
		//handle sub arrays
		if sub[i] == '[' {
			j = matching(sub[1:]) + 1
		}
		//handle end element
		if j < 0 {
			j = len(sub)
		}
		//load sub element
		arr = append(arr, load(sub[i:j]))
		//advance to next element
		if len(sub) > j+1 {
			sub = sub[j+1:]
		} else {
			sub = sub[j:]
		}
	}
	return arr
}

func compair(a, b interface{}) int {
	switch a.(type) {
	case []interface{}:
		switch b.(type) {
		case []interface{}:
			var i int
			for i = 0; i < len(b.([]interface{})); i++ {
				if i >= len(a.([]interface{})) {
					return 1
				}
				val := compair(a.([]interface{})[i], b.([]interface{})[i])
				if val != 0 {
					return val
				}
			}
			if i < len(a.([]interface{})) {
				return -1
			}
		default:
			return compair(a, []interface{}{b})
		}
	default:
		switch b.(type) {
		case []interface{}:
			return compair([]interface{}{a}, b)
		default:
			if a.(int) > b.(int) {
				return -1
			}
			if a.(int) < b.(int) {
				return 1
			}
		}
	}
	return 0
}

func part1(comps []interface{}) {
	sum := 0
	for i := 0; i < len(comps); i += 2 {
		if compair(comps[i], comps[i+1]) == 1 {
			sum += i/2 + 1
		}
	}

	fmt.Printf("Part1: %+v\n", sum)
}

func part2(comps []interface{}) {
	comps = append(comps, load("[[2]]"))
	comps = append(comps, load("[[6]]"))

	sort.SliceStable(comps, func(i, j int) bool {
		return compair(comps[i], comps[j]) == 1
	})

	a := 0
	b := 0

	for i, comp := range comps {
		if fmt.Sprintf("%+v", comp) == "[[2]]" {
			a = i + 1
		}
		if fmt.Sprintf("%+v", comp) == "[[6]]" {
			b = i + 1
		}
	}
	fmt.Printf("Part2: %+v\n", a*b)
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error opening input file: %+v\n", err)
		return
	}

	comps := []interface{}{}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if len(line) < 1 {
			continue
		}

		comps = append(comps, load(line))
	}

	part1(comps)
	part2(comps)
}
