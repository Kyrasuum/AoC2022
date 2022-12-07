package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	max    int = 70000000
	used   int = 0
	target int = max - 30000000
	del    int = 70000000
)

type dir struct {
	name   string
	size   int
	childs map[string]interface{}
	parent interface{}
}

type file struct {
	name string
	size int
}

func getSize(curr *dir) int {
	if curr.size != 0 {
		return curr.size
	}

	size := 0
	for _, child := range curr.childs {
		switch child.(type) {
		case *dir:
			size += getSize(child.(*dir))
		case *file:
			size += child.(*file).size
		}
	}

	curr.size = size
	return size
}

func indent(step int) string {
	ret := ""
	for i := 0; i < step; i++ {
		ret += "   "
	}
	return ret
}

func ls(curr *dir, step int) {
	fmt.Printf("%s%d %s\n", indent(step), curr.size, curr.name)
	if curr.size >= used-target && curr.size < del {
		del = curr.size
	}
	for _, child := range curr.childs {
		switch child.(type) {
		case *dir:
			ls(child.(*dir), step+1)
		case *file:
			fmt.Printf("%s%d %s\n", indent(step+1), child.(*file).size, child.(*file).name)
		}
	}
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error opening input file: %+v\n", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	root := &dir{name: "/", size: 0, childs: map[string]interface{}{}}
	curr := root
	for _, line := range lines {
		//skip empty lines
		if len(line) < 1 {
			continue
		}

		if line[0] == '$' {
			//commands
			if strings.Index(line, "cd") == 2 {
				//change current
				tgt := line[5:]
				if tgt == ".." {
					curr = curr.parent.(*dir)
				} else if tgt == "/" {
					curr = root
				} else {
					curr = curr.childs[tgt].(*dir)
				}
			}
		} else if strings.Index(line, "dir") == 0 {
			//directories
			tgt := line[4:]
			curr.childs[tgt] = &dir{name: tgt, size: 0, childs: map[string]interface{}{}, parent: curr}
		} else {
			//files
			i := strings.Index(line, " ")
			tgt := line[i+1:]
			size, _ := strconv.Atoi(line[:i])
			curr.childs[tgt] = &file{name: tgt, size: size}
		}
	}
	getSize(root)
	used = root.size
	ls(root, 0)
	fmt.Printf("%+v\n", used)
	fmt.Printf("%+v\n", used-target)
	fmt.Printf("%+v\n", del)
}
