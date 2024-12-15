package main

import (
	"bufio"
	"os"
	"strings"
)

func getTaskInput() ([]string, [][]string) {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var rules []string = []string{}
	var data [][]string = [][]string{}

	scanningRules := true

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Trim(line, " ") == "" {
			scanningRules = false
			continue
		}

		if scanningRules {
			rules = append(rules, line)
		} else {
			data = append(data, strings.Split(line, ","))
		}
	}

	return rules, data
}
