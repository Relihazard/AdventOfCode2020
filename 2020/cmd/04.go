package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func checkFields(m map[string]string, counter *int) {
	lenMap := len(m)
	_, found := m["cid"]
	if lenMap == 8 || (lenMap == 7 && !found) {
		*counter++
	}
}

func dayFourPartOne(a []map[string]string) {
	var counter int

	for _, m := range a {
		checkFields(m, &counter)
	}

	fmt.Printf("\033[1;36m%d\033[0m\n", counter)
}

func checkLimits(data string, limitMin int, limitMax int) bool {
	dataInt, err := strconv.ParseInt(data, 10, 0)

	if err != nil || (int(dataInt) < limitMin || int(dataInt) > limitMax) {
		return false
	}

	return true
}

func checkByr(m map[string]string) bool {
	data, ok := m["byr"]

	if !ok || len(data) != 4 {
		return false
	}

	return checkLimits(data, 1920, 2002)
}

func checkIyr(m map[string]string) bool {
	data, ok := m["iyr"]

	if !ok || len(data) != 4 {
		return false
	}

	return checkLimits(data, 2010, 2020)
}

func checkEyr(m map[string]string) bool {
	data, ok := m["eyr"]

	if !ok || len(data) != 4 {
		return false
	}

	return checkLimits(data, 2020, 2030)
}

func checkHgt(m map[string]string) bool {
	data, ok := m["hgt"]

	if !ok {
		return false
	}

	unit := data[len(data)-2:]

	if unit == "cm" {
		return checkLimits(data[0:3], 150, 193)
	} else if unit == "in" {
		return checkLimits(data[0:2], 59, 76)
	} else {
		return false
	}
}

func checkHclChars(data string) bool {
	for _, char := range data {
		if unicode.IsDigit(char) {
			continue
		} else if unicode.IsLower(char) {
			if char >= 'a' && char <= 'f' {
				continue
			}
		} else {
			return false
		}
	}

	return true
}

func checkHcl(m map[string]string) bool {
	data, ok := m["hcl"]

	if !ok || data[0] != '#' || len(data[1:]) != 6 || !checkHclChars(data[1:]) {
		return false
	}

	return true
}

func checkEcl(m map[string]string) bool {
	data, ok := m["ecl"]

	if !ok {
		return false
	}

	switch data {
	case
		"amb",
		"blu",
		"brn",
		"gry",
		"grn",
		"hzl",
		"oth":
		return true
	default:
		return false
	}
}

func checkPid(m map[string]string) bool {
	data, ok := m["pid"]

	if !ok || len(data) != 9 {
		return false
	}

	for _, char := range data {
		if !unicode.IsDigit(char) {
			return false
		}
	}

	return true
}

func dayFourParTwo(a []map[string]string) {
	var counter int

	for _, m := range a {
		if len(m) < 7 {
			continue
		} else {
			if !checkByr(m) || !checkIyr(m) || !checkEyr(m) || !checkHgt(m) || !checkHcl(m) || !checkEcl(m) || !checkPid(m) {
				continue
			}
			counter++
		}
	}

	fmt.Printf("\033[1;36m%d\033[0m\n", counter)
}

func dayFour() {
	fmt.Printf("\033[1;34m%s\033[0m\n", "Day four:")

	file, err := os.Open("inputs/2020/04")

	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()

	var a []map[string]string
	m := make(map[string]string, 8)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			a = append(a, m)
			m = make(map[string]string, 8)
			continue
		}

		splits := strings.Split(line, " ")

		for _, value := range splits {
			fields := strings.Split(value, ":")
			m[fields[0]] = fields[1]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}

	a = append(a, m)

	dayFourPartOne(a)
	dayFourParTwo(a)
}
