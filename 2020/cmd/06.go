package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Group struct {
	questions map[rune]int
	size      int
}

func newGroup() *Group {
	g := Group{questions: make(map[rune]int)}

	return &g
}

func daySixPartOne(a []Group) {
	result := 0

	for _, group := range a {
		result += len(group.questions)
	}

	fmt.Printf("\033[1;36m%d\033[0m\n", result)
}

func daySixPartTwo(a []Group) {
	result := 0

	for _, group := range a {
		for _, value := range group.questions {
			if value == group.size {
				result++
			}
		}
	}

	fmt.Printf("\033[1;36m%d\033[0m\n", result)
}

func daySix() {
	fmt.Printf("\033[1;34m%s\033[0m\n", "Day six:")

	file, err := os.Open("inputs/2020/06")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var a []Group
	g := newGroup()

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			a = append(a, *g)
			g = newGroup()
			continue
		}
		for _, c := range line {
			g.questions[c]++
		}
		g.size++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	a = append(a, *g)

	daySixPartOne(a)
	daySixPartTwo(a)
}
