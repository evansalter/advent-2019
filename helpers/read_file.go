package helpers

import (
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

// ReadInputFile reads the input.txt for a particular day and returns the contents split by newline
func ReadInputFile(day int) []string {
	curDir, _ := os.Open(".")
	files, _ := curDir.Readdirnames(0)

	wd, _ := os.Getwd()
	p := []string{wd}
	for _, f := range files {
		if f == "go.mod" {
			p = append(p, strconv.Itoa(day))
		}
	}
	p = append(p, "input.txt")
	fp := path.Join(p...)
	data, _ := ioutil.ReadFile(fp)
	return strings.Split(string(data), "\n")
}
