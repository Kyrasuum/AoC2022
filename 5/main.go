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
	crates := [][]string{}
	for i, line := range lines {
		if len(line) < 1 {
			continue
		}
		if !strings.Contains(line, "move") {
			if !strings.Contains(line, "[") {
				cols := strings.Split(line, " ")
				for _, col := range cols {
					if len(col) > 0 {
						crates = append(crates, []string{})
						colnum, _ := strconv.Atoi(col)
						pos := strings.Index(line, col)
						for j, _ := range lines[:i] {
							line := lines[i-j-1]
							char := line[pos]
							if char != ' ' {
								crates[colnum-1] = append(crates[colnum-1], string(char))
							}
						}
					}
				}
				fmt.Printf("%+v\n", crates)
			}
		} else {
			cmds := strings.Split(line, " ")
			num, _ := strconv.Atoi(cmds[1])
			fi, _ := strconv.Atoi(cmds[3])
			ti, _ := strconv.Atoi(cmds[5])

			from := crates[fi-1]
			to := crates[ti-1]

			fmt.Printf("%+v   :   %+v\n", to, from)
			to = append(to, from[len(from)-num:]...)
			from = from[:len(from)-num]
			fmt.Printf("%+v   :   %+v\n", to, from)

			crates[fi-1] = from
			crates[ti-1] = to
		}
	}

	for _, crate := range crates {
		fmt.Printf("%+v", crate[len(crate)-1])
	}
}
