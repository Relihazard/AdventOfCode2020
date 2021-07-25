package twenty

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

func daySevenPartOne(ruleset Ruleset) int {
	list := make(map[string]bool)

	ruleset.findBagsContainingColor("shiny gold", &list)

	return len(list)
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

func daySevenPartTwo(ruleset Ruleset) int {
	return ruleset.countBags("shiny gold")
}

func DaySeven() (int, int) {
	file, err := os.Open("data/inputs/twenty/07")

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

	return daySevenPartOne(ruleset), daySevenPartTwo(ruleset)
}
