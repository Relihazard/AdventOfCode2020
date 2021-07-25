package twenty

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func dayOnePartOne(a []int) int {
	m := make(map[int]bool)
	for _, value := range a {
		if m[value] {
			return value * (2020 - value)
		}

		m[2020-value] = true
	}

	return -1
}

func dayOnePartTwo(a []int) int {
	m := make(map[int]bool)
	for i, value := range a {
		currSum := 2020 - value

		for _, value2 := range a[i+1:] {
			if m[currSum-value2] {
				return value * value2 * (currSum - value2)
			}
			m[value2] = true
		}
	}

	return -1
}

func DayOne() (int, int) {
	file, err := os.Open("data/inputs/twenty/01")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var a []int

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if i != 2020 {
			a = append(a, i)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return dayOnePartOne(a), dayOnePartTwo(a)
}
