package part2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/evansalter/advent-2019/helpers"
)

const (
	OpAdd      = "01"
	OpMultiply = "02"
	OpInput    = "03"
	OpOutput   = "04"
	OpHalt     = "99"

	ModePosition  = "0"
	ModeImmediate = "1"
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

func getInput(program []string, mode string, arg int) int {
	var str string
	switch mode {
	case ModePosition:
		if arg >= len(program) {
			return 0
		}
		str = program[arg]
	case ModeImmediate:
		return arg
	}

	if str == "" {
		return 0
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("Error converting %s to int: %s", str, err.Error()))
	}

	return i
}

func inputChar() rune {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	char, _, err := reader.ReadRune()
	if err != nil {
		panic(fmt.Sprintf("Error reading character: %s", err.Error()))
	}
	return char
}

func getNumToIncrement(op string) int {
	switch op {
	case OpAdd, OpMultiply:
		return 4
	case OpInput, OpOutput:
		fallthrough
	default:
		return 2
	}
}

func runProgram(program []string) []string {
	i := 0
	for i < len(program) {
		opcode := program[i]

		// Left-pad with zeros
		for len(opcode) < 5 {
			opcode = fmt.Sprintf("0%s", opcode)
		}
		modes, op := opcode[:3], opcode[3:]
		if op == OpHalt {
			break
		}

		param1Mode, param2Mode := modes[2:], modes[1:2]

		args, err := stringsToInts(program[i+1], program[i+2], program[i+3])
		if err != nil {
			fmt.Printf("Error converting strings to ints: %s\n", err.Error())
			return nil
		}
		out := args[2]

		input1, input2 := getInput(program, param1Mode, args[0]), getInput(program, param2Mode, args[1])
		switch op {
		case OpAdd:
			program[out] = strconv.Itoa(input1 + input2)
		case OpMultiply:
			program[out] = strconv.Itoa(input1 * input2)
		case OpInput:
			char := inputChar()
			program[input1] = string(char)
		case OpOutput:
			fmt.Println(input1)
		}

		numToIncrement := getNumToIncrement(op)
		i += numToIncrement
	}
	return program
}

func Run() {
	input := helpers.ReadInputFile(5)[0]
	program := strings.Split(input, ",")
	runProgram(program)
}
