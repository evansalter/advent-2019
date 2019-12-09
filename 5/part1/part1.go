package part1

import (
	"bufio"
	"fmt"
	"os"
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

func getInput(program []string, mode string, arg int) int {
	var str string
	switch mode {
	case "0":
		if arg >= len(program) {
			return 0
		}
		str = program[arg]
	case "1":
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

func runProgram(program []string) []string {
	i := 0
	for i < len(program) {
		opcode := program[i]

		// Left-pad with zeros
		for len(opcode) < 5 {
			opcode = fmt.Sprintf("0%s", opcode)
		}
		modes, op := opcode[:3], opcode[3:]
		if op == "99" {
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
		case "01": // Addition
			program[out] = strconv.Itoa(input1 + input2)
		case "02": // Multiplication
			program[out] = strconv.Itoa(input1 * input2)
		case "03": // Save input
			char := inputChar()
			program[input1] = string(char)
		case "04": // Output
			fmt.Println(input1)
		}

		numToIncrement := 4
		if op == "03" || op == "04" {
			numToIncrement = 2
		}
		i += numToIncrement
	}
	return program
}

func Run() {
	input := helpers.ReadInputFile(5)[0]
	program := strings.Split(input, ",")
	// fmt.Println(runProgram(program))
	runProgram(program)

	// for i := 0; i <= 99; i++ {
	// 	for j := 0; j <= 99; j++ {
	// 		program := strings.Split(input, ",")

	// 		program[1] = strconv.Itoa(i)
	// 		program[2] = strconv.Itoa(j)

	// 		program = runProgram(program)
	// 		if program[0] == "19690720" {
	// 			fmt.Println(100*i + j)
	// 			return
	// 		}
	// 	}
	// }
	// fmt.Println("No answer found")
}
