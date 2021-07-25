package twenty

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func checkNumber(arr []int64, number int64) bool {
	for i := range arr {
		for _, val := range arr[i+1:] {
			if arr[i]+val == number {
				return true
			}
		}
	}

	return false
}

func arraySum(arr []int64) int64 {
	var sum int64

	for _, val := range arr {
		sum += val
	}

	return sum
}

func DayNine() (int, int) {
	file, err := os.Open("data/inputs/twenty/09")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data []int64

	for scanner.Scan() {
		number, _ := strconv.ParseInt(scanner.Text(), 10, 0)

		data = append(data, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var preamble int

	if len(data) < 25 {
		preamble = 5
	} else {
		preamble = 25
	}

	arr := data[0:preamble]

	var limit int64
	var partOne, partTwo int

	for _, val := range data[preamble:] {
		if !checkNumber(arr, val) {
			partOne = int(val)
			limit = val
			break
		}
		arr = arr[1:]
		arr = append(arr, val)
	}

	arr = nil

	for i := range data {
		arr = append(arr, data[i])
		for _, val := range data[i+1:] {
			arr = append(arr, val)

			sum := arraySum(arr)

			if sum == limit {
				sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
				partTwo = int(arr[0] + arr[len(arr)-1])
			} else if sum > limit {
				arr = nil
				break
			}
		}
	}

	return partOne, partTwo
}
