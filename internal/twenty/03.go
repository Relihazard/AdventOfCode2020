package twenty

import (
	"bufio"
	"log"
	"os"
)

func dayThreePartOne(a []string, lineLen int, limit int) int {
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

	return trees
}

func dayThreePartTwo(a []string, lineLen int, limit int) int {
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

	return total
}

func DayThree() (int, int) {
	file, err := os.Open("data/inputs/twenty/03")

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

	lineLen := len(a[0])
	limit := len(a)

	return dayThreePartOne(a, lineLen, limit), dayThreePartTwo(a, lineLen, limit)
}
