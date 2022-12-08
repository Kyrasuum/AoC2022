package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func part1(heights [][]int) {
	visible := [][]int{}
	count := 0

	for i, row := range heights {
		visible = append(visible, []int{})
		for range row {
			visible[i] = append(visible[i], 0)
		}
	}

	for i, row := range heights {
		max := -1
		for j, tree := range row {
			if tree > max {
				max = tree
				if visible[i][j] != 1 {
					visible[i][j] = 1
					count++
				}
			}
		}
		max = -1
		for j := len(row) - 1; j >= 0; j-- {
			tree := row[j]
			if tree > max {
				max = tree
				if visible[i][j] != 1 {
					visible[i][j] = 1
					count++
				}
			}
		}
	}
	for j := len(heights[0]) - 1; j >= 0; j-- {
		max := -1
		for i := 0; i < len(heights); i++ {
			row := heights[i]
			tree := row[j]
			if tree > max {
				max = tree
				if visible[i][j] != 1 {
					visible[i][j] = 1
					count++
				}
			}
		}
		max = -1
		for i := len(heights) - 1; i >= 0; i-- {
			row := heights[i]
			tree := row[j]
			if tree > max {
				max = tree
				if visible[i][j] != 1 {
					visible[i][j] = 1
					count++
				}
			}
		}
	}
	for _, row := range visible {
		fmt.Printf("%+v\n", row)
	}
	fmt.Printf("%+v\n", count)
}

func visLeft(heights [][]int, x int, y int) int {
	count := 0
	for i := x - 1; i >= 0; i-- {
		count++
		if heights[y][i] >= heights[y][x] {
			break
		}
	}
	return count
}

func visRight(heights [][]int, x int, y int) int {
	count := 0
	for i := x + 1; i < len(heights[y]); i++ {
		count++
		if heights[y][i] >= heights[y][x] {
			break
		}
	}
	return count
}

func visUp(heights [][]int, x int, y int) int {
	count := 0
	for i := y + 1; i < len(heights); i++ {
		count++
		if heights[i][x] >= heights[y][x] {
			break
		}
	}
	return count
}

func visDown(heights [][]int, x int, y int) int {
	count := 0
	for i := y - 1; i >= 0; i-- {
		count++
		if heights[i][x] >= heights[y][x] {
			break
		}
	}
	return count
}

func part2(heights [][]int) int {
	max := 0

	for y, row := range heights {
		for x, _ := range row {
			up := visUp(heights, x, y)
			dn := visDown(heights, x, y)
			lf := visLeft(heights, x, y)
			rt := visRight(heights, x, y)
			cur := up * dn * lf * rt
			if cur > max {
				fmt.Printf("\n\n%+v * %+v * %+v * %+v = %+v [%+v, %+v]\n", up, dn, lf, rt, cur, x, y)
				testPrint(heights, x, y)
				max = cur
			}
		}
	}
	return max
}

func testPrint(heights [][]int, x int, y int) {
	for j, row := range heights {
		fmt.Printf("[")
		for i, tree := range row {
			if i == x && j == y {
				fmt.Printf("%+v ", Red(tree))
			} else {
				fmt.Printf("%+v ", tree)
			}
		}
		fmt.Printf("]\n")
	}
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error opening input file: %+v\n", err)
		return
	}

	heights := [][]int{}

	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if len(line) < 1 {
			continue
		}
		heights = append(heights, []int{})

		for _, char := range line {
			height, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}
			heights[i] = append(heights[i], height)
		}
	}
	// part1(heights)
	fmt.Printf("%+v\n", part2(heights))
}
