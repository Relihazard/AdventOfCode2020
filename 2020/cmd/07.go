package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Ruleset map[string]map[string]int

func (r *Ruleset) findBagsContainingColor(name string, list *map[string]bool) {
	for color, contents := range *r {
		if _, found := contents[name]; found {
			if _, ok := (*list)[color]; !ok {
				(*list)[color] = true
				r.findBagsContainingColor(color, list)
			}
		}
	}
}

func daySevenPartOne(ruleset Ruleset) {
	list := make(map[string]bool)

	ruleset.findBagsContainingColor("shiny gold", &list)

	fmt.Printf("\033[1;36m%d\033[0m\n", len(list))
}

func (r *Ruleset) countBags(name string) int {
	total := 0
	contents := (*r)[name]
	for color, count := range contents {
		subcount := r.countBags(color)
		if subcount > 0 {
			total += count * subcount
		}
		total += count
	}
	return total
}

func daySevenPartTwo(ruleset Ruleset) {
	fmt.Printf("\033[1;36m%d\033[0m\n", ruleset.countBags("shiny gold"))
}

func daySeven() {
	fmt.Printf("\033[1;34m%s\033[0m\n", "Day seven:")

	file, err := os.Open("inputs/2020/07")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	ruleset := make(Ruleset)

	for scanner.Scan() {
		line := scanner.Text()

		splits := strings.Split(line, " bags contain ")
		color := splits[0]
		regex := regexp.MustCompile(`\d\s\w*\s\w*`)
		substrings := regex.FindAllString(splits[1], -1)

		contents := make(map[string]int)
		for _, value := range substrings {
			count, description, bagColor := 0, "", ""
			_, _ = fmt.Sscanf(value, "%d %s %s", &count, &description, &bagColor)
			contents[description+" "+bagColor] = count
		}

		ruleset[color] = contents

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	daySevenPartOne(ruleset)
	daySevenPartTwo(ruleset)
}
