package main

import (
	"fmt"
)

func taskOne() int {
    var taskOneInput [][]int = getTaskInput();

    var count int = 0;

    for i := 0; i < len(taskOneInput); i++ {
        if isReportSafe(taskOneInput[i], false) {
            count++;
        }
    }

    return count;
}



func taskTwo() int {
    var taskOneInput [][]int = getTaskInput();

    var count int = 0;

    for i := 0; i < len(taskOneInput); i++ {
        if isReportSafe(taskOneInput[i], true) {
            count++;
        }
    }

    return count;
}


func main() {
    taskOneResult := taskOne();

    fmt.Println("Task on result: ", taskOneResult);

    taskTwoResult := taskTwo();

    fmt.Println("Task two result: ", taskTwoResult);
}
