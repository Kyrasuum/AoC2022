package main

import (
	"fmt"
	"io/ioutil"
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
	header := ""
	for _, line := range lines {
		for i, char := range line {
			if strings.Contains(header, string(char)) {
				header = header[strings.Index(header, string(char))+1:]
			}
			header = header + string(char)
			if len(header) == 14 {
				fmt.Printf("%+v\n", i+1)
				fmt.Printf("%+v\n", header)
				break
			}
		}
	}
}
