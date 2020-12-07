// starter template
package main

import (
	"bufio"
	"fmt"
	"os"
)

// part 1 answer is 6911
// part 2 ansewr is not 5699 (too high); answer is 3473

// part 1 runtime: 632 milliseconds
// part 2 runtime: 753 milliseconds

// in the conversion to part 2, i made the dumb mistake of forgetting to sum all the people as i read them in. adding in a counter for
// people immediately spit out the correct answer. runtime complexity is linear since we go through the list to read it in and then
// go through the number of entries put in twice. I also decided to time a measure runtime for the first itme and it was really cool to see
// 2000 lines of text read, processed, computed, and reported in less than a second

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type CustomsForm struct {
	form   []int
	people int
}

func main() {
	// select part one or two
	partOne := true
	if len(os.Args) == 2 && os.Args[1] == "2" {
		partOne = false
	}

	f, err := os.Open("./answers.txt")
	check(err)
	defer func() {
		err := f.Close()
		check(err)
	}()

	customsForms := make([]CustomsForm, 0)
	activeForm := CustomsForm{}
	activeForm.form = make([]int, 26)

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()

		if line == "" {
			customsForms = append(customsForms, activeForm)
			activeForm = CustomsForm{}
			activeForm.form = make([]int, 26)
		} else {
			activeForm.people++
			for _, char := range line {
				activeForm.form[char-97]++
			}
		}
	}
	customsForms = append(customsForms, activeForm)

	if partOne {

		yesses := 0

		for _, form := range customsForms {
			for _, ans := range form.form {
				if ans > 0 {
					yesses++
				}
			}
		}

		fmt.Printf("Part 1 answer: %v", yesses)

	} else {

		yesses := 0

		for _, form := range customsForms {
			for _, ans := range form.form {
				if ans == form.people {
					yesses++
				}
			}
		}

		fmt.Printf("Part 2 answer: %v", yesses)
	}
}
