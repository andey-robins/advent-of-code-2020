// starter template
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	computer "github.com/andey-robins/advent-of-code-2020/computer"
)

// part 1 solution is 1331
// part 2 solution is not 761 (too low); solution is 1121

// part 1 solution runs in 635 milliseconds
// part 2 solution runs in 634 milliseconds

// i'm pretty sure this is grosly memory innefficient; however, it is really performant.
//

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getOpcode(command string) computer.Opr {
	if command == "jmp" {
		return computer.Jmp
	} else if command == "acc" {
		return computer.Acc
	} else if command == "nop" {
		return computer.Nop
	}

	return nil
}

func flipNopJmp(newCompiledCode []computer.Command, i int) []computer.Command {
	if newCompiledCode[i].GetOpcode() == getOpcode("jmp") {
		newCompiledCode[i] = computer.NewCommand(getOpcode("nop"), newCompiledCode[i].GetArgument())
	} else if newCompiledCode[i].GetOpcode() == getOpcode("nop") {
		newCompiledCode[i] = computer.NewCommand(getOpcode("jmp"), newCompiledCode[i].GetArgument())
	}

	return newCompiledCode
}

func main() {
	// select part one or two
	partOne := true
	fileName := "./bootcode.txt"
	if len(os.Args) == 2 && os.Args[1] == "2" {
		partOne = false
	} else if len(os.Args) == 2 && os.Args[1] == "s" {
		fileName = "./bootcodesmall.txt"
	}

	f, err := os.Open(fileName)
	check(err)
	defer func() {
		err := f.Close()
		check(err)
	}()

	s := bufio.NewScanner(f)
	// compiler
	compiledCode := make([]computer.Command, 0)
	for s.Scan() {

		if s.Text()[0] == 'j' {
			argument, err := strconv.Atoi(s.Text()[4:])
			check(err)
			compiledCode = append(compiledCode, computer.NewCommand(getOpcode("jmp"), argument))
		} else if s.Text()[0] == 'a' {
			argument, err := strconv.Atoi(s.Text()[4:])
			check(err)
			compiledCode = append(compiledCode, computer.NewCommand(getOpcode("acc"), argument))
		} else if s.Text()[0] == 'n' {
			argument, err := strconv.Atoi(s.Text()[4:])
			check(err)
			compiledCode = append(compiledCode, computer.NewCommand(getOpcode("nop"), argument))
		}
	}
	// end of compiler

	if partOne {

		bootProgram := computer.NewProgram(compiledCode)
		result, _ := bootProgram.Run()

		fmt.Printf("Part 1 solution: %v\n", result)

	} else {
		// here comes a brute force solution, oh boy

		newCompiledCode := compiledCode

		for i := 0; i < len(compiledCode); i++ {

			// flip a command
			newCompiledCode = flipNopJmp(newCompiledCode, i)

			// run the program to see if it fails
			bootProgram := computer.NewProgram(newCompiledCode)
			result, err := bootProgram.Run()
			if err == nil {
				fmt.Printf("Part 2 solutions: %v\n", result)
				break
			}

			// flip it back
			newCompiledCode = flipNopJmp(newCompiledCode, i)

		}
	}
}
