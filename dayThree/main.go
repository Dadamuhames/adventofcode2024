package main

import (
	"fmt"
	"strconv"
	"strings"
)



func mul(nums []string) int {

    if len(nums) != 2 {
        panic("nums length invalid");
    }

    a, _ := strconv.Atoi(nums[0]);
    b, _ := strconv.Atoi(nums[1]);

    return a * b;
}


func solution(taskInput []string) int {
    result := 0;

    for i := 0; i < len(taskInput); i++ {
        str := strings.Replace(taskInput[i], "mul(", "", 1);
        str = strings.Replace(str, ")", "", 1);
        
        numsPair := strings.Split(str, ",");
        
        result += mul(numsPair);
    }

    return result;
}


func main() {
    taskOneInput := parseValidExpressions();

    taskOneResult := solution(taskOneInput);

    fmt.Println("Task one result: ", taskOneResult);

    
    taskTwoInput := parseValidExpressionsWithInstructions();

    taskTwoResult := solution(taskTwoInput);

    fmt.Println("Task two result: ", taskTwoResult);
}
