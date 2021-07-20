package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const Preamble = 25

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

func dayNine() {
	fmt.Printf("\033[1;34m%s\033[0m\n", "Day nine:")

	file, err := os.Open("inputs/2020/09")

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

	arr := data[0:Preamble]

	var limit int64

	for _, val := range data[25:] {
		if !checkNumber(arr, val) {
			fmt.Printf("\033[1;36m%d\033[0m\n", val)
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
				fmt.Printf("\033[1;36m%d\033[0m\n", arr[0]+arr[len(arr)-1])
			} else if sum > limit {
				arr = nil
				break
			}
		}
	}
}
