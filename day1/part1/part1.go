package part1

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

func zip(list1, list2 []int) [][2]int {
	result := make([][2]int, len(list1))
	for i := range list1 {
		result[i] = [2]int{list1[i], list2[i]}
	}
	return result
}

func Execute(input string) int {
	list1 := []int{}
	list2 := []int{}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 0 {
			parts := strings.SplitN(line, "   ", 2)
			if len(parts) != 2 {
				continue
			}
			pv1, _ := strconv.Atoi(parts[0])
			pv2, _ := strconv.Atoi(parts[1])
			list1 = append(list1, pv1)
			list2 = append(list2, pv2)
		}
	}

	sort.Ints(list1)
	sort.Ints(list2)

	totalDiff := 0
	for _, pair := range zip(list1, list2) {
		x, y := pair[0], pair[1]
		if x >= y {
			totalDiff += x - y
		} else {
			totalDiff += y - x
		}
	}

	return totalDiff
}
