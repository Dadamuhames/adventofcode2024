package main

import (
	"bufio"
	"os"
	"strings"
)

func getTaskInput() taskMap {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	grd := guard{}
	var data [][]int = [][]int{}

	x := 0

	for scanner.Scan() {
		line := scanner.Text()

		lineAsList := strings.Split(line, "")

		var lineAsIntList []int = []int{}

		for i := 0; i < len(lineAsList); i++ {
			if lineAsList[i] == "#" {
				lineAsIntList = append(lineAsIntList, 1)
			} else if lineAsList[i] == "." {
				lineAsIntList = append(lineAsIntList, 0)
			} else {
				grd.posX = x
				grd.posY = i
				grd.dir = lineAsList[i]
				lineAsIntList = append(lineAsIntList, 0)
			}
		}

		data = append(data, lineAsIntList)
		x++
	}

	return taskMap{g: grd, data: data, exited: false}
}
