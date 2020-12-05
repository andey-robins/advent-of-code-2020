package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// part 1 answer is not 234 (too low) it is 235
// part 2 answer is not 173 (too low); 195 (too high) it is 194

// so initially i had a simple error with not appending the last read passport, leading to an incorrect first attempt
// on the answer to part 1. the first simple error in part 2 was using && instead of || in the eye color checks, the second was
// caused by not checking the first character of the hcl string was a '#' if the rest of the string would otherwise be valid
// overall today was one that i did at the end of a caffeine crash and in the lead up to finals week, so chock it up
// to not being totally mentally there during work

type Passport struct {
	Byr int
	Iyr int
	Eyr int
	Hgt int
	Hcl string
	Ecl string
	Pid string
	Cid int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// open and read in the scanned entries
	f, err := os.Open("./passports.txt")
	check(err)
	defer func() {
		err := f.Close()
		check(err)
	}()

	s := bufio.NewScanner(f)

	passports := []Passport{}
	activePassport := Passport{}

	for s.Scan() {

		line := s.Text()

		if line == "" {
			passports = append(passports, activePassport)
			activePassport = Passport{}
		} else {
			lineData := strings.Split(line, " ")
			for _, entry := range lineData {
				entries := strings.Split(entry, ":")

				if entries[0] == "byr" {

					data, err := strconv.Atoi(entries[1])
					check(err)
					// validation
					if data <= 2002 && data >= 1920 {
						activePassport.Byr = data
					}

				} else if entries[0] == "iyr" {

					data, err := strconv.Atoi(entries[1])
					check(err)
					// validation
					if data <= 2020 && data >= 2010 {
						activePassport.Iyr = data
					}

				} else if entries[0] == "eyr" {

					data, err := strconv.Atoi(entries[1])
					check(err)
					// validation
					if data <= 2030 && data >= 2020 {
						activePassport.Eyr = data
					}

				} else if entries[0] == "hgt" {

					// validation
					if strings.IndexAny(entries[1], "cm") > -1 {
						// centimeter
						data, err := strconv.Atoi(entries[1][:len(entries[1])-2])
						check(err)
						if data <= 193 && data >= 150 {
							activePassport.Hgt = data
						}
					} else if strings.IndexAny(entries[1], "in") > -1 {
						// inches
						data, err := strconv.Atoi(entries[1][:len(entries[1])-2])
						check(err)
						if data >= 59 && data <= 76 {
							activePassport.Hgt = data
						}
					}

				} else if entries[0] == "hcl" {

					// validation
					_, err := strconv.ParseInt(entries[1][1:], 16, 0)
					if err == nil && entries[1][:1] == "#" {
						activePassport.Hcl = entries[1]
					}

				} else if entries[0] == "ecl" {

					// validation
					if strings.Index(entries[1], "amb") > -1 || strings.Index(entries[1], "blu") > -1 || strings.Index(entries[1], "brn") > -1 || strings.Index(entries[1], "gry") > -1 || strings.Index(entries[1], "grn") > -1 || strings.Index(entries[1], "hzl") > -1 || strings.Index(entries[1], "oth") > -1 {
						activePassport.Ecl = entries[1]
					}

				} else if entries[0] == "pid" {

					// validation
					_, err := strconv.Atoi(entries[1])
					if len(entries[1]) == 9 && err == nil {
						activePassport.Pid = entries[1]
					}

				} else if entries[0] == "cid" {

					// no validaiton needed, simply accepted
					data, err := strconv.Atoi(entries[1])
					check(err)
					activePassport.Cid = data

				} else {
					panic("invalid entry")
				}
			}
		}
	}

	passports = append(passports, activePassport)

	validPassports := 0
	for _, passport := range passports {
		if passport.Byr != 0 && passport.Ecl != "" && passport.Eyr != 0 && passport.Hcl != "" && passport.Hgt != 0 && passport.Iyr != 0 && passport.Pid != "" {
			validPassports++
		}
	}

	fmt.Printf("Found %v valid passports", validPassports)

}
