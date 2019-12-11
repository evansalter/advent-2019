package part2

import (
	"fmt"
	"strings"
	"sync"

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

func copyStringSlice(s []string) []string {
	new := make([]string, len(s))
	for i := range s {
		new[i] = s[i]
	}
	return new
}

func Run() {
	lines := helpers.ReadInputFile(7)[0]
	input := strings.Split(lines, ",")

	perms := make([][]int, 0)
	for i := 5; i <= 9; i++ {
		for j := 5; j <= 9; j++ {
			for y := 5; y <= 9; y++ {
				for x := 5; x <= 9; x++ {
					for z := 5; z <= 9; z++ {
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
		inChans := make([]chan int, 5)
		outChans := make([]chan int, 5)
		for i := range inChans {
			inChans[i] = make(chan int, 5)
			outChans[i] = make(chan int, 5)
		}

		wg := &sync.WaitGroup{}
		for i := range phases {
			wg.Add(1)
			go func(i int) {
				p := NewProgram(copyStringSlice(input), inChans[i], outChans[i])
				p.Run()
				wg.Done()
			}(i)
		}
		for i, phase := range phases {
			inChans[i] <- phase
			go func(i int) {
				prev := i - 1
				if prev < 0 {
					prev = 4
				}
				for v := range outChans[prev] {
					inChans[i] <- v
					if prev == 4 {
						outputs = append(outputs, v)
					}
				}
			}(i)
		}
		inChans[0] <- 0
		wg.Wait()
	}

	fmt.Println(findMaxInt(outputs))
}
