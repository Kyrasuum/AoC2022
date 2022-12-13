package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var ()

type monkey struct {
	name  int
	items []int
	count int
	op    func(int) int
	test  func(int) bool
	pass  int
	fail  int
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error opening input file: %+v\n", err)
		return
	}

	lcm := 1
	monkeys := []monkey{}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if len(line) < 1 {
			continue
		}

		if strings.Index(line, "Monkey") == 0 {
			monkeys = append(monkeys, monkey{name: len(monkeys)})
			continue
		}
		if strings.Index(line, "Starting") == 2 {
			items := strings.Split(line[strings.Index(line, ": ")+2:], ", ")
			for _, item := range items {
				value, err := strconv.Atoi(item)
				if err != nil {
					continue
				}
				monkeys[len(monkeys)-1].items = append(monkeys[len(monkeys)-1].items, value)
			}
		}
		if strings.Index(line, "Operation") == 2 {
			cmds := strings.Split(line[strings.Index(line, "=")+2:], " ")

			var op func(int, int) int

			switch cmds[1] {
			case "+":
				op = func(a int, b int) int { return a + b }
			case "-":
				op = func(a int, b int) int { return a - b }
			case "*":
				op = func(a int, b int) int { return a * b }
			case "/":
				op = func(a int, b int) int { return a / b }
			}

			switch cmds[0] {
			case "old":
				switch cmds[2] {
				case "old":
					monkeys[len(monkeys)-1].op = func(old int) int { return op(old, old) }
				default:
					arg2, err := strconv.Atoi(cmds[2])
					if err != nil {
						continue
					}
					monkeys[len(monkeys)-1].op = func(old int) int { return op(old, arg2) }
				}
			default:
				arg1, err := strconv.Atoi(cmds[0])
				if err != nil {
					continue
				}
				switch cmds[2] {
				case "old":
					monkeys[len(monkeys)-1].op = func(old int) int { return op(arg1, old) }
				default:
					arg2, err := strconv.Atoi(cmds[2])
					if err != nil {
						continue
					}
					monkeys[len(monkeys)-1].op = func(old int) int { return op(arg1, arg2) }
				}
			}
		}
		if strings.Index(line, "Test") == 2 {
			cmds := strings.Split(line, " ")
			num, err := strconv.Atoi(cmds[len(cmds)-1])
			if err != nil {
				continue
			}
			lcm = lcm * num
			monkeys[len(monkeys)-1].test = func(count int) bool { return count%num == 0 }
		}
		if strings.Index(line, "If true") == 4 {
			cmds := strings.Split(line, " ")
			num, err := strconv.Atoi(cmds[len(cmds)-1])
			if err != nil {
				continue
			}
			monkeys[len(monkeys)-1].pass = num
		}
		if strings.Index(line, "If false") == 4 {
			cmds := strings.Split(line, " ")
			num, err := strconv.Atoi(cmds[len(cmds)-1])
			if err != nil {
				continue
			}
			monkeys[len(monkeys)-1].fail = num
		}
	}

	for i := 0; i < 10000; i++ {
		for j, monkey := range monkeys {
			monkey.count += len(monkey.items)
			for _, item := range monkey.items {
				item = monkey.op(item)
				item = item % lcm

				if monkey.test(item) {
					monkeys[monkey.pass].items = append(monkeys[monkey.pass].items, item)
				} else {
					monkeys[monkey.fail].items = append(monkeys[monkey.fail].items, item)
				}
			}
			monkey.items = []int{}
			monkeys[j] = monkey
			fmt.Printf("%+v\n", monkey)
		}
		fmt.Println()
		fmt.Println()
	}

	counts := []int{}
	for _, monkey := range monkeys {
		counts = append(counts, monkey.count)
	}
	sort.Ints(counts)

	fmt.Printf("%+v\n", counts)
	fmt.Printf("%+v\n", counts[len(counts)-1]*counts[len(counts)-2])
}
