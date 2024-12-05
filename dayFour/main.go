package main

import (
	"fmt"
	"strings"
)



func findMasForX(input [][]string, x int, y int) int {
    count := 0;


    // for X and down
   if x < len(input) - 3 {
        if input[x+1][y] == "M" && input[x+2][y] == "A" && input[x+3][y] == "S" {
            count++;
        }
    }

    // for X and up
    if x >= 3 {
        if input[x-1][y] == "M" && input[x-2][y] == "A" && input[x-3][y] == "S" {
            count++;
        }
    }


    // for X and left
    if y >= 3 {
      if input[x][y-1] == "M" && input[x][y-2] == "A" && input[x][y-3] == "S" {
        count++;
      }
    }


    // for X and right
    if y < len(input[0]) - 3 {
        if input[x][y+1] == "M" && input[x][y+2] == "A" && input[x][y+3] == "S" {
            count++;
        }
    }

    // for X and to left top corner
    if y >= 3 && x >= 3 {
        if input[x-1][y-1] == "M" && input[x-2][y-2] == "A" && input[x-3][y-3] == "S" {
            count++;;
        }
    }


    // for X and to right top corner
    if y < len(input[0]) - 3 && x >= 3 {
        if input[x-1][y+1] == "M" && input[x-2][y+2] == "A" && input[x-3][y+3] == "S" {
            count++;
        }
    }


    // for X and to left bottoom corner 
    if y >= 3 && x < len(input) - 3 {
        if input[x+1][y-1] == "M" && input[x+2][y-2] == "A" && input[x+3][y-3] == "S" {
            count++;
        }
    }


    // for X and to right bottom corner
    if y < len(input[0]) - 3 && x < len(input) - 3 {
        if input[x+1][y+1] == "M" && input[x+2][y+2] == "A" && input[x+3][y+3] == "S" {
            count++;
        }
    }



    return count;
}



func taskOne(input [][]string) int {
    count := 0;

    for i := 0; i < len(input); i++ {
        for j := 0; j < len(input[i]); j++ {
            if input[i][j] == "X" {
                count += findMasForX(input, i, j); 
            }
        }
    }    

    return count;
}


func isXMas(input [][]string, x int, y int) bool {
    if x >= 1 && x < len(input) - 1 && y >= 1 && y < len(input[0]) - 1 {
        ms := "";

        ms += input[x-1][y-1];
        ms += input[x-1][y+1];
        ms += input[x+1][y-1];
        ms += input[x+1][y+1];
        
        if  strings.Count(ms, "M") == 2 && strings.Count(ms, "S") == 2 {
            fmt.Println(ms);
        }

        if ms == "MSSM" || ms == "SMMS" {
            return false;
        }

        return strings.Count(ms, "M") == 2 && strings.Count(ms, "S") == 2;
    }

    return false;
}


func taskTwo(input [][]string) int {
    count := 0;

    for i := 0; i < len(input); i++ {
        for j := 0; j < len(input[i]); j++ {
            if input[i][j] == "A" {
               if isXMas(input, i, j) {
                    count++;
               } 
            }
        }
    }

    return count;
}


func main() {
    taskOneInput := getTaskInput();

    taskOneResult := taskOne(taskOneInput);

    fmt.Println("Task one result: ", taskOneResult);


    taskTwoResult := taskTwo(taskOneInput);

    fmt.Println("Task two result: ", taskTwoResult);
}
