package intcode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	// MaxArguments is the max number of arguments an operation can have
	MaxArguments = 3
)

// OpCode represents the action to be performed
type OpCode string

const (
	Add         OpCode = "01"
	Multiply           = "02"
	Input              = "03"
	Output             = "04"
	JumpIfTrue         = "05"
	JumpIfFalse        = "06"
	LessThan           = "07"
	Equals             = "08"
	Halt               = "99"
)

// Mode represents the method in which an argument value is retrieved
type Mode string

const (
	// Position means the argument value is a program index
	Position Mode = "0"
	// Immediate means the argument value should be used as is
	Immediate = "1"
	// Relative means the argument value is a memory address relative to the relative base
	Relative = "2"
)

// Program contains the instruction counter, as well as the opcode program
type Program struct {
	Counter int
	Program []string
	running bool
}

// NewProgram initializes a program with the given opcode slice
func NewProgram(p []string) *Program {
	return &Program{
		Counter: 0,
		Program: p,
		running: true,
	}
}

// Run runs the program to completion
func (p *Program) Run() {
	for p.HasMore() {
		cmd := p.GetNextCommand()
		cmd.Execute()
		cmd.IncrementCounter()
	}
}

// Halt stops the running program
func (p *Program) Halt() {
	p.running = false
}

// HasMore returns true if there are more instructions to execute, false otherwise
func (p *Program) HasMore() bool {
	return p.Counter < len(p.Program)-1 && p.running
}

// IncrCounter increments the program counter by n
func (p *Program) IncrCounter(n int) {
	p.Counter += n
}

// SetStringValue sets the value at the given position
func (p *Program) SetStringValue(pos int, val string) {
	p.Program[pos] = val
}

// SetIntValue sets the value at the given position, after converting it to a string
func (p *Program) SetIntValue(pos, val int) {
	p.SetStringValue(pos, strconv.Itoa(val))
}

// GetNextCommand returns the next command to be executed
func (p *Program) GetNextCommand() Command {
	opcode := p.Program[p.Counter]

	// Left-pad with zeros
	for len(opcode) < 5 {
		opcode = fmt.Sprintf("0%s", opcode)
	}
	op := OpCode(opcode[3:])
	if op == Halt {
		return NewHaltCommand(p)
	}

	switch op {
	case Add:
		return NewAddCommand(p)
	case Multiply:
		return NewMultiplyCommand(p)
	case Input:
		return NewInputCommand(p)
	case Output:
		return NewOutputCommand(p)
	case JumpIfTrue:
		return NewJumpIfTrueCommand(p)
	case JumpIfFalse:
		return NewJumpIfFalseCommand(p)
	case LessThan:
		return NewLessThanCommand(p)
	case Equals:
		return NewEqualsCommand(p)
	default:
		panic(fmt.Sprintf("Unexpected command: %s", op))
	}
}

// GetNArguments returns the next n arguments in the program
func (p *Program) GetNArguments(n int) []*Argument {
	if n > MaxArguments {
		panic(fmt.Sprintf("Cannot request more than %d arguments, requested %d", MaxArguments, n))
	}
	opcode := p.Program[p.Counter]

	// Left-pad with zeros
	for len(opcode) < 5 {
		opcode = fmt.Sprintf("0%s", opcode)
	}
	modeStr := opcode[:3]
	mode1, mode2 := modeStr[2:], modeStr[1:2]
	modes := []Mode{Mode(mode1), Mode(mode2), Position}

	args := make([]*Argument, n)
	for i := 0; i < len(args); i++ {
		idx := p.Counter + i + 1
		val, err := strconv.Atoi(p.Program[idx])
		if err != nil {
			panic(fmt.Sprintf("Error converting %s to int: %s", p.Program[idx], err.Error()))
		}
		args[i] = NewArgument(modes[i], val)
	}

	return args
}

// Argument is a parameter to a command, containing a mode and a value
type Argument struct {
	Mode Mode
	Val  int
}

