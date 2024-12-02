package part1

import (
	"bufio"
	"strconv"
	"strings"
)

func Execute(input string) (safe int) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 0 {
			if check(line) {
				safe++
			}
		}
	}
	return
}

func check(line string) (result bool) {
	parts := strings.Split(line, " ")
	status := 0
	for index, _ := range parts {
		if index == 0 {
			continue
		}
		v1, _ := strconv.Atoi(parts[index])
		v2, _ := strconv.Atoi(parts[index-1])
		if v1 > v2 {
			if status == 2 {
				return false
			}
			status = 1
		}
		if v1 < v2 {
			if status == 1 {
				return false
			}
			status = 2
		}

		if v1 == v2 {
			return false
		}
		if v1-v2 > 3 || v1-v2 < -3 {
			return false
		}
	}
	return true
}
