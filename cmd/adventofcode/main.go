package main

import (
	"flag"
	"fmt"

	"github.com/relihazard/adventofcode/internal/twenty"
)

func main() {
	days := make(map[int][]func() (int, int))

	days[2020] = []func() (int, int){
		twenty.DayOne,
		twenty.DayTwo,
		twenty.DayThree,
		twenty.DayFour,
		twenty.DayFive,
		twenty.DaySix,
		twenty.DaySeven,
		twenty.DayEight,
		twenty.DayNine,
		twenty.DayTen,
	}

	yearPtr := flag.Int("y", 2020, "The year to be executed")
	flag.IntVar(yearPtr, "year", 2020, "Same as -y")

	dayPtr := flag.Int("d", 0, "The day to be executed")
	flag.IntVar(dayPtr, "day", 0, "Same as -d")

	flag.Parse()

	if *dayPtr <= 0 || *dayPtr > 24 {
		for i, f := range days[*yearPtr] {
			fmt.Printf("\033[1;34mDay %d:\033[0m\n", i+1)
			partOne, partTwo := f()
			fmt.Printf("\033[1;36m%d\033[0m\n", partOne)
			fmt.Printf("\033[1;36m%d\033[0m\n", partTwo)
			fmt.Println()
		}
	} else {
		days[*yearPtr][*dayPtr-1]()
		fmt.Println()
	}
}
