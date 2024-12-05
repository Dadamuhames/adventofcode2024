package main

import (
	"bufio"
	"os"
	"strings"
)



func getTaskInput() [][]string {
     file, err := os.Open("input.txt");

    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file);

    result := make([][]string, 0);

    for scanner.Scan() {
        line := scanner.Text();

        result = append(result, strings.Split(line, ""));
    }
    
    return result;   
}
