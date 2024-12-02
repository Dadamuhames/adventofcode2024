package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func getTaskInput() [][]int {
    file, err := os.Open("input.txt")

    if err != nil {
        panic(err);
    }

    scanner := bufio.NewScanner(file);

    var result [][]int = [][]int {};

    for scanner.Scan() {
        line := scanner.Text();

        var lineAsCharArray []string = strings.Split(line, " ");
        var lineAsIntArray []int = []int {};

        for i := 0; i < len(lineAsCharArray); i++ {
            num, _ := strconv.Atoi(lineAsCharArray[i]);

            lineAsIntArray = append(lineAsIntArray, num);
        }

        result = append(result, lineAsIntArray);
    }

    return result;    
}
