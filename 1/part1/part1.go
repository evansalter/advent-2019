package part1

import (
	"fmt"
	"strconv"

	"github.com/evansalter/advent-2019/helpers"
)

func Run() {
	lines := helpers.ReadInputFile(1)
	total := 0
	for _, l := range lines {
		i, err := strconv.Atoi(l)
		if err != nil {
			fmt.Printf("Error converting %s to int: %s", l, err.Error())
			return
		}

		f := (i / 3) - 2
		total += f
	}
	fmt.Println(total)
}
