package part1

import (
	"fmt"
	"github.com/evansalter/advent-2019/helpers"
	"math"
	"strconv"
	"strings"
)

func intToDigits(in int) []int {
	numLen := int(math.Log10(float64(in))) + 1
	digits := make([]int, numLen)

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = in % 10
		in /= 10
	}
	return digits
}

func isValidPassword(in int) bool {
	digits := intToDigits(in)

	if len(digits) != 6 {
		return false
	}

	numDoubleDigits := 0
	lastDigit := -1
	for _, d := range digits {
		if lastDigit > d {
			return false
		}

		if lastDigit == d {
			numDoubleDigits++
		}
		lastDigit = d
	}

	if numDoubleDigits == 0 {
		return false
	}

	return true
}

func Run() {
	lines := helpers.ReadInputFile(4)
	r := strings.Split(lines[0], "-")

	rangeStart, err := strconv.Atoi(r[0])
	if err != nil {
		fmt.Printf("Error converting %s to int: %s\n", r[0], err.Error())
		return
	}
	rangeEnd, err := strconv.Atoi(r[1])
	if err != nil {
		fmt.Printf("Error converting %s to int: %s\n", r[1], err.Error())
		return
	}

	validPasswords := 0
	for i := rangeStart; i <= rangeEnd; i++ {
		if isValidPassword(i) {
			validPasswords++
		}
	}

	fmt.Println(validPasswords)
}
