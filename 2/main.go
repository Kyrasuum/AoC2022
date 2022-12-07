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

	scores := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}
	strategy := map[string]map[string]string{
		"X": {
			"A": "Z",
			"B": "X",
			"C": "Y",
		},
		"Y": {
			"A": "X",
			"B": "Y",
			"C": "Z",
		},
		"Z": {
			"A": "Y",
			"B": "Z",
			"C": "X",
		},
	}
	total := 0

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		chars := strings.Split(line, " ")
		if len(chars) < 2 {
			continue
		}
		theirs := chars[0]
		result := chars[1]
		ours := strategy[result][theirs]
		if score, ok := scores[theirs+" "+ours]; ok {
			total += score
		}
	}
	fmt.Printf("%+v\n", total)
}
