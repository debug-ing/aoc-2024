package part2

import (
	"math"
	"strconv"
	"strings"
)

func Execute(input string) (safe int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		numberStrings := strings.Fields(line)
		lineNumbers := make([]uint32, len(numberStrings))
		for i, n := range numberStrings {
			num, err := strconv.ParseUint(n, 10, 32)
			if err != nil {
				panic(err)
			}
			lineNumbers[i] = uint32(num)
		}

		for i := 0; i < len(lineNumbers); i++ {
			tempLineNumbers := append([]uint32{}, lineNumbers[:i]...)
			tempLineNumbers = append(tempLineNumbers, lineNumbers[i+1:]...)

			if len(tempLineNumbers) < 2 {
				continue
			}

			first := tempLineNumbers[0]
			second := tempLineNumbers[1]
			increasing := first < second

			diff := uint32(math.Abs(float64(int(second) - int(first))))
			if diff < 1 || diff > 3 {
				continue
			}

			valid := true
			for j := 1; j < len(tempLineNumbers); j++ {
				first := tempLineNumbers[j-1]
				second := tempLineNumbers[j]

				if increasing && first >= second {
					valid = false
					break
				} else if !increasing && first <= second {
					valid = false
					break
				}

				diff := uint32(math.Abs(float64(int(second) - int(first))))
				if diff < 1 || diff > 3 {
					valid = false
					break
				}
			}

			if valid {
				safe += 1
				break
			}
		}
	}

	return
}
