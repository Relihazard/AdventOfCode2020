package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func dayFivePartOne(a []string) []int {
	var list []int

	for _, line := range a {
		row := []int{0, 127}
		column := []int{0, 7}

		rowString := line[0:7]
		columnString := line[7:]

		for _, c := range rowString {
			if c == 'F' {
				row[1] = (row[0] + row[1]) / 2
			} else {
				row[0] = (row[0]+row[1])/2 + 1
			}
		}

		for _, c := range columnString {
			if c == 'L' {
				column[1] = (column[0] + column[1]) / 2
			} else {
				column[0] = (column[0]+column[1])/2 + 1
			}
		}

		id := row[0]*8 + column[0]
		list = append(list, id)
	}

	sort.Ints(list)
	fmt.Printf("\033[1;36m%d\033[0m\n", list[len(list)-1])

	return list
}

func dayFivePartTwo(list []int) {
	for i, value := range list {
		if i > 0 && list[i-1]+1 != value {
			fmt.Printf("\033[1;36m%d\033[0m\n", value-1)
			return
		}
	}
}

func dayFive() {
	fmt.Printf("\033[1;34m%s\033[0m\n", "Day five:")

	file, err := os.Open("inputs/2020/05")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var a []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 0 {
			a = append(a, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	list := dayFivePartOne(a)
	dayFivePartTwo(list)
}
