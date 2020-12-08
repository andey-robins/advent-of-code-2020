package main

import (
	"bufio"
	"fmt"
	"os"
)

type Bag struct {
	adj   string
	color string
}

type BagTreeNode struct {
	bag      Bag
	children []Bag
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// select part one or two
	partOne := true
	if len(os.Args) == 2 && os.Args[1] == "2" {
		partOne = false
	}

	f, err := os.Open("./rules.txt")
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
