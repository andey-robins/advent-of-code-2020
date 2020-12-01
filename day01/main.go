package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// part 1 answer 1016619
// part 2 answer 218767230

// i figured a brute force solution here is simplest to implement and works since performance isn't really necessary at all.
// this was more of a setup to get used to file operations in go since i've never had to do them before

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	expenses := make([]int, 1)

	// open file and load it into the slice
	f, err := os.Open("./expenses.txt")
	check(err)
	defer func() {
		if err = f.Close(); err != nil {
			panic(err)
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		num, err := strconv.Atoi(s.Text())
		check(err)
		expenses = append(expenses, num)
	}
	err = s.Err()
	check(err)

	// finished with reading in the file

	partOne := false

	if partOne {
		for _, n := range expenses {
			for _, m := range expenses {
				if n+m == 2020 {
					fmt.Println(n * m)
					return
				}
			}
		}
		fmt.Println("Didn't find anything")
	} else {
		for _, n := range expenses {
			for _, m := range expenses {
				for _, o := range expenses {
					if n+m+o == 2020 && n != 0 && m != 0 && o != 0 {
						fmt.Println(n * m * o)
						return
					}
				}
			}
		}
	}
}