// NewArgument returns a new argument with the given mode and value
func NewArgument(mode Mode, val int) *Argument {
	return &Argument{
		Mode: mode,
		Val:  val,
	}
}

// GetValue looks up the argument in the program, using the correct mode
func (a *Argument) GetValue(p *Program) int {
	var str string
	switch a.Mode {
	case Position:
		if a.Val >= len(p.Program) {
			fmt.Printf("Value %d is greater then length of program %d\n", a.Val, len(p.Program))
			return 0
		}
		str = p.Program[a.Val]
	case Immediate:
		return a.Val
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

// Command is a unit of work supported by the opcode machine
type Command interface {
	Execute()
	IncrementCounter()
}

// AddCommand adds arg1 and arg2 together, placing the result in output
type AddCommand struct {
	program *Program
	arg1    *Argument
	arg2    *Argument
	output  *Argument
}

// NewAddCommand returns an AddCommand instance
func NewAddCommand(p *Program) *AddCommand {
	args := p.GetNArguments(3)
	return &AddCommand{
		program: p,
		arg1:    args[0],
		arg2:    args[1],
		output:  args[2],
	}
}

// Execute runs the command
func (c *AddCommand) Execute() {
	c.program.SetIntValue(c.output.Val, c.arg1.GetValue(c.program)+c.arg2.GetValue(c.program))
}

// IncrementCounter moves the program counter the correct number of places
func (c *AddCommand) IncrementCounter() {
	c.program.IncrCounter(4)
}

// MultiplyCommand multiplies arg1 and arg2, placing the result in output
type MultiplyCommand struct {
	program *Program
	arg1    *Argument
	arg2    *Argument
	output  *Argument
}

// NewMultiplyCommand returns a new MultiplyCommand instance
func NewMultiplyCommand(p *Program) *MultiplyCommand {
	args := p.GetNArguments(3)
	return &MultiplyCommand{
		program: p,
		arg1:    args[0],
		arg2:    args[1],
		output:  args[2],
	}
}

// Execute runs the command
func (c *MultiplyCommand) Execute() {
	c.program.SetIntValue(c.output.Val, c.arg1.GetValue(c.program)*c.arg2.GetValue(c.program))
}

// IncrementCounter moves the program counter the correct number of places
func (c *MultiplyCommand) IncrementCounter() {
	c.program.IncrCounter(4)
}

// InputCommand asks the user for a character input and places it in the output
type InputCommand struct {
	program *Program
	output  *Argument
}

// NewInputCommand returns an instance of InputCommand
func NewInputCommand(p *Program) *InputCommand {
	args := p.GetNArguments(1)
	return &InputCommand{
		program: p,
		output:  args[0],
	}
}

// Execute runs the command
func (c *InputCommand) Execute() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input: ")
	char, _, err := reader.ReadRune()
	if err != nil {
		panic(fmt.Sprintf("Error reading character: %s", err.Error()))
	}
	c.program.SetStringValue(c.output.Val, string(char))
}

// IncrementCounter moves the program counter the correct number of places
func (c *InputCommand) IncrementCounter() {
	c.program.IncrCounter(2)
}

// OutputCommand prints the given field to the console
type OutputCommand struct {
	program *Program
	field   *Argument
}

// NewOutputCommand returns a new instance of OutputCommand
func NewOutputCommand(p *Program) *OutputCommand {
	args := p.GetNArguments(1)
	return &OutputCommand{
		program: p,
		field:   args[0],
	}
}

// Execute runs the command
func (c *OutputCommand) Execute() {
	fmt.Println(c.program.Program[c.field.Val])
}

// IncrementCounter moves the program counter the correct number of places
func (c *OutputCommand) IncrementCounter() {
	c.program.IncrCounter(2)
}

// JumpIfTrueCommand jumps to the target instructor if the argument is not zero
type JumpIfTrueCommand struct {
	program    *Program
	arg1       *Argument
	target     *Argument
	shouldJump bool
}

// NewJumpIfTrueCommand returns a new instance of JumpIfTrueCommand
func NewJumpIfTrueCommand(p *Program) *JumpIfTrueCommand {
	args := p.GetNArguments(2)
	return &JumpIfTrueCommand{
		program: p,
		arg1:    args[0],
		target:  args[1],
	}
}

// Execute runs the command
func (c *JumpIfTrueCommand) Execute() {
	if c.arg1.GetValue(c.program) != 0 {
		c.shouldJump = true
	}
}

// IncrementCounter moves the program counter the correct number of places
func (c *JumpIfTrueCommand) IncrementCounter() {
	if c.shouldJump {
		c.program.Counter = c.target.GetValue(c.program)
	} else {
		c.program.IncrCounter(3)
	}
}

// JumpIfFalseCommand jumps to the target instructor if the argument is zero
type JumpIfFalseCommand struct {
	program    *Program
	arg1       *Argument
	target     *Argument
	shouldJump bool
}

// NewJumpIfFalseCommand returns a new instance of JumpIfFalseCommand
func NewJumpIfFalseCommand(p *Program) *JumpIfFalseCommand {
	args := p.GetNArguments(2)
	return &JumpIfFalseCommand{
		program: p,
		arg1:    args[0],
		target:  args[1],
	}
}

// Execute runs the command
func (c *JumpIfFalseCommand) Execute() {
	if c.arg1.GetValue(c.program) == 0 {
		c.shouldJump = true
	}
}

// IncrementCounter moves the program counter the correct number of places
func (c *JumpIfFalseCommand) IncrementCounter() {
	if c.shouldJump {
		c.program.Counter = c.target.GetValue(c.program)
	} else {
		c.program.IncrCounter(3)
	}
}

// LessThanCommand puts 1 in the output if arg1 < arg2, 0 otherwise
type LessThanCommand struct {
	program *Program
	arg1    *Argument
	arg2    *Argument
	output  *Argument
}

// NewLessThanCommand returns a new instance of LessThanCommand
func NewLessThanCommand(p *Program) *LessThanCommand {
	args := p.GetNArguments(3)
	return &LessThanCommand{
		program: p,
		arg1:    args[0],
		arg2:    args[1],
		output:  args[2],
	}
}

// Execute runs the command
func (c *LessThanCommand) Execute() {
	out := 0
	if c.arg1.GetValue(c.program) < c.arg2.GetValue(c.program) {
		out = 1
	}
	c.program.SetIntValue(c.output.Val, out)
}

// IncrementCounter moves the program counter the correct number of places
func (c *LessThanCommand) IncrementCounter() {
	c.program.IncrCounter(4)
}

// EqualsCommand puts 1 in the output if arg1 == arg2, 0 otherwise
type EqualsCommand struct {
	program *Program
	arg1    *Argument
	arg2    *Argument
	output  *Argument
}

// NewEqualsCommand returns a new instance of EqualsCommand
func NewEqualsCommand(p *Program) *EqualsCommand {
	args := p.GetNArguments(3)
	return &EqualsCommand{
		program: p,
		arg1:    args[0],
		arg2:    args[1],
		output:  args[2],
	}
}

// Execute runs the command
func (c *EqualsCommand) Execute() {
	out := 0
	if c.arg1.GetValue(c.program) == c.arg2.GetValue(c.program) {
		out = 1
	}
	c.program.SetIntValue(c.output.Val, out)
}

// IncrementCounter moves the program counter the correct number of places
func (c *EqualsCommand) IncrementCounter() {
	c.program.IncrCounter(4)
}

// HaltCommand halts the running program
type HaltCommand struct {
	program *Program
}

// NewHaltCommand returns a new instance of HaltCommand
func NewHaltCommand(p *Program) *HaltCommand {
	return &HaltCommand{
		program: p,
	}
}

// Execute runs the command
func (c *HaltCommand) Execute() {
	c.program.Halt()
}

// IncrementCounter moves the program counter the correct number of places
func (c *HaltCommand) IncrementCounter() {
	return
}
