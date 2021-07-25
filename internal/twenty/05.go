package twenty

import (
	"bufio"
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

	return list
}

func dayFivePartTwo(list []int) int {
	for i, value := range list {
		if i > 0 && list[i-1]+1 != value {
			return value - 1
		}
	}
	return -1
}

func DayFive() (int, int) {
	file, err := os.Open("data/inputs/twenty/05")

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
	return list[len(list)-1], dayFivePartTwo(list)
}
