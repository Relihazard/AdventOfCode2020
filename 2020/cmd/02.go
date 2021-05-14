package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func dayTwoPartOne(a []string) {
	var counter int

	for _, line := range a {
		splits := strings.Split(line, " ")
		limits := strings.Split(splits[0], "-")
		lowerLimit, _ := strconv.ParseInt(limits[0], 10, 0)
		upperLimit, _ := strconv.ParseInt(limits[1], 10, 0)
		c := splits[1][:1]
		count := strings.Count(line, c)

		if count >= int(lowerLimit) && count <= int(upperLimit) {
			counter++
		}
	}

	fmt.Printf("\033[1;36m%d\033[0m\n", counter)
}

func dayTwoPartTwo(a []string) {
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

	fmt.Printf("\033[1;36m%d\033[0m\n", counter)
}

func dayTwo() {
	fmt.Printf("\033[1;34m%s\033[0m\n", "Day two:")

	file, err := os.Open("inputs/2020/02")

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

	dayTwoPartOne(a)
	dayTwoPartTwo(a)
}
