package main

import (
	"fmt"
	"slices"
	"strconv"
)

func taskOne(rules []string, data [][]string) int {
	sum := 0

	for l := 0; l < len(data); l++ {
		if isLineValid(rules, data[l]) {
			mid := len(data[l]) / 2
			n, _ := strconv.Atoi(data[l][mid])
			sum += n
		}
	}

	return sum
}

func taskTwo(rules []string, data [][]string) int {
	sum := 0

	for i := 0; i < len(data); i++ {
		valid := true

		for l := 0; l < len(data[i]); l++ {
			for j := l + 1; j < len(data[i]); j++ {
				rule := fmt.Sprintf("%s|%s", data[i][j], data[i][l])

				if slices.Contains(rules, rule) {
					if valid {
						valid = false
					}

					data[i][j], data[i][l] = data[i][l], data[i][j]
				}
			}
		}

		if !valid {
			mid := len(data[i]) / 2
			n, _ := strconv.Atoi(data[i][mid])
			sum += n
		}
	}

	return sum
}

func main() {
	rules, data := getTaskInput()

	taskOneResult := taskOne(rules, data)

	fmt.Println("Task one result: ", taskOneResult)

	taskTwoResult := taskTwo(rules, data)

	fmt.Println("Task two result: ", taskTwoResult)
}
