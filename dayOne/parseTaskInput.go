package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)


func getTaskInput() ([]int, []int) {
    file, err := os.Open("input.txt")

    if err != nil {
        panic(err);
    }

    scanner := bufio.NewScanner(file);

    var leftList []int = []int {};
    var rightList []int = []int {};

    for scanner.Scan() {
        line := scanner.Text();

        lineAsArray := strings.Split(line, "   ");

        numberOne, _ := strconv.Atoi(lineAsArray[0]);
        numberTwo, _ := strconv.Atoi(lineAsArray[1]);

        leftList = append(leftList, numberOne);
        rightList = append(rightList, numberTwo);
    }


    return leftList, rightList;
}

