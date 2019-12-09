package part1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/evansalter/advent-2019/helpers"
)

const (
	wire         = "w"
	intersection = "X"
)

type point struct {
	X int
	Y int
}

func (p point) Manhattan() int {
	x, y := p.X, p.Y
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}

func (p point) ToPointDistance() pointDistance {
	return pointDistance{
		P:    p,
		Dist: p.Manhattan(),
	}
}

type pointDistance struct {
	P    point
	Dist int
}

func computePoint(curX, curY int, dir string) (x, y int) {
	x, y = curX, curY
	switch dir {
	case "U":
		y++
	case "R":
		x++
	case "D":
		y--
	case "L":
		x--
	}
	return x, y
}

func allIntsInSliceNonZero(slice []int) bool {
	for _, x := range slice {
		if x <= 0 {
			return false
		}
	}
	return true
}

func Run() {
	lines := helpers.ReadInputFile(3)

	wirePoints := make([][]point, len(lines))

	for i, l := range lines {
		commands := strings.Split(l, ",")
		wirePoints[i] = make([]point, 0)

		x, y := 0, 0
		for _, c := range commands {
			dir := c[:1]
			mag, err := strconv.Atoi(c[1:])
			if err != nil {
				fmt.Printf("Error converting %s to int: %s\n", c[1:], err.Error())
				return
			}

			for m := mag; m > 0; m-- {
				x, y = computePoint(x, y, dir)
				wirePoints[i] = append(wirePoints[i], point{X: x, Y: y})
			}
		}
	}

	pointCountMap := make(map[point][]int)
	for i, w := range wirePoints {
		for _, p := range w {
			if _, ok := pointCountMap[p]; !ok {
				pointCountMap[p] = make([]int, len(wirePoints))
			}
			pointCountMap[p][i]++
		}
	}

	pointDistances := make([]pointDistance, 0)
	for p, c := range pointCountMap {
		if allIntsInSliceNonZero(c) {
			pointDistances = append(pointDistances, p.ToPointDistance())
		}
	}

	sort.Slice(pointDistances, func(i, j int) bool {
		return pointDistances[i].Dist < pointDistances[j].Dist
	})

	fmt.Println(pointDistances[0].Dist)
}
