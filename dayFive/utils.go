package main

import (
	"fmt"
	"slices"
)

func isLineValid(rules []string, line []string) bool {
	for i := 0; i < len(line); i++ {
		for j := i + 1; j < len(line); j++ {
			rule := fmt.Sprintf("%s|%s", line[j], line[i])

			if slices.Contains(rules, rule) {
				return false
			}
		}
	}

	return true
}
