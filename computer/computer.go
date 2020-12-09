package computer

import "fmt"

//
// ERRORS
//

// LoopError is returned if a loop in the code is detected
type LoopError struct {
	prob string
}

func (e *LoopError) Error() string {
	return fmt.Sprintf("%s", e.prob)
}

// Op is an enum for tracking what command operation is executed
type Op int

// enum for opcodes
const (
	Acc Op = iota
	Jmp
	Nop
)

// Command is a structure that combines an opcode and an argument to represent a line of code
type Command struct {
	opcode   Opr
	argument int
}

// NewCommand generates a new command object from a given opcode and argument
func NewCommand(operation Opr, argument int) Command {
	return Command{opcode: operation, argument: argument}
}

// GetOpcode returns the Opr of a command object
func (c *Command) GetOpcode() Opr {
	return c.opcode
}

// GetArgument returns the argument associated with a Command object
func (c *Command) GetArgument() int {
	return c.argument
}

// Program is a struct that contains an entire program. load the program into here and "run" it
type Program struct {
	code        []Command
	accumulator int
}

// NewProgram return a New Program object
func NewProgram(code []Command) Program {
	return Program{code: code, accumulator: 0}
}

// Opr is an interface to opcodes
type Opr interface {
	Op() Op
}

// Op fulfills the Opr interface
func (o Op) Op() Op {
	return o
}

// Run runs a program and returns data based on the output of running a program
func (p Program) Run() (int, error) {
	loopDetection := make([]bool, len(p.code))

	for programCounter := 0; programCounter < len(p.code); programCounter++ {

		if !loopDetection[programCounter] {
			loopDetection[programCounter] = true
		} else {
			return p.accumulator, &LoopError{"Loop Detected"}
		}

		if p.code[programCounter].opcode == Acc {
			p.accumulator += p.code[programCounter].argument
		} else if p.code[programCounter].opcode == Jmp {
			programCounter += p.code[programCounter].argument
			programCounter--
		}

		// nop doesn't do anything cause... y'know... nop
	}

	return p.accumulator, nil
}
