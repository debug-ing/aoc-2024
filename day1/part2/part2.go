package part2

import (
	"bufio"
	"strconv"
	"strings"
)

type Data struct {
	Repeat int
	InLeft bool
}

func Execute(input string) int {
	list := make(map[int]*Data)
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
			if v, exists := list[pv1]; exists {
				v.InLeft = true
			} else {
				list[pv1] = &Data{Repeat: 0, InLeft: true}
			}
			if v, exists := list[pv2]; exists {
				v.Repeat++
			} else {
				list[pv2] = &Data{Repeat: 1, InLeft: false}
			}
		}
	}
	sum := 0
	for k, v := range list {
		if v.InLeft {
			sum += k * v.Repeat
		}
	}

	return sum
}
