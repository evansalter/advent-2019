package part1

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/evansalter/advent-2019/helpers"
)

const (
	ImgWidth  = 25
	ImgHeight = 6
)

func breakIntoLayers(imgData string) []string {
	pixelsPerLayer := ImgWidth * ImgHeight
	layers := make([]string, 0)
	for len(imgData) > 0 {
		layers = append(layers, imgData[:pixelsPerLayer])
		imgData = imgData[pixelsPerLayer:]
	}
	return layers
}

func findNumDigitInString(digit int, str string) int {
	return strings.Count(str, strconv.Itoa(digit))
}

func Run() {
	input := helpers.ReadInputFile(8)[0]
	layers := breakIntoLayers(input)

	lowestZeroCount, lowestZeroLayer := math.MaxInt32, -1
	for i, l := range layers {
		if n := findNumDigitInString(0, l); n < lowestZeroCount {
			lowestZeroCount = n
			lowestZeroLayer = i
		}
	}

	l := layers[lowestZeroLayer]
	fmt.Println(findNumDigitInString(1, l) * findNumDigitInString(2, l))
}
