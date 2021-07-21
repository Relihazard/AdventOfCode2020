package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type IntSlice []int64

func (list IntSlice) Has(i int64) bool {
	for _, val := range list {
		if val == i {
			return true
		}
	}
	return false
}

func dayTen() {
	fmt.Printf("\033[1;34m%s\033[0m\n", "Day ten:")

	file, err := os.Open("inputs/2020/10")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var arr IntSlice

	for scanner.Scan() {
		val, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		arr = append(arr, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	arr = append(arr, 0)

	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	arr = append(arr, arr[len(arr)-1]+3)

	var ones, threes int

	for i := range arr {
		if arr.Has(arr[i] + 1) {
			ones++
		} else if arr.Has(arr[i] + 3) {
			threes++
		}
	}

	fmt.Printf("\033[1;36m%d\033[0m\n", ones*threes)

	m := make(map[int64]int64)
	m[0] = 1

	var one, two, three int64

	for _, val := range arr[1:] {
		if val2, ok := m[val-1]; ok {
			one = val2
		}
		if val2, ok := m[val-2]; ok {
			two = val2
		}
		if val2, ok := m[val-3]; ok {
			three = val2
		}
		m[val] = one + two + three
		one, two, three = 0, 0, 0
	}

	fmt.Printf("\033[1;36m%d\033[0m\n", m[arr[len(arr)-1]])
}
