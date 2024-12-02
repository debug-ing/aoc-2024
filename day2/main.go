package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/debug-ing/aoc-2024/day2/part1"
	"github.com/debug-ing/aoc-2024/day2/part2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <1 or 2>")
		return
	}

	mode, err := strconv.Atoi(os.Args[1])
	if err != nil || (mode != 1 && mode != 2) {
		fmt.Println("Please provide a valid mode: 1 or 2")
		return
	}

	input, err := os.ReadFile("input")
	if err != nil {
		panic("Error reading file:" + err.Error())
	}
	var result int
	if mode == 1 {
		result = part1.Execute(string(input))
	} else if mode == 2 {
		result = part2.Execute(string(input))
	}
	fmt.Println(result)

}
