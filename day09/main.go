// starter template
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// part 1 solution is 375054920
// part 2 solution is not 47412131 (too low); it's 54142584

// i should really start doing these when i'm not dog tired. For part two instead of adding the smallest and largest, i added the first and the last
// this one was a really interesting problem because it forced me to make a solution that was far more linear and integrated betweens part 1 and 2 than
// most of the other days. Otherwise though, a nice problem set.

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// check for sum returns true if nextNumber is a valid next number from the list numbers
func checkForSum(numbers []int, nextNumber int) bool {
	for i, numOne := range numbers {
		for j, numTwo := range numbers {
			if i != j && numOne+numTwo == nextNumber {
				return true
			}
		}
	}

	return false
}

func main() {
	// select part one or two
	filename := "./xmas.txt"
	preambleLen := 25
	if len(os.Args) == 2 && os.Args[1] == "s" {
		filename = "./xmassmall.txt"
		preambleLen = 5
	}

	f, err := os.Open(filename)
	check(err)
	defer func() {
		err := f.Close()
		check(err)
	}()

	allNumbers := make([]int, 0)
	lastNumbers := make([]int, preambleLen)
	numberIndex := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		number, err := strconv.Atoi(s.Text())
		check(err)

		if numberIndex < preambleLen {

			// load in preamble
			lastNumbers[numberIndex%preambleLen] = number
			allNumbers = append(allNumbers, number)

		} else if checkForSum(lastNumbers, number) {

			// the next number is a valid next number
			lastNumbers[numberIndex%preambleLen] = number
			allNumbers = append(allNumbers, number)

		} else {

			// not a valid next number
			badNumber := number
			fmt.Printf("Part 1 solution: %v\n", number)

			for i, firstNumber := range allNumbers {

				sum := 0
				smallest := firstNumber
				largest := firstNumber

				for j := i; j < len(allNumbers); j++ {
					sum += allNumbers[j]

					if allNumbers[j] > largest {
						largest = allNumbers[j]
					}

					if allNumbers[j] < smallest {
						smallest = allNumbers[j]
					}

					if sum == badNumber {
						fmt.Printf("Part 2 solution: %v + %v = %v\n", smallest, largest, smallest+largest)
						return
					}
				}
			}
		}

		numberIndex++
	}

	// if partOne {
	// 	fmt.Println("Part 1")
	// } else {
	// 	fmt.Println("Part 2")
	// }
}
