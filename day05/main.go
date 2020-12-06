package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// part 1 answer is not 970 (too high); is 896
// part 2 answer is 659

// the big hickup today was that i switched up columns and rows in my naming convention. it was something that could be easily refactored,
// but i instead fixed it when loading into the data struct so that it would behave as expected after loading, but the loading would look
// goofy. \/^o^\/

// otherwise, today was realitvely simple and the only other problem that arrose was i forgot my seat wouldn't have a seatid associated with it
// and we would need to calculate it from the surrounding seats.

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type BoardingPass struct {
	row    int64
	column int64
	seatId int64
}

type Seat struct {
	full   bool
	seatId int64
}

// replaces a rune at location loc in a string
func replaceAtIndex(input string, replacement rune, loc int) string {
	out := []rune(input)
	out[loc] = replacement
	return string(out)
}

func main() {
	// select part one or two
	partOne := true
	if len(os.Args) == 2 && os.Args[1] == "2" {
		partOne = false
	}

	f, err := os.Open("./boardingpass.txt")
	check(err)
	defer func() {
		err := f.Close()
		check(err)
	}()

	s := bufio.NewScanner(f)
	boardingPasses := []BoardingPass{}
	for s.Scan() {
		line := s.Text()

		// i assigned the column and row variables backwards during reading. It is fixed in the struct
		column := line[:7]
		for i, char := range line[:7] {
			if char == 'F' {
				column = replaceAtIndex(column, '0', i)
			} else {
				column = replaceAtIndex(column, '1', i)
			}
		}
		column = "0b" + column

		row := line[7:]
		for i, char := range line[7:] {
			if char == 'R' {
				row = replaceAtIndex(row, '1', i)
			} else {
				row = replaceAtIndex(row, '0', i)
			}
		}
		row = "0b" + row

		rowNumber, err := strconv.ParseInt(row, 0, 8)
		check(err)
		columnNumber, err := strconv.ParseInt(column, 0, 8)
		check(err)
		nextBoardingPass := BoardingPass{
			column: rowNumber,
			row:    columnNumber,
			seatId: columnNumber*8 + rowNumber,
		}
		boardingPasses = append(boardingPasses, nextBoardingPass)
	}

	if partOne {
		var max = int64(0)
		for _, pass := range boardingPasses {
			if pass.seatId > max {
				max = pass.seatId
			}
		}

		fmt.Printf("Part 1 answer: %v\n", max)
	} else {

		planeSeats := make([]Seat, 128*8)
		for _, pass := range boardingPasses {
			planeSeats[pass.row*8+pass.column] = Seat{
				full:   true,
				seatId: pass.seatId,
			}
		}

		throughBeginning := false
		var mySeatId = int64(0)
		for i, seat := range planeSeats {
			if seat.full && !throughBeginning {
				throughBeginning = true
			} else if !seat.full && throughBeginning && mySeatId == 0 {
				mySeatId = planeSeats[i-1].seatId + 1
			}
		}
		fmt.Printf("Part 2 answer: %v\n", mySeatId)
	}
}
