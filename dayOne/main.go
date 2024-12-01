package main

import "fmt"

func taskOne() int {
    leftList, rightList := getTaskInput();
    
    bubbleSort(leftList);
    bubbleSort(rightList);

    inputLen := len(leftList);

    var result int = 0;

    for i := 0; i < inputLen; i++ {
        dif := leftList[i] - rightList[i];

        if(dif < 0) {
            dif *= -1;
        }

        result += dif;
    }

    return result;
}

func appearsCount(n int, rightList []int) int {
    var count = 0;


    for i := 0; i < len(rightList); i++ {
        if n == rightList[i] {
            count++
        }
    }

    return count;
}



func taskTwo() int {
    leftList, rightList := getTaskInput();

    var result int = 0;


    for i := 0; i < len(leftList); i++ {
        value := leftList[i];

        count := appearsCount(value, rightList); 

        result += value * count;
    }

    return result;
}


func main() {
	taskOneResult := taskOne()

    fmt.Println("ID's difference: ", taskOneResult);


    taskTwoResult := taskTwo();

    fmt.Println("Task two result: ", taskTwoResult);
}
