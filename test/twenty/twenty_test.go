package twenty_test

import (
	"testing"

	"github.com/relihazard/adventofcode/internal/twenty"
)

const PartOneFailMessage = "Part one is wrong: expected %d, got %d"
const PartTwoFailMessage = "Part two is wrong: expected %d, got %d"

func TestDayOne(t *testing.T) {
	partOneExpected, partTwoExpected := 514579, 241861950

	partOne, partTwo := twenty.DayOne()
	if partOne != partOneExpected {
		t.Fatalf(PartOneFailMessage, partOneExpected, partOne)
	}
	if partTwo != partTwoExpected {
		t.Fatalf(PartTwoFailMessage, partTwoExpected, partTwo)
	}
}
func TestDayTwo(t *testing.T) {
	partOneExpected, partTwoExpected := 2, 1

	partOne, partTwo := twenty.DayTwo()
	if partOne != partOneExpected {
		t.Fatalf(PartOneFailMessage, partOneExpected, partOne)
	}
	if partTwo != partTwoExpected {
		t.Fatalf(PartTwoFailMessage, partTwoExpected, partTwo)
	}
}

func TestDayThree(t *testing.T) {
	partOneExpected, partTwoExpected := 7, 336

	partOne, partTwo := twenty.DayThree()
	if partOne != partOneExpected {
		t.Fatalf(PartOneFailMessage, partOneExpected, partOne)
	}
	if partTwo != partTwoExpected {
		t.Fatalf(PartTwoFailMessage, partTwoExpected, partTwo)
	}
}

func TestDayFour(t *testing.T) {
	partOneExpected, partTwoExpected := 2, 2

	partOne, partTwo := twenty.DayFour()
	if partOne != partOneExpected {
		t.Fatalf(PartOneFailMessage, partOneExpected, partOne)
	}
	if partTwo != partTwoExpected {
		t.Fatalf(PartTwoFailMessage, partTwoExpected, partTwo)
	}
}

func TestDayFive(t *testing.T) {
	partOneExpected, partTwoExpected := 820, 566

	partOne, partTwo := twenty.DayFive()
	if partOne != partOneExpected {
		t.Fatalf(PartOneFailMessage, partOneExpected, partOne)
	}
	if partTwo != partTwoExpected {
		t.Fatalf(PartTwoFailMessage, partTwoExpected, partTwo)
	}
}

func TestDaySix(t *testing.T) {
	partOneExpected, partTwoExpected := 11, 6

	partOne, partTwo := twenty.DaySix()
	if partOne != partOneExpected {
		t.Fatalf(PartOneFailMessage, partOneExpected, partOne)
	}
	if partTwo != partTwoExpected {
		t.Fatalf(PartTwoFailMessage, partTwoExpected, partTwo)
	}
}

func TestDaySeven(t *testing.T) {
	partOneExpected, partTwoExpected := 4, 32

	partOne, partTwo := twenty.DaySeven()
	if partOne != partOneExpected {
		t.Fatalf(PartOneFailMessage, partOneExpected, partOne)
	}
	if partTwo != partTwoExpected {
		t.Fatalf(PartTwoFailMessage, partTwoExpected, partTwo)
	}
}

func TestDayEight(t *testing.T) {
	partOneExpected, partTwoExpected := 5, 8

	partOne, partTwo := twenty.DayEight()
	if partOne != partOneExpected {
		t.Fatalf(PartOneFailMessage, partOneExpected, partOne)
	}
	if partTwo != partTwoExpected {
		t.Fatalf(PartTwoFailMessage, partTwoExpected, partTwo)
	}
}

func TestDayNine(t *testing.T) {
	partOneExpected, partTwoExpected := 127, 62

	partOne, partTwo := twenty.DayNine()
	if partOne != partOneExpected {
		t.Fatalf(PartOneFailMessage, partOneExpected, partOne)
	}
	if partTwo != partTwoExpected {
		t.Fatalf(PartTwoFailMessage, partTwoExpected, partTwo)
	}
}

func TestDayTen(t *testing.T) {
	partOneExpected, partTwoExpected := 35, 8

	partOne, partTwo := twenty.DayTen()
	if partOne != partOneExpected {
		t.Fatalf(PartOneFailMessage, partOneExpected, partOne)
	}
	if partTwo != partTwoExpected {
		t.Fatalf(PartTwoFailMessage, partTwoExpected, partTwo)
	}
}
