package part1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/evansalter/advent-2019/helpers"
)

func stringsToInts(strs ...string) ([]int, error) {
	out := make([]int, len(strs))
	var err error
	for i, s := range strs {
		out[i], err = strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func runProgram(program []string) []string {
	for i := 0; i < len(program); i += 4 {
		op := program[i]
		if op == "99" {
			break
		}
		args, err := stringsToInts(program[i+1], program[i+2], program[i+3])
		if err != nil {
			fmt.Printf("Error converting strings to ints: %s\n", err.Error())
			return nil
		}
		out := args[2]
		inputs, err := stringsToInts(program[args[0]], program[args[1]])
		if err != nil {
			fmt.Printf("Error converting strings to ints: %s\n", err.Error())
			return nil
		}
		input1, input2 := inputs[0], inputs[1]
		var result int
		switch op {
		case "1": // Addition
			result = input1 + input2
		case "2": // Multiplication
			result = input1 * input2
		}
		program[out] = strconv.Itoa(result)
	}
	return program
}

func Run() {
	input := helpers.ReadInputFile(2)[0]
	program := strings.Split(input, ",")

	program[1] = "12"
	program[2] = "2"

	program = runProgram(program)
	fmt.Println(program[0])
}
