package part2

import (
	"fmt"
	"strconv"

	"github.com/evansalter/advent-2019/helpers"
)

func calcFuel(mass int) int {
	f := (mass / 3) - 2
	if f <= 0 {
		return 0
	}
	return calcFuel(f) + f
}

func Run() {
	lines := helpers.ReadInputFile(1)
	total := 0
	for _, l := range lines {
		i, err := strconv.Atoi(l)
		if err != nil {
			fmt.Printf("Error converting %s to int: %s", l, err.Error())
			return
		}

		total += calcFuel(i)
	}
	fmt.Println(total)
}
