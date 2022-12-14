package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

func testPrint(minx, maxx, miny, maxy int, walls map[string]bool, sands map[string]bool, path map[string]bool, entry []int) {
	for j := miny; j <= maxy; j++ {
		for i := minx; i <= maxx; i++ {
			if has(walls, fmt.Sprintf("%d %d", i, j)) {
				fmt.Printf("#")
			} else if has(sands, fmt.Sprintf("%d %d", i, j)) {
				fmt.Printf(Yellow("o"))
			} else if has(path, fmt.Sprintf("%d %d", i, j)) {
				fmt.Printf(Purple("~"))
			} else if i == entry[0] && j == entry[1] {
				fmt.Printf(Green("S"))
			} else {
				fmt.Printf(Black("."))
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func has(m map[string]bool, k string) bool {
	_, ok := m[k]
	return ok
}

func part1(minx, maxx, miny, maxy int, walls map[string]bool) {
	sands := map[string]bool{}
	path := map[string]bool{}
	entry := []int{500, 0}
	cnt := 0

	sand := []int{entry[0], entry[1]}
	for {
		if sand[0] < minx || sand[0] > maxx || sand[1] > maxy {
			break
		}

		if !has(walls, fmt.Sprintf("%d %d", sand[0], sand[1]+1)) && !has(sands, fmt.Sprintf("%d %d", sand[0], sand[1]+1)) {
			path[fmt.Sprintf("%d %d", sand[0], sand[1])] = true
			sand[1]++
		} else if !has(walls, fmt.Sprintf("%d %d", sand[0]-1, sand[1]+1)) && !has(sands, fmt.Sprintf("%d %d", sand[0]-1, sand[1]+1)) {
			path[fmt.Sprintf("%d %d", sand[0], sand[1])] = true
			sand[1]++
			sand[0]--
		} else if !has(walls, fmt.Sprintf("%d %d", sand[0]+1, sand[1]+1)) && !has(sands, fmt.Sprintf("%d %d", sand[0]+1, sand[1]+1)) {
			path[fmt.Sprintf("%d %d", sand[0], sand[1])] = true
			sand[1]++
			sand[0]++
		} else {
			sands[fmt.Sprintf("%d %d", sand[0], sand[1])] = true
			sand = []int{entry[0], entry[1]}
			path = map[string]bool{}
			cnt++
		}
	}
	testPrint(minx, maxx, miny, maxy, walls, sands, path, entry)

	fmt.Printf("%+v\n", cnt)
}

func part2(minx, maxx, miny, maxy int, walls map[string]bool) {
	maxy += 2

	for i := minx; i <= maxx; i++ {
		walls[fmt.Sprintf("%d %d", i, maxy)] = true
	}

	sands := map[string]bool{}
	path := map[string]bool{}
	entry := []int{500, 0}
	cnt := 0

	sand := []int{entry[0], entry[1]}
	for {
		if sand[0] < minx || sand[0] > maxx || sand[1] > maxy {
			if sand[0] > maxx {
				maxx++
				walls[fmt.Sprintf("%d %d", maxx, maxy)] = true
			}
			if sand[0] < minx {
				minx--
				walls[fmt.Sprintf("%d %d", minx, maxy)] = true
			}
			sand = []int{entry[0], entry[1]}
			path = map[string]bool{}
		}

		if !has(walls, fmt.Sprintf("%d %d", sand[0], sand[1]+1)) && !has(sands, fmt.Sprintf("%d %d", sand[0], sand[1]+1)) {
			path[fmt.Sprintf("%d %d", sand[0], sand[1])] = true
			sand[1]++
		} else if !has(walls, fmt.Sprintf("%d %d", sand[0]-1, sand[1]+1)) && !has(sands, fmt.Sprintf("%d %d", sand[0]-1, sand[1]+1)) {
			path[fmt.Sprintf("%d %d", sand[0], sand[1])] = true
			sand[1]++
			sand[0]--
		} else if !has(walls, fmt.Sprintf("%d %d", sand[0]+1, sand[1]+1)) && !has(sands, fmt.Sprintf("%d %d", sand[0]+1, sand[1]+1)) {
			path[fmt.Sprintf("%d %d", sand[0], sand[1])] = true
			sand[1]++
			sand[0]++
		} else {
			sands[fmt.Sprintf("%d %d", sand[0], sand[1])] = true
			path = map[string]bool{}
			cnt++
			if sand[0] == entry[0] && sand[1] == entry[1] {
				break
			}
			sand = []int{entry[0], entry[1]}
		}
	}
	testPrint(minx, maxx, miny, maxy, walls, sands, path, entry)

	fmt.Printf("%+v\n", cnt)
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error opening input file: %+v\n", err)
		return
	}

	minx := 500
	miny := 0
	maxx := 500
	maxy := 0

	walls := map[string]bool{}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if len(line) < 1 {
			continue
		}

		prev := []int{}
		verts := strings.Split(line, " -> ")
		for _, vert := range verts {
			coords := strings.Split(vert, ",")
			pos := []int{}

			for _, coord := range coords {
				tmp, err := strconv.Atoi(coord)
				if err != nil {
					fmt.Printf("failed to convert num: %s", coord)
					continue
				}
				pos = append(pos, tmp)
			}

			if len(prev) > 0 {
				for i := int(math.Min(float64(prev[0]), float64(pos[0]))); i <= int(math.Max(float64(prev[0]), float64(pos[0]))); i++ {
					for j := int(math.Min(float64(prev[1]), float64(pos[1]))); j <= int(math.Max(float64(prev[1]), float64(pos[1]))); j++ {
						walls[fmt.Sprintf("%d %d", i, j)] = true
						minx = int(math.Min(float64(i), float64(minx)))
						maxx = int(math.Max(float64(i), float64(maxx)))
						miny = int(math.Min(float64(j), float64(miny)))
						maxy = int(math.Max(float64(j), float64(maxy)))
					}
				}
			}
			prev = pos
		}
	}

	part1(minx, maxx, miny, maxy, walls)
	part2(minx, maxx, miny, maxy, walls)
}
