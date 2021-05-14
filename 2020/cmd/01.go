package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func dayOnePartOne(a []int) {
	m := make(map[int]bool)
	for _, value := range a {
		if m[value] {
			fmt.Printf("\033[1;36m%d\033[0m\n", value*(2020-value))
		}

		m[2020-value] = true
	}
}

func dayOnePartTwo(a []int) {
	m := make(map[int]bool)
	for i, value := range a {
		currSum := 2020 - value

		for _, value2 := range a[i+1:] {
			if m[currSum-value2] {
				fmt.Printf("\033[1;36m%d\033[0m\n", value*value2*(currSum-value2))
				return
			}
			m[value2] = true
		}
	}
}

func dayOne() {
	fmt.Printf("\033[1;34m%s\033[0m\n", "Day one:")

	file, err := os.Open("inputs/2020/01")

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
			return
		}
		if i != 2020 {
			a = append(a, i)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	dayOnePartOne(a)
	dayOnePartTwo(a)
}
