// starter template
package main

import (
	"bufio"
	"fmt"
	"os"
)

// part 1 answer 247
// part 2 answer 2983070376

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

	// file operations
	f, err := os.Open("./map.txt")
	check(err)
	defer func() {
		err := f.Close()
		check(err)
	}()

	hitTrees := 0
	s := bufio.NewScanner(f)
	treeMap := [][]rune{}

	for s.Scan() {
		row := []rune(s.Text())
		treeMap = append(treeMap, row)
	}

	if partOne {
		x := 0
		y := 0

		for y < 323 {
			if treeMap[y][x%31] == '#' {
				hitTrees++
			}

			x += 3
			y++
		}

		fmt.Printf("Hit %v trees", hitTrees)
	} else {

		totalHitTrees := 1

		slope := [5][5]int{
			{1, 3, 5, 7, 1},
			{1, 1, 1, 1, 2},
		}

		attempt := 0

		for attempt < 5 {
			fmt.Printf("Attempting run with slope %v, %v\n", slope[0][attempt], slope[1][attempt])
			x := 0
			y := 0
			hitTrees = 0

			for y < 323 {
				if treeMap[y][x%31] == '#' {
					hitTrees++
				}

				x += slope[0][attempt]
				y += slope[1][attempt]
			}

			attempt++
			totalHitTrees *= hitTrees
			fmt.Printf("Hit %v trees with the tested slope\n", hitTrees)
		}

		fmt.Printf("Part 2 answer: %v", totalHitTrees)
	}
}
