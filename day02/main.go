package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// part 1 answer 469
// part 2 answer 267

// this one went super smoothly. it's been nice getting a lot of experience with parsing in go. this solution felt even more straight forward than
// yesterdays, but that might be because i'm feeling warmed up and more in the mood for writing code. also, i added in a command line arg for switching
// which part's code is run based on that, which is a nice quality of life feature. i could document it and validate input, but i don't really feel
// like doing that because this is just for fun and i'm the only one that will ever use it (and frankly i doubt i'll ever run this code again after today)

// i really like this little helper function that i found in a guide
// so i think it's likely it'll be here every day
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// select part one or two
	partOne := true
	if os.Args[1] == "2" {
		partOne = false
	}

	// file operations
	f, err := os.Open("./passwords.txt")
	check(err)
	defer func() {
		err := f.Close()
		check(err)
	}()

	validPasswords := 0

	s := bufio.NewScanner(f)
	for s.Scan() {

		if partOne {
			splitLine := strings.Split(s.Text(), " ")
			numbers := strings.Split(splitLine[0], "-")
			character := splitLine[1][0]
			password := splitLine[2]

			count := 0
			for i, _ := range password {
				if password[i] == character {
					count++
				}
			}

			lo, err := strconv.Atoi(numbers[0])
			check(err)
			hi, err := strconv.Atoi(numbers[1])
			check(err)

			if count >= lo && count <= hi {
				validPasswords++
			}
		} else {
			splitLine := strings.Split(s.Text(), " ")
			numbers := strings.Split(splitLine[0], "-")
			character := splitLine[1][0]
			password := splitLine[2]
			lo, err := strconv.Atoi(numbers[0])
			check(err)
			hi, err := strconv.Atoi(numbers[1])
			check(err)

			first := false
			second := false

			if password[lo-1] == character {
				first = true
			}
			if password[hi-1] == character {
				second = true
			}

			if first != second {
				validPasswords++
			}
		}

	}

	fmt.Println(validPasswords)
}
