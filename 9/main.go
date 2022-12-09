package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func resPrint(visited map[string]bool) {
	maxx := 0
	minx := 0
	maxy := 0
	miny := 0

	for key, _ := range visited {
		pos := strings.Split(key, " ")
		x, _ := strconv.Atoi(pos[0])
		y, _ := strconv.Atoi(pos[1])
		if x < minx {
			minx = x
		}
		if y < miny {
			miny = y
		}
		if x > maxx {
			maxx = x
		}
		if y > maxy {
			maxy = y
		}
	}

	for i := maxy; i >= miny; i-- {
		for j := minx; j <= maxx; j++ {
			if _, ok := visited[fmt.Sprintf("%d %d", j, i)]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func testPrint(visited map[string]bool, rope [][]int) {
	maxx := 0
	minx := 0
	maxy := 0
	miny := 0

	for key, _ := range visited {
		pos := strings.Split(key, " ")
		x, _ := strconv.Atoi(pos[1])
		y, _ := strconv.Atoi(pos[0])
		if x < minx {
			minx = x
		}
		if y < miny {
			miny = y
		}
		if x > maxx {
			maxx = x
		}
		if y > maxy {
			maxy = y
		}
	}

	fmt.Printf("%+v\n", rope)
	pos := map[string]bool{}
	for k := 0; k < len(rope); k++ {
		pos[fmt.Sprintf("%d %d", rope[k][0], rope[k][1])] = true
		if rope[k][0] < minx {
			minx = rope[k][0]
		}
		if rope[k][0] > maxx {
			maxx = rope[k][0]
		}
		if rope[k][1] < miny {
			miny = rope[k][1]
		}
		if rope[k][1] > maxy {
			maxy = rope[k][1]
		}
	}
	for i := maxy; i >= miny; i-- {
		for j := minx; j <= maxx; j++ {
			if _, ok := pos[fmt.Sprintf("%d %d", j, i)]; ok {
				for k := 0; k < len(rope); k++ {
					if rope[k][0] == j && rope[k][1] == i {
						fmt.Printf("%d", k)
						break
					}
				}
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error opening input file: %+v\n", err)
		return
	}

	rlen := 10
	rope := [][]int{}
	visited := map[string]bool{}
	for i := 0; i < rlen; i++ {
		rope = append(rope, []int{0, 0})
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		visited[fmt.Sprintf("%d %d", rope[len(rope)-1][0], rope[len(rope)-1][1])] = true
		if len(line) < 1 {
			continue
		}

		cmds := strings.Split(line, " ")
		cmd := cmds[0]
		num, err := strconv.Atoi(cmds[1])
		if err != nil {
			continue
		}

		for i := 0; i < num; i++ {
			switch cmd {
			case "R":
				rope[0][0]++
			case "L":
				rope[0][0]--
			case "D":
				rope[0][1]--
			case "U":
				rope[0][1]++
			}
			for j := 1; j < rlen; j++ {
				difx := rope[j-1][0] - rope[j][0]
				dify := rope[j-1][1] - rope[j][1]

				if math.Abs(float64(dify)) > 1 || math.Abs(float64(difx)) > 1 {
					rope[j][0] += int(math.Max(-1, math.Min(1, float64(difx))))
					rope[j][1] += int(math.Max(-1, math.Min(1, float64(dify))))
				}
			}
			visited[fmt.Sprintf("%d %d", rope[len(rope)-1][0], rope[len(rope)-1][1])] = true
			// testPrint(visited, rope)
		}
	}

	// fmt.Printf("%+v\n", visited)
	fmt.Printf("%+v\n", len(visited))
	// resPrint(visited)
}
