package twenty

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Op struct {
	name    string
	arg     int
	f       func(int, *int) int
	visited bool
}

func nop(arg int, acc *int) int {
	return 1
}

func jmp(arg int, acc *int) int {
	return arg
}

func acc(arg int, acc *int) int {
	*acc += arg
	return 1
}

func flip(name string) (string, func(int, *int) int) {
	if name == "nop" {
		return "jmp", jmp
	}
	return "nop", nop
}

func move(ops []Op) int {
	var acc int
	var i int

	for {
		if ops[i].visited {
			return acc
		}
		ops[i].visited = true
		i += ops[i].f(ops[i].arg, &acc)
	}
}

func movePartTwo(ops []Op) int {
	var acc int
	var i int

	for {
		if i >= len(ops) {
			return acc
		}
		if ops[i].visited {
			return 0
		}
		ops[i].visited = true
		i += ops[i].f(ops[i].arg, &acc)
	}
}

func DayEight() (int, int) {
	file, err := os.Open("data/inputs/twenty/08")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ops []Op

	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, " ")

		arg, _ := strconv.ParseInt(splits[1], 10, 0)

		var f func(int, *int) int

		switch splits[0] {
		case "nop":
			f = nop
		case "jmp":
			f = jmp
		case "acc":
			f = acc
		}

		ops = append(ops, Op{splits[0], int(arg), f, false})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	moveCount := move(ops)

	for i := range ops {
		ops[i].visited = false
	}

	for i, val := range ops {
		opsCopy := make([]Op, len(ops))

		copy(opsCopy, ops)

		if val.name == "nop" || val.name == "jmp" {
			opsCopy[i].name, opsCopy[i].f = flip(val.name)
		}

		if acc := movePartTwo(opsCopy); acc != 0 {
			return moveCount, acc
		}
	}

	return -1, -1
}
