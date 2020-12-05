// starter template
package main

import (
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
	if len(os.Args) == 2 && os.Args[1] == "2" {
		partOne = false
	}

	if partOne {
		fmt.Println("Part 1")
	} else {
		fmt.Println("Part 2")
	}
}
