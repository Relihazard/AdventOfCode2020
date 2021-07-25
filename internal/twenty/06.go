package twenty

import (
	"bufio"
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

func daySixPartOne(a []Group) int {
	result := 0

	for _, group := range a {
		result += len(group.questions)
	}

	return result
}

func daySixPartTwo(a []Group) int {
	result := 0

	for _, group := range a {
		for _, value := range group.questions {
			if value == group.size {
				result++
			}
		}
	}

	return result
}

func DaySix() (int, int) {
	file, err := os.Open("data/inputs/twenty/06")

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

	return daySixPartOne(a), daySixPartTwo(a)
}
