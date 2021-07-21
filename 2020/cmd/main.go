package main

import (
	"flag"
	"fmt"
)

func main() {
	days := []func(){
		dayOne,
		dayTwo,
		dayThree,
		dayFour,
		dayFive,
		daySix,
		daySeven,
		dayEight,
		dayNine,
		dayTen,
	}

	dayPtr := flag.Int("d", 0, "The day to be executed")
	flag.IntVar(dayPtr, "day", 0, "Same as -d")

	flag.Parse()

	if *dayPtr == 0 {
		for _, f := range days {
			f()
			fmt.Println()
		}
	} else {
		days[*dayPtr-1]()
		fmt.Println()
	}
}
