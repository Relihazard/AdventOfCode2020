package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func dayThreePartOne(a []string, lineLen int, limit int) {
	var trees, i, j int

	for {
		j += 3
		if j >= lineLen {
			j = j - lineLen
		}
		i++
		if a[i][j] == '#' {
			trees++
		}
		if i == (limit - 1) {
			break
		}
	}

	fmt.Printf("\033[1;36m%d\033[0m\n", trees)
}

func dayThreePartTwo(a []string, lineLen int, limit int) {
	slopes := [][]int{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}
	trees := make([]int, 5)

	for index, value := range slopes {
		right := value[1]
		bottom := value[0]
		var i, j int

		for {
			j += right
			if j >= lineLen {
				j = j - lineLen
			}
			i += bottom
			if a[i][j] == '#' {
				trees[index]++
			}
			if i == (limit - 1) {
				break
			}
		}
	}

	total := 1

	for _, value := range trees {
		total = total * value
	}

	fmt.Printf("\033[1;36m%d\033[0m\n", total)
}

func dayThree() {
	fmt.Printf("\033[1;34m%s\033[0m\n", "Day three:")

	file, err := os.Open("inputs/2020/03")

	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var a []string

	for scanner.Scan() {
		a = append(a, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}

	lineLen := len(a[0])
	limit := len(a)

	dayThreePartOne(a, lineLen, limit)
	dayThreePartTwo(a, lineLen, limit)
}
