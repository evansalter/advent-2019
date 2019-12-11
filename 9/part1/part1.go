package part1

import (
	"strings"

	"github.com/evansalter/advent-2019/9/intcode"
	"github.com/evansalter/advent-2019/helpers"
)

func Run() {
	lines := helpers.ReadInputFile(9)[0]
	input := strings.Split(lines, ",")

	p := intcode.NewProgram(input)
	p.Run()
}
