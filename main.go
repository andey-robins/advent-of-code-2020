// starter template
package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// select part one or two
	partOne := true
	filename := "./filename"
	if len(os.Args) == 2 && os.Args[1] == "2" {
		partOne = false
	} else if len(os.Args) == 2 && os.Args[1] == "s" {
		filename = "./filenamesmall"
	}

	f, err := os.Open(filename)
	check(err)
	defer func() {
		err := f.Close()
		check(err)
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {

	}

	if partOne {
		fmt.Println("Part 1")
	} else {
		fmt.Println("Part 2")
	}
}
