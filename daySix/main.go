package main

import (
	"fmt"
	"slices"
)

func taskOne(tMap taskMap) int {
	positions := map[string]bool{}

	key := fmt.Sprintf("%v|%v", tMap.g.posX, tMap.g.posY)

	positions[key] = true

	for !tMap.exited {
		tMap.move()

		key := fmt.Sprintf("%v|%v", tMap.g.posX, tMap.g.posY)

		positions[key] = true
	}

	return len(positions)
}

func main() {
	tMap := getTaskInput()

	taskOneResult := taskOne(tMap)

	fmt.Println("Task one result: ", taskOneResult)
}

func drawPath(tMap taskMap, path []string) {
	const colorGreen = "\033[32m"
	const colorRed = "\033[0;31m"
	const colorNone = "\033[0m"

	for i := 0; i < len(tMap.data); i++ {
		for l := 0; l < len(tMap.data[i]); l++ {
			point := fmt.Sprintf("%v|%v", i, l)

			if 89 == i && 51 == l {
				fmt.Print(colorGreen)
				fmt.Print("^")
				fmt.Print(colorNone)
				continue
			}

			if slices.Contains(path, point) {
				fmt.Print(colorRed)
				fmt.Print("X")
				fmt.Print(colorNone)

				continue
			}

			fmt.Print(tMap.data[i][l])
		}

		fmt.Print("\n")
	}
}
