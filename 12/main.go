package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

var ()

func Contains(s [][]int, e []int) bool {
	for _, a := range s {
		if fmt.Sprintf("%+v", a) == fmt.Sprintf("%+v", e) {
			return true
		}
	}
	return false
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error opening input file: %+v\n", err)
		return
	}

	heights := [][]int{}
	steps := [][]int{}
	start := []int{}
	ends := [][]int{}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if len(line) < 1 {
			continue
		}
		heights = append(heights, []int{})
		steps = append(steps, []int{})

		for _, char := range line {
			if char == 'S' {
				heights[len(heights)-1] = append(heights[len(heights)-1], 0)
				ends = append(ends, []int{len(heights) - 1, len(heights[len(heights)-1]) - 1})
				steps[len(steps)-1] = append(steps[len(steps)-1], -1)
			} else if char == 'a' {
				heights[len(heights)-1] = append(heights[len(heights)-1], 0)
				ends = append(ends, []int{len(heights) - 1, len(heights[len(heights)-1]) - 1})
				steps[len(steps)-1] = append(steps[len(steps)-1], -1)
			} else if char == 'E' {
				heights[len(heights)-1] = append(heights[len(heights)-1], int('z'-'a'))
				start = []int{len(heights) - 1, len(heights[len(heights)-1]) - 1}
				steps[len(steps)-1] = append(steps[len(steps)-1], 0)
			} else {
				heights[len(heights)-1] = append(heights[len(heights)-1], int(char-'a'))
				steps[len(steps)-1] = append(steps[len(steps)-1], -1)
			}
		}
	}

	pos := start
	slice := [][]int{start}
	for len(slice) != 0 {
		pos = slice[0]
		slice = slice[1:]

		//up
		if pos[0]-1 >= 0 && pos[0]-1 < len(heights) && heights[pos[0]][pos[1]]-1 <= heights[pos[0]-1][pos[1]] {
			//update steps
			if steps[pos[0]-1][pos[1]] == -1 || steps[pos[0]-1][pos[1]] > steps[pos[0]][pos[1]]+1 {
				steps[pos[0]-1][pos[1]] = steps[pos[0]][pos[1]] + 1
				//add to slice
				if !Contains(slice, []int{pos[0] - 1, pos[1]}) {
					slice = append(slice, []int{pos[0] - 1, pos[1]})
				}
			}
		}

		//down
		if pos[0]+1 >= 0 && pos[0]+1 < len(heights) && heights[pos[0]][pos[1]]-1 <= heights[pos[0]+1][pos[1]] {
			//update steps
			if steps[pos[0]+1][pos[1]] == -1 || steps[pos[0]+1][pos[1]] > steps[pos[0]][pos[1]]+1 {
				steps[pos[0]+1][pos[1]] = steps[pos[0]][pos[1]] + 1
				//add to slice
				if !Contains(slice, []int{pos[0] + 1, pos[1]}) {
					slice = append(slice, []int{pos[0] + 1, pos[1]})
				}
			}
		}

		//left
		if pos[1]-1 >= 0 && pos[1]-1 < len(heights[pos[0]]) && heights[pos[0]][pos[1]]-1 <= heights[pos[0]][pos[1]-1] {
			//update steps
			if steps[pos[0]][pos[1]-1] == -1 || steps[pos[0]][pos[1]-1] > steps[pos[0]][pos[1]]+1 {
				steps[pos[0]][pos[1]-1] = steps[pos[0]][pos[1]] + 1
				//add to slice
				if !Contains(slice, []int{pos[0], pos[1] - 1}) {
					slice = append(slice, []int{pos[0], pos[1] - 1})
				}
			}
		}

		//right
		if pos[1]+1 >= 0 && pos[1]+1 < len(heights[pos[0]]) && heights[pos[0]][pos[1]]-1 <= heights[pos[0]][pos[1]+1] {
			//update steps
			if steps[pos[0]][pos[1]+1] == -1 || steps[pos[0]][pos[1]+1] > steps[pos[0]][pos[1]]+1 {
				steps[pos[0]][pos[1]+1] = steps[pos[0]][pos[1]] + 1
				//add to slice
				if !Contains(slice, []int{pos[0], pos[1] + 1}) {
					slice = append(slice, []int{pos[0], pos[1] + 1})
				}
			}
		}

		//sort
		sort.SliceStable(slice, func(i, j int) bool {
			a := -1
			b := -1
			for _, end := range ends {
				if int(math.Abs(float64(slice[i][0]-end[0]))+math.Abs(float64(slice[i][1]-end[1]))) < a || a == -1 {
					a = int(math.Abs(float64(slice[i][0]-end[0])) + math.Abs(float64(slice[i][1]-end[1])))
				}
				if int(math.Abs(float64(slice[j][0]-end[0]))+math.Abs(float64(slice[j][1]-end[1]))) < b || b == -1 {
					b = int(math.Abs(float64(slice[j][0]-end[0])) + math.Abs(float64(slice[j][1]-end[1])))
				}
			}
			return a < b
		})
	}

	min := -1
	for _, end := range ends {
		if steps[end[0]][end[1]] != -1 && steps[end[0]][end[1]] < min || min == -1 {
			min = steps[end[0]][end[1]]
		}
	}
	fmt.Printf("\n%+v\n", min)
}
