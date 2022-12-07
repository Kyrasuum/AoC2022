package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var ()

func weight(char rune) int {
	if char >= 'a' && char <= 'z' {
		return int(char) - int('a') + 1
	} else {
		return int(char) - int('A') + 27
	}
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error opening input file: %+v\n", err)
		return
	}

	sum := 0
	lines := strings.Split(string(data), "\n")
	for i := 0; i < len(lines); i += 3 {
		for _, char := range lines[i] {
			if strings.ContainsRune(lines[i+1], char) && strings.ContainsRune(lines[i+2], char) {
				sum += weight(char)
				break
			}
		}
	}
	fmt.Printf("%+v\n", sum)
}
