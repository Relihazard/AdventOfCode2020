package twenty

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func dayTwoPartOne(a []string) int {
	var counter int

	for _, line := range a {
		splits := strings.Split(line, " ")
		limits := strings.Split(splits[0], "-")
		lowerLimit, _ := strconv.ParseInt(limits[0], 10, 0)
		upperLimit, _ := strconv.ParseInt(limits[1], 10, 0)
		c := splits[1][0]
		count := 0

		for _, char := range splits[2] {
			if rune(c) == char {
				count++
			}
		}

		if count >= int(lowerLimit) && count <= int(upperLimit) {
			counter++
		}
	}

	return counter
}

func dayTwoPartTwo(a []string) int {
	var counter int

	for _, line := range a {
		splits := strings.Split(line, " ")
		limits := strings.Split(splits[0], "-")
		firstPosition, _ := strconv.Atoi(limits[0])
		secondPosition, _ := strconv.Atoi(limits[1])
		c := splits[1][:1]

		if splits[2][firstPosition-1] == c[0] && splits[2][secondPosition-1] == c[0] {
			continue
		}
		if splits[2][firstPosition-1] == c[0] || splits[2][secondPosition-1] == c[0] {
			counter++
		}
	}

	return counter
}

func DayTwo() (int, int) {
	file, err := os.Open("data/inputs/twenty/02")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var a []string

	for scanner.Scan() {
		a = append(a, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return dayTwoPartOne(a), dayTwoPartTwo(a)
}
