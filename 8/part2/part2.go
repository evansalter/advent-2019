package part2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/evansalter/advent-2019/helpers"
)

const (
	ImgWidth  = 25
	ImgHeight = 6

	PixelBlack       = "0"
	PixelWhite       = "1"
	PixelTransparent = "2"
)

type Pixel struct {
	values []string
}

func (p *Pixel) ComputePixel() string {
	for _, v := range p.values {
		if v == PixelTransparent {
			continue
		}
		return v
	}
	return PixelTransparent
}

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

func printImage(img []string, width int) {
	for len(img) > 0 {
		fmt.Println(strings.Join(img[:width], ""))
		img = img[width:]
	}
}

func Run() {
	input := helpers.ReadInputFile(8)[0]
	layers := breakIntoLayers(input)
	imgSize := ImgWidth * ImgHeight

	pixels := make([]*Pixel, imgSize)
	for i := range pixels {
		vals := make([]string, imgSize)
		for j, l := range layers {
			vals[j] = string([]rune(l)[i])
		}
		pixels[i] = &Pixel{values: vals}
	}

	img := make([]string, imgSize)
	for i, p := range pixels {
		img[i] = p.ComputePixel()
	}

	printImage(img, ImgWidth)
}
