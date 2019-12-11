package part1

import (
	"fmt"
	"strings"

	"github.com/evansalter/advent-2019/helpers"
)

func intSliceHasDupes(s []int) bool {
	m := make(map[int]int)
	for _, i := range s {
		m[i]++
	}

	for _, i := range m {
		if i > 1 {
			return true
		}
	}

	return false
}

func findMaxInt(s []int) int {
	highest := 0
	for _, i := range s {
		if i > highest {
			highest = i
		}
	}
	return highest
}

func Run() {
	lines := helpers.ReadInputFile(7)[0]
	input := strings.Split(lines, ",")

	perms := make([][]int, 0)
	for i := 0; i <= 4; i++ {
		for j := 0; j <= 4; j++ {
			for y := 0; y <= 4; y++ {
				for x := 0; x <= 4; x++ {
					for z := 0; z <= 4; z++ {
						perm := []int{i, j, y, x, z}
						if !intSliceHasDupes(perm) {
							perms = append(perms, perm)
						}
					}
				}
			}
		}
	}

	outputs := make([]int, 0)
	for _, phases := range perms {
		signal := 0
		inChan, outChan := make(chan int), make(chan int)

		for _, phase := range phases {
			p := NewProgram(input, inChan, outChan)
			go p.Run()

			inChan <- phase
			inChan <- signal
			signal = <-outChan
		}
		outputs = append(outputs, signal)
	}

	fmt.Println(findMaxInt(outputs))
}
